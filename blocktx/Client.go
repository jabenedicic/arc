package blocktx

import (
	"context"
	"time"

	"github.com/TAAL-GmbH/arc/blocktx/blocktx_api"
	"github.com/TAAL-GmbH/arc/tracing"
	"github.com/ordishs/go-utils"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ClientI is the interface for the block-tx transactionHandler
type ClientI interface {
	Start(minedBlockChan chan *blocktx_api.Block)
	LocateTransaction(ctx context.Context, transaction *blocktx_api.Transaction) (string, error)
	RegisterTransaction(ctx context.Context, transaction *blocktx_api.TransactionAndSource) error
	GetBlock(ctx context.Context, blockHash []byte) (*blocktx_api.Block, error)
	GetMinedTransactionsForBlock(ctx context.Context, blockAndSource *blocktx_api.BlockAndSource) (*blocktx_api.MinedTransactions, error)
}

type Client struct {
	address string
	logger  utils.Logger
	conn    *grpc.ClientConn
	client  blocktx_api.BlockTxAPIClient
}

func NewClient(l utils.Logger, address string) ClientI {
	btc := &Client{
		address: address,
		logger:  l,
	}

	conn, err := btc.dialGRPC()
	if err != nil {
		btc.logger.Fatalf("Failed to connect to block-tx server at %s: %v", btc.address, err)
	}
	btc.logger.Infof("Connected to block-tx server at %s", btc.address)

	btc.conn = conn
	btc.client = blocktx_api.NewBlockTxAPIClient(conn)

	return btc
}

func (btc *Client) Start(minedBlockChan chan *blocktx_api.Block) {
	for {
		stream, err := btc.client.GetBlockNotificationStream(context.Background(), &blocktx_api.Height{})
		if err != nil {
			btc.logger.Errorf("Error getting block notification stream: %v", err)
		} else {
			btc.logger.Infof("Connected to block-tx server at %s", btc.address)

			var block *blocktx_api.Block
			for {
				block, err = stream.Recv()
				if err != nil {
					break
				}

				btc.logger.Infof("Block %s\n", utils.HexEncodeAndReverseBytes(block.Hash))
				utils.SafeSend(minedBlockChan, block)
			}
		}

		time.Sleep(10 * time.Second)
	}
}

func (btc *Client) LocateTransaction(ctx context.Context, transaction *blocktx_api.Transaction) (string, error) {
	location, err := btc.client.LocateTransaction(ctx, transaction)
	if err != nil {
		return "", err
	}

	return location.Source, nil
}

func (btc *Client) RegisterTransaction(ctx context.Context, transaction *blocktx_api.TransactionAndSource) error {
	if _, err := btc.client.RegisterTransaction(ctx, transaction); err != nil {
		return err
	}

	return nil
}

func (btc *Client) GetBlock(ctx context.Context, blockHash []byte) (*blocktx_api.Block, error) {
	block, err := btc.client.GetBlock(ctx, &blocktx_api.Hash{Hash: blockHash})
	if err != nil {
		return nil, err
	}

	return block, nil
}

func (btc *Client) GetMinedTransactionsForBlock(ctx context.Context, blockAndSource *blocktx_api.BlockAndSource) (*blocktx_api.MinedTransactions, error) {
	mt, err := btc.client.GetMinedTransactionsForBlock(ctx, blockAndSource)
	if err != nil {
		return nil, err
	}

	return mt, nil
}

func (btc *Client) dialGRPC() (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`), // This sets the initial balancing policy.
	}

	return grpc.Dial(btc.address, tracing.AddGRPCDialOptions(opts)...)
}
