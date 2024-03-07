package metamorph

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/bitcoin-sv/arc/blocktx/blocktx_api"
	"github.com/bitcoin-sv/arc/metamorph/metamorph_api"
	"github.com/bitcoin-sv/arc/metamorph/processor_response"
	"github.com/bitcoin-sv/arc/metamorph/store"
	"github.com/libsv/go-p2p"
	"github.com/libsv/go-p2p/chaincfg/chainhash"
	"github.com/opentracing/opentracing-go"
	"github.com/ordishs/go-utils/stat"
	"github.com/ordishs/gocore"
)

const (
	// MaxRetries number of times we will retry announcing transaction if we haven't seen it on the network
	MaxRetries = 15
	// length of interval for checking transactions if they are seen on the network
	// if not we resend them again for a few times
	unseenTransactionRebroadcastingInterval = 60

	mapExpiryTimeDefault = 24 * time.Hour
	LogLevelDefault      = slog.LevelInfo

	failedToUpdateStatus       = "Failed to update status"
	dataRetentionPeriodDefault = 14 * 24 * time.Hour // 14 days

	maxMonitoriedTxs          = 100000
	loadUnminedLimit          = int64(5000)
	minimumHealthyConnections = 2

	processStatusUpdatesIntervalDefault  = 500 * time.Millisecond
	processStatusUpdatesBatchSizeDefault = 1000
)

var (
	ErrUnhealthy = errors.New("processor has less than 2 healthy peer connections")
)

type Processor struct {
	store                   store.MetamorphStore
	ProcessorResponseMap    *ProcessorResponseMap
	pm                      p2p.PeerManagerI
	mqClient                MessageQueueClient
	logger                  *slog.Logger
	mapExpiryTime           time.Duration
	dataRetentionPeriod     time.Duration
	now                     func() time.Time
	processExpiredTxsTicker *time.Ticker
	httpClient              HttpClient
	maxMonitoredTxs         int64

	quitListenTxChannel         chan struct{}
	quitListenTxChannelComplete chan struct{}
	minedTxsChan                chan *blocktx_api.TransactionBlocks

	storageStatusUpdateCh            chan store.UpdateStatus
	quitListenStatusUpdateCh         chan struct{}
	quitListenStatusUpdateChComplete chan struct{}
	processStatusUpdatesInterval     time.Duration
	processStatusUpdatesBatchSize    int

	startTime           time.Time
	queueLength         atomic.Int32
	queuedCount         atomic.Int32
	stored              *stat.AtomicStat
	announcedToNetwork  *stat.AtomicStats
	requestedByNetwork  *stat.AtomicStats
	sentToNetwork       *stat.AtomicStats
	acceptedByNetwork   *stat.AtomicStats
	seenInOrphanMempool *stat.AtomicStats
	seenOnNetwork       *stat.AtomicStats
	rejected            *stat.AtomicStats
	mined               *stat.AtomicStat
	retries             *stat.AtomicStat
}

