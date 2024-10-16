package broadcaster_test

import (
	"context"
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/bitcoin-sv/arc/internal/broadcaster"
	"github.com/bitcoin-sv/arc/internal/broadcaster/mocks"
	"github.com/bitcoin-sv/arc/internal/metamorph/metamorph_api"
	"github.com/bitcoin-sv/arc/internal/testdata"
	"github.com/bitcoin-sv/arc/pkg/keyset"
	"github.com/bitcoin-sv/go-sdk/script"
	sdkTx "github.com/bitcoin-sv/go-sdk/transaction"
	"github.com/stretchr/testify/require"
)

func TestUTXOCreator(t *testing.T) {
	ks, err := keyset.New()
	require.NoError(t, err)

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	tests := []struct {
		name                   string
		getUTXOsResp           sdkTx.UTXOs
		getBalanceFunc         func(ctx context.Context, address string, constantBackoff time.Duration, retries uint64) (int64, int64, error)
		expectedBroadcastCalls int
		expectedError          error
		requestedUTXOs         int
		requestedAmountPerUTXO uint64
	}{
		{
			name:         "success - creates correct UTXOs",
			getUTXOsResp: sdkTx.UTXOs{{TxID: testdata.TX1Hash[:], Vout: 0, LockingScript: ks.Script, Satoshis: 401}},
			getBalanceFunc: func(ctx context.Context, address string, constantBackoff time.Duration, retries uint64) (int64, int64, error) {
				return 400, 0, nil
			},
			expectedBroadcastCalls: 1,
			expectedError:          nil,
			requestedUTXOs:         4,
			requestedAmountPerUTXO: 100,
		},
		{
			name:         "Insufficient balance",
			getUTXOsResp: sdkTx.UTXOs{{TxID: testdata.TX1Hash[:], Vout: 0, LockingScript: ks.Script, Satoshis: 50}},
			getBalanceFunc: func(ctx context.Context, address string, constantBackoff time.Duration, retries uint64) (int64, int64, error) {
				return 100, 0, nil
			},
			expectedBroadcastCalls: 0,
			expectedError:          broadcaster.ErrRequestedSatoshisTooHigh,
			requestedUTXOs:         4,
			requestedAmountPerUTXO: 100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Given
			mockArcClient := &mocks.ArcClientMock{
				BroadcastTransactionsFunc: func(ctx context.Context, txs sdkTx.Transactions, waitForStatus metamorph_api.Status, callbackURL, callbackToken string, fullStatusUpdates, skipFeeValidation bool) ([]*metamorph_api.TransactionStatus, error) {
					statuses := make([]*metamorph_api.TransactionStatus, len(txs))
					for i, tx := range txs {
						statuses[i] = &metamorph_api.TransactionStatus{Txid: tx.TxID(), Status: metamorph_api.Status_SEEN_ON_NETWORK}
					}
					return statuses, nil
				},
			}

			mockUtxoClient := &mocks.UtxoClientMock{
				GetBalanceWithRetriesFunc: tt.getBalanceFunc,
				GetUTXOsWithRetriesFunc: func(ctx context.Context, lockingScript *script.Script, address string, constantBackoff time.Duration, retries uint64) (sdkTx.UTXOs, error) {
					return tt.getUTXOsResp, nil
				},
			}

			utxoCreator, err := broadcaster.NewUTXOCreator(logger, mockArcClient, ks, mockUtxoClient, false)
			require.NoError(t, err)
			defer utxoCreator.Shutdown()

			// When
			actualError := utxoCreator.Start(tt.requestedUTXOs, tt.requestedAmountPerUTXO)

			// Then
			if tt.expectedError != nil {
				require.ErrorIs(t, actualError, tt.expectedError)
				return
			} else {
				require.NoError(t, actualError)
			}

			time.Sleep(500 * time.Millisecond)

			require.Equal(t, tt.expectedBroadcastCalls, len(mockArcClient.BroadcastTransactionsCalls()), "Unexpected number of BroadcastTransactions calls")

			if tt.expectedBroadcastCalls > 0 {
				for _, call := range mockArcClient.BroadcastTransactionsCalls() {
					for _, tx := range call.Txs {
						for _, output := range tx.Outputs {
							require.Equal(t, tt.requestedAmountPerUTXO, output.Satoshis, "Each UTXO output should have the requested amount")
						}
					}
				}
			}
		})
	}
}