type Option func(f *Processor)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func NewProcessor(s store.MetamorphStore, pm p2p.PeerManagerI, opts ...Option) (*Processor, error) {
	if s == nil {
		return nil, errors.New("store cannot be nil")
	}

	if pm == nil {
		return nil, errors.New("peer manager cannot be nil")
	}

	p := &Processor{
		startTime:               time.Now().UTC(),
		store:                   s,
		pm:                      pm,
		dataRetentionPeriod:     dataRetentionPeriodDefault,
		mapExpiryTime:           mapExpiryTimeDefault,
		now:                     time.Now,
		processExpiredTxsTicker: time.NewTicker(unseenTransactionRebroadcastingInterval * time.Second),
		maxMonitoredTxs:         maxMonitoriedTxs,

		quitListenTxChannel:         make(chan struct{}),
		quitListenTxChannelComplete: make(chan struct{}),

		quitListenStatusUpdateCh:         make(chan struct{}),
		quitListenStatusUpdateChComplete: make(chan struct{}),
		processStatusUpdatesInterval:     processStatusUpdatesIntervalDefault,
		processStatusUpdatesBatchSize:    processStatusUpdatesBatchSizeDefault,
		storageStatusUpdateCh:            make(chan store.UpdateStatus, processStatusUpdatesBatchSizeDefault),
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
		stored:              stat.NewAtomicStat(),
		announcedToNetwork:  stat.NewAtomicStats(),
		requestedByNetwork:  stat.NewAtomicStats(),
		sentToNetwork:       stat.NewAtomicStats(),
		acceptedByNetwork:   stat.NewAtomicStats(),
		seenInOrphanMempool: stat.NewAtomicStats(),
		seenOnNetwork:       stat.NewAtomicStats(),
		rejected:            stat.NewAtomicStats(),
		mined:               stat.NewAtomicStat(),
		retries:             stat.NewAtomicStat(),
	}

	p.logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: LogLevelDefault})).With(slog.String("service", "mtm"))

	// apply options to processor
	for _, opt := range opts {
		opt(p)
	}

	p.ProcessorResponseMap = NewProcessorResponseMap(p.mapExpiryTime, WithNowResponseMap(p.now))

	p.logger.Info("Starting processor", slog.String("cacheExpiryTime", p.mapExpiryTime.String()))

	// Start a goroutine to resend transactions that have not been seen on the network
	go p.processExpiredTransactions()
	go p.processMinedCallbacks()
	go p.processStatusUpdatesInStorage()

	gocore.AddAppPayloadFn("mtm", func() interface{} {
		return p.GetStats(false)
	})

	_ = newPrometheusCollector(p)

	return p, nil
}

// Shutdown closes all channels and goroutines gracefully
func (p *Processor) Shutdown() {
	p.logger.Info("Shutting down processor")

	err := p.unlockItems()
	if err != nil {
		p.logger.Error("Failed to unlock all hashes", slog.String("err", err.Error()))
	}
	p.processExpiredTxsTicker.Stop()
	p.quitListenTxChannel <- struct{}{}
	<-p.quitListenTxChannelComplete
	p.quitListenStatusUpdateCh <- struct{}{}
	<-p.quitListenStatusUpdateChComplete
}

func (p *Processor) unlockItems() error {
	items := p.ProcessorResponseMap.Items()
	hashes := make([]*chainhash.Hash, len(items))
	index := 0
	for key := range items {
		hash, err := chainhash.NewHash(key.CloneBytes())
		if err != nil {
			return err
		}
		hashes[index] = hash
		index++
	}

	p.logger.Info("unlocking items", slog.Int("number", len(hashes)))
	return p.store.SetUnlocked(context.Background(), hashes)
}

func (p *Processor) processMinedCallbacks() {
	go func() {
		defer func() {
			p.quitListenTxChannelComplete <- struct{}{}
		}()

		for {
			select {
			case <-p.quitListenTxChannel:
				return
			case txBlocks := <-p.minedTxsChan:
				updatedData, err := p.store.UpdateMined(context.Background(), txBlocks)
				if err != nil {
					p.logger.Error("failed to register transactions", slog.String("err", err.Error()))
				}

				for _, data := range updatedData {
					if data.CallbackUrl != "" {
						go p.SendCallback(p.logger, data)
					}
				}
			}
		}
	}()
}

func (p *Processor) CheckAndUpdate(statusUpdatesMap *map[chainhash.Hash]store.UpdateStatus, skipCheck bool) {
	if (skipCheck && len(*statusUpdatesMap) < p.processStatusUpdatesBatchSize) || len(*statusUpdatesMap) == 0 {
		return
	}

	statusUpdates := make([]store.UpdateStatus, 0, p.processStatusUpdatesBatchSize)
	for _, distinctStatusUpdate := range *statusUpdatesMap {
		statusUpdates = append(statusUpdates, distinctStatusUpdate)
	}

	err := p.statusUpdateWithCallback(statusUpdates)
	if err != nil {
		p.logger.Error("failed to bulk update statuses", slog.String("err", err.Error()))
	}

	*statusUpdatesMap = map[chainhash.Hash]store.UpdateStatus{}
}

func (p *Processor) processStatusUpdatesInStorage() {
	ticker := time.NewTicker(p.processStatusUpdatesInterval)

	go func() {
		defer func() {
			p.quitListenStatusUpdateChComplete <- struct{}{}
		}()

		statusUpdatesMap := map[chainhash.Hash]store.UpdateStatus{}

		for {
			select {
			case <-p.quitListenStatusUpdateCh:
				return
			case statusUpdate := <-p.storageStatusUpdateCh:
				// Ensure no duplicate hashes, overwrite value if the status has higher value than existing status
				foundStatusUpdate, found := statusUpdatesMap[statusUpdate.Hash]
				if !found || (found && statusValueMap[foundStatusUpdate.Status] < statusValueMap[statusUpdate.Status]) {
					statusUpdatesMap[statusUpdate.Hash] = statusUpdate
				}

				p.CheckAndUpdate(&statusUpdatesMap, true)
			case <-ticker.C:
				p.CheckAndUpdate(&statusUpdatesMap, false)
			}
		}
	}()
}

func (p *Processor) statusUpdateWithCallback(statusUpdates []store.UpdateStatus) error {
	updatedData, err := p.store.UpdateStatusBulk(context.Background(), statusUpdates)
	if err != nil {
		return err
	}

	for _, data := range updatedData {
		if ((data.Status == metamorph_api.Status_SEEN_ON_NETWORK || data.Status == metamorph_api.Status_SEEN_IN_ORPHAN_MEMPOOL) && data.FullStatusUpdates || data.Status == metamorph_api.Status_REJECTED) && data.CallbackUrl != "" {
			go p.SendCallback(p.logger, data)
		}
	}
	return nil
}

func (p *Processor) processExpiredTransactions() {
	span, _ := opentracing.StartSpanFromContext(context.Background(), "Processor:processExpiredTransactions")
	dbctx := opentracing.ContextWithSpan(context.Background(), span)
	defer span.Finish()

	// Periodically read unmined transactions from database and announce them again
	for range p.processExpiredTxsTicker.C {
		// define from what point in time we are interested in unmined transactions
		getUnminedSince := p.now().Add(-1 * p.mapExpiryTime)
		var offset int64

		for {
			// get all transactions since then chunk by chunk
			fmt.Println(loadUnminedLimit, offset)
			unminedTxs, err := p.store.GetUnmined(dbctx, getUnminedSince, loadUnminedLimit, offset)
			if err != nil {
				p.logger.Error("Failed to get unmined transactions", slog.String("err", err.Error()))
				continue
			}

			offset += loadUnminedLimit
			p.logger.Info("loaded unmined transactions", slog.Int("number", len(unminedTxs)), slog.Int64("limit", loadUnminedLimit))
			if len(unminedTxs) == 0 {
				break
			}

			p.logger.Info("Resending unmined transactions", slog.Int("number", len(unminedTxs)))
			for _, tx := range unminedTxs {
				// mark that we retried processing this transaction once more
				if err = p.store.IncrementRetries(dbctx, tx.Hash); err != nil {
					p.logger.Error("Failed to increment retries in database", slog.String("err", err.Error()))
				}

				if tx.Retries > MaxRetries {
					// Sending GETDATA to peers to see if they have it
					p.logger.Debug("Re-getting expired tx", slog.String("hash", tx.Hash.String()))
					p.pm.RequestTransaction(tx.Hash)
				} else {
					p.logger.Debug("Re-announcing expired tx", slog.String("hash", tx.Hash.String()))
					p.pm.AnnounceTransaction(tx.Hash, nil)
				}

				p.retries.AddDuration(time.Since(time.Now()))
			}
		}

	}
}

// GetPeers returns a list of connected and a list of disconnected peers
func (p *Processor) GetPeers() ([]string, []string) {
	peers := p.pm.GetPeers()
	peersConnected := make([]string, 0, len(peers))
	peersDisconnected := make([]string, 0, len(peers))
	for _, peer := range peers {
		if peer.Connected() {
			peersConnected = append(peersConnected, peer.String())
		} else {
			peersDisconnected = append(peersDisconnected, peer.String())
		}
	}

	return peersConnected, peersDisconnected
}

// func (p *Processor) LoadUnmined() {
// 	span, spanCtx := opentracing.StartSpanFromContext(context.Background(), "Processor:LoadUnmined")
// 	defer span.Finish()

// 	limit := loadUnminedLimit
// 	margin := p.maxMonitoredTxs - int64(len(p.ProcessorResponseMap.Items()))

// 	if margin < limit {
// 		limit = margin
// 	}

// 	if limit <= 0 {
// 		return
// 	}

// 	getUnminedSince := p.now().Add(-1 * p.mapExpiryTime)

// 	unminedTxs, err := p.store.GetUnmined(spanCtx, getUnminedSince, limit)
// 	if err != nil {
// 		p.logger.Error("Failed to get unmined transactions", slog.String("err", err.Error()))
// 		return
// 	}

// 	p.logger.Info("loaded unmined transactions", slog.Int("number", len(unminedTxs)), slog.Int64("limit", limit))

// 	if len(unminedTxs) == 0 {
// 		return
// 	}

// 	for _, record := range unminedTxs {
// 		pr := processor_response.NewProcessorResponseWithStatus(record.Hash, record.Status)
// 		pr.NoStats = true
// 		pr.Start = record.StoredAt
// 		p.ProcessorResponseMap.Set(record.Hash, pr)
// 	}
// }

var statusValueMap = map[metamorph_api.Status]int{
	metamorph_api.Status_UNKNOWN:                0,
	metamorph_api.Status_QUEUED:                 1,
	metamorph_api.Status_RECEIVED:               2,
	metamorph_api.Status_STORED:                 3,
	metamorph_api.Status_ANNOUNCED_TO_NETWORK:   4,
	metamorph_api.Status_REQUESTED_BY_NETWORK:   5,
	metamorph_api.Status_SENT_TO_NETWORK:        6,
	metamorph_api.Status_SEEN_IN_ORPHAN_MEMPOOL: 7,
	metamorph_api.Status_ACCEPTED_BY_NETWORK:    8,
	metamorph_api.Status_SEEN_ON_NETWORK:        9,
	metamorph_api.Status_REJECTED:               10,
	metamorph_api.Status_MINED:                  11,
	metamorph_api.Status_CONFIRMED:              12,
}

func (p *Processor) SendStatusForTransaction(hash *chainhash.Hash, status metamorph_api.Status, source string, statusErr error) error {
	span, _ := opentracing.StartSpanFromContext(context.Background(), "Processor:SendStatusForTransaction")
	defer span.Finish()

	// make sure we update the transaction status in database
	var rejectReason string
	if statusErr != nil {
		rejectReason = statusErr.Error()
	}

	p.storageStatusUpdateCh <- store.UpdateStatus{
		Hash:         *hash,
		Status:       status,
		RejectReason: rejectReason,
	}

	// if we recieve new update check if we have client connection waiting for status and send it
	processorResponse, ok := p.ProcessorResponseMap.Get(hash)
	if ok {
		processorResponse.UpdateStatus(&processor_response.ProcessorResponseStatusUpdate{
			Status:    status,
			StatusErr: nil,
		})
	}

	p.logger.Debug("Status reported for tx", slog.String("status", status.String()), slog.String("hash", hash.String()))
	return nil
}

func (p *Processor) ProcessTransaction(ctx context.Context, req *ProcessorRequest) {
	startNanos := time.Now().UnixNano()

	// we need to decouple the Context from the request, so that we don't get cancelled
	// when the request is cancelled
	callerSpan := opentracing.SpanFromContext(ctx)
	newctx := opentracing.ContextWithSpan(context.Background(), callerSpan)
	span, spanCtx := opentracing.StartSpanFromContext(newctx, "Processor:processTransaction")
	defer span.Finish()

	p.queuedCount.Add(1)

	// check if tx already stored, return it
	data, err := p.store.Get(ctx, req.Data.Hash[:])
	if err == nil {
		/*
			When transaction is re-submitted we make inserted_at_num to be now()
			to make sure it will be loaded and re-broadcasted if needed.
		*/
		insertedAtNum, _ := strconv.Atoi(p.now().Format("2006010215"))
		data.InsertedAtNum = insertedAtNum
		if set_err := p.store.Set(spanCtx, req.Data.Hash[:], data); set_err != nil {
			p.logger.Error("Failed to store transaction", slog.String("hash", req.Data.Hash.String()), slog.String("err", set_err.Error()))
		}

		var rejectErr error
		if data.RejectReason != "" {
			rejectErr = errors.New(data.RejectReason)
		}

		// notify the client instantly and return without waiting for any specific status
		req.ResponseChannel <- processor_response.StatusAndError{
			Hash:   data.Hash,
			Status: data.Status,
			Err:    rejectErr,
		}

		return
	}

	// register transaction in blocktx using message queue
	if err = p.mqClient.PublishRegisterTxs(req.Data.Hash[:]); err != nil {
		p.logger.Error("failed to register tx in blocktx", slog.String("hash", req.Data.Hash.String()), slog.String("err", err.Error()))
	}

	processorResponse := processor_response.NewProcessorResponseWithChannel(req.Data.Hash, req.ResponseChannel, req.Timeout)

	// store in database
	req.Data.Status = metamorph_api.Status_STORED
	insertedAtNum, _ := strconv.Atoi(p.now().Format("2006010215"))
	req.Data.InsertedAtNum = insertedAtNum
	err = p.store.Set(spanCtx, req.Data.Hash[:], req.Data)
	if err != nil {
		p.logger.Error("Failed to store transaction", slog.String("hash", req.Data.Hash.String()), slog.String("err", err.Error()))
	}

	// broadcast that transaction is stored to client
	processorResponse.UpdateStatus(&processor_response.ProcessorResponseStatusUpdate{
		Status:    metamorph_api.Status_STORED,
		StatusErr: err,
	})

	// Add this transaction to the map of transactions that client is listening to with open connection.
	p.ProcessorResponseMap.Set(req.Data.Hash, processorResponse)

	// we no longer need processor response object after client disconnects
	go func() {
		time.Sleep(req.Timeout + time.Second)
		p.ProcessorResponseMap.Delete(req.Data.Hash)
	}()

	// Announce transaction to network and save peers
	p.logger.Debug("announcing transaction", slog.String("hash", req.Data.Hash.String()))
	p.pm.AnnounceTransaction(req.Data.Hash, nil)

	// notify existing client about new status
	processorResponse, ok := p.ProcessorResponseMap.Get(req.Data.Hash)
	if ok {
		processorResponse.UpdateStatus(&processor_response.ProcessorResponseStatusUpdate{
			Status:    metamorph_api.Status_ANNOUNCED_TO_NETWORK,
			StatusErr: nil,
		})
	}

	// broadcast that transaction is announced to network (eventually active clientს catch that)
	p.storageStatusUpdateCh <- store.UpdateStatus{
		Hash:         *req.Data.Hash,
		Status:       metamorph_api.Status_ANNOUNCED_TO_NETWORK,
		RejectReason: "",
	}

	gocore.NewStat("processor").AddTime(startNanos)
}

func (p *Processor) Health() error {
	healthyConnections := 0

	for _, peer := range p.pm.GetPeers() {
		if peer.Connected() && peer.IsHealthy() {
			healthyConnections++
		}
	}

	if healthyConnections < minimumHealthyConnections {
		p.logger.Warn("Less than expected healthy peers - ", slog.Int("number", healthyConnections))
		return nil
	}

	if healthyConnections == 0 {
		p.logger.Error("Metamorph not healthy")
		return ErrUnhealthy
	}

	return nil
}
