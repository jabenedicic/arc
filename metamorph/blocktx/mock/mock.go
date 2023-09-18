// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	"github.com/bitcoin-sv/arc/blocktx"
	"github.com/bitcoin-sv/arc/blocktx/blocktx_api"
	"github.com/libsv/go-p2p/chaincfg/chainhash"
	"sync"
)

// Ensure, that ClientIMock does implement blocktx.ClientI.
// If this is not the case, regenerate this file with moq.
var _ blocktx.ClientI = &ClientIMock{}

// ClientIMock is a mock implementation of blocktx.ClientI.
//
// 	func TestSomethingThatUsesClientI(t *testing.T) {
//
// 		// make and configure a mocked blocktx.ClientI
// 		mockedClientI := &ClientIMock{
// 			GetBlockFunc: func(ctx context.Context, blockHash *chainhash.Hash) (*blocktx_api.Block, error) {
// 				panic("mock out the GetBlock method")
// 			},
// 			GetLastProcessedBlockFunc: func(ctx context.Context) (*blocktx_api.Block, error) {
// 				panic("mock out the GetLastProcessedBlock method")
// 			},
// 			GetMinedTransactionsForBlockFunc: func(ctx context.Context, blockAndSource *blocktx_api.BlockAndSource) (*blocktx_api.MinedTransactions, error) {
// 				panic("mock out the GetMinedTransactionsForBlock method")
// 			},
// 			GetTransactionBlockFunc: func(ctx context.Context, transaction *blocktx_api.Transaction) (*blocktx_api.RegisterTransactionResponse, error) {
// 				panic("mock out the GetTransactionBlock method")
// 			},
// 			GetTransactionBlocksFunc: func(ctx context.Context, transaction *blocktx_api.Transactions) (*blocktx_api.TransactionBlocks, error) {
// 				panic("mock out the GetTransactionBlocks method")
// 			},
// 			LocateTransactionFunc: func(ctx context.Context, transaction *blocktx_api.Transaction) (string, error) {
// 				panic("mock out the LocateTransaction method")
// 			},
// 			RegisterTransactionFunc: func(ctx context.Context, transaction *blocktx_api.TransactionAndSource) (*blocktx_api.RegisterTransactionResponse, error) {
// 				panic("mock out the RegisterTransaction method")
// 			},
// 			StartFunc: func(minedBlockChan chan *blocktx_api.Block)  {
// 				panic("mock out the Start method")
// 			},
// 		}
//
// 		// use mockedClientI in code that requires blocktx.ClientI
// 		// and then make assertions.
//
// 	}
type ClientIMock struct {
	// GetBlockFunc mocks the GetBlock method.
	GetBlockFunc func(ctx context.Context, blockHash *chainhash.Hash) (*blocktx_api.Block, error)

	// GetLastProcessedBlockFunc mocks the GetLastProcessedBlock method.
	GetLastProcessedBlockFunc func(ctx context.Context) (*blocktx_api.Block, error)

	// GetTransactionMerklePath returns merkle path of the transaction
	GetTransactionMerklePath(ctx context.Context, transaction *blocktx_api.Transaction) (string, error)

	// GetMinedTransactionsForBlockFunc mocks the GetMinedTransactionsForBlock method.
	GetMinedTransactionsForBlockFunc func(ctx context.Context, blockAndSource *blocktx_api.BlockAndSource) (*blocktx_api.MinedTransactions, error)

	// GetTransactionBlockFunc mocks the GetTransactionBlock method.
	GetTransactionBlockFunc func(ctx context.Context, transaction *blocktx_api.Transaction) (*blocktx_api.RegisterTransactionResponse, error)

	// GetTransactionBlocksFunc mocks the GetTransactionBlocks method.
	GetTransactionBlocksFunc func(ctx context.Context, transaction *blocktx_api.Transactions) (*blocktx_api.TransactionBlocks, error)

	// LocateTransactionFunc mocks the LocateTransaction method.
	LocateTransactionFunc func(ctx context.Context, transaction *blocktx_api.Transaction) (string, error)

	// RegisterTransactionFunc mocks the RegisterTransaction method.
	RegisterTransactionFunc func(ctx context.Context, transaction *blocktx_api.TransactionAndSource) (*blocktx_api.RegisterTransactionResponse, error)

	// StartFunc mocks the Start method.
	StartFunc func(minedBlockChan chan *blocktx_api.Block)

	// calls tracks calls to the methods.
	calls struct {
		// GetBlock holds details about calls to the GetBlock method.
		GetBlock []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// BlockHash is the blockHash argument value.
			BlockHash *chainhash.Hash
		}
		// GetLastProcessedBlock holds details about calls to the GetLastProcessedBlock method.
		GetLastProcessedBlock []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// GetMinedTransactionsForBlock holds details about calls to the GetMinedTransactionsForBlock method.
		GetMinedTransactionsForBlock []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// BlockAndSource is the blockAndSource argument value.
			BlockAndSource *blocktx_api.BlockAndSource
		}
		// GetTransactionBlock holds details about calls to the GetTransactionBlock method.
		GetTransactionBlock []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Transaction is the transaction argument value.
			Transaction *blocktx_api.Transaction
		}
		// GetTransactionBlocks holds details about calls to the GetTransactionBlocks method.
		GetTransactionBlocks []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Transaction is the transaction argument value.
			Transaction *blocktx_api.Transactions
		}
		// LocateTransaction holds details about calls to the LocateTransaction method.
		LocateTransaction []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Transaction is the transaction argument value.
			Transaction *blocktx_api.Transaction
		}
		// RegisterTransaction holds details about calls to the RegisterTransaction method.
		RegisterTransaction []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Transaction is the transaction argument value.
			Transaction *blocktx_api.TransactionAndSource
		}
		// Start holds details about calls to the Start method.
		Start []struct {
			// MinedBlockChan is the minedBlockChan argument value.
			MinedBlockChan chan *blocktx_api.Block
		}
	}
	lockGetBlock                     sync.RWMutex
	lockGetLastProcessedBlock        sync.RWMutex
	lockGetMinedTransactionsForBlock sync.RWMutex
	lockGetTransactionBlock          sync.RWMutex
	lockGetTransactionBlocks         sync.RWMutex
	lockLocateTransaction            sync.RWMutex
	lockRegisterTransaction          sync.RWMutex
	lockStart                        sync.RWMutex
}

// GetBlock calls GetBlockFunc.
func (mock *ClientIMock) GetBlock(ctx context.Context, blockHash *chainhash.Hash) (*blocktx_api.Block, error) {
	if mock.GetBlockFunc == nil {
		panic("ClientIMock.GetBlockFunc: method is nil but ClientI.GetBlock was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		BlockHash *chainhash.Hash
	}{
		Ctx:       ctx,
		BlockHash: blockHash,
	}
	mock.lockGetBlock.Lock()
	mock.calls.GetBlock = append(mock.calls.GetBlock, callInfo)
	mock.lockGetBlock.Unlock()
	return mock.GetBlockFunc(ctx, blockHash)
}

// GetBlockCalls gets all the calls that were made to GetBlock.
// Check the length with:
//     len(mockedClientI.GetBlockCalls())
func (mock *ClientIMock) GetBlockCalls() []struct {
	Ctx       context.Context
	BlockHash *chainhash.Hash
} {
	var calls []struct {
		Ctx       context.Context
		BlockHash *chainhash.Hash
	}
	mock.lockGetBlock.RLock()
	calls = mock.calls.GetBlock
	mock.lockGetBlock.RUnlock()
	return calls
}

// GetLastProcessedBlock calls GetLastProcessedBlockFunc.
func (mock *ClientIMock) GetLastProcessedBlock(ctx context.Context) (*blocktx_api.Block, error) {
	if mock.GetLastProcessedBlockFunc == nil {
		panic("ClientIMock.GetLastProcessedBlockFunc: method is nil but ClientI.GetLastProcessedBlock was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetLastProcessedBlock.Lock()
	mock.calls.GetLastProcessedBlock = append(mock.calls.GetLastProcessedBlock, callInfo)
	mock.lockGetLastProcessedBlock.Unlock()
	return mock.GetLastProcessedBlockFunc(ctx)
}

// GetLastProcessedBlockCalls gets all the calls that were made to GetLastProcessedBlock.
// Check the length with:
//     len(mockedClientI.GetLastProcessedBlockCalls())
func (mock *ClientIMock) GetLastProcessedBlockCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockGetLastProcessedBlock.RLock()
	calls = mock.calls.GetLastProcessedBlock
	mock.lockGetLastProcessedBlock.RUnlock()
	return calls
}

// GetMinedTransactionsForBlock calls GetMinedTransactionsForBlockFunc.
func (mock *ClientIMock) GetMinedTransactionsForBlock(ctx context.Context, blockAndSource *blocktx_api.BlockAndSource) (*blocktx_api.MinedTransactions, error) {
	if mock.GetMinedTransactionsForBlockFunc == nil {
		panic("ClientIMock.GetMinedTransactionsForBlockFunc: method is nil but ClientI.GetMinedTransactionsForBlock was just called")
	}
	callInfo := struct {
		Ctx            context.Context
		BlockAndSource *blocktx_api.BlockAndSource
	}{
		Ctx:            ctx,
		BlockAndSource: blockAndSource,
	}
	mock.lockGetMinedTransactionsForBlock.Lock()
	mock.calls.GetMinedTransactionsForBlock = append(mock.calls.GetMinedTransactionsForBlock, callInfo)
	mock.lockGetMinedTransactionsForBlock.Unlock()
	return mock.GetMinedTransactionsForBlockFunc(ctx, blockAndSource)
}

// GetMinedTransactionsForBlockCalls gets all the calls that were made to GetMinedTransactionsForBlock.
// Check the length with:
//     len(mockedClientI.GetMinedTransactionsForBlockCalls())
func (mock *ClientIMock) GetMinedTransactionsForBlockCalls() []struct {
	Ctx            context.Context
	BlockAndSource *blocktx_api.BlockAndSource
} {
	var calls []struct {
		Ctx            context.Context
		BlockAndSource *blocktx_api.BlockAndSource
	}
	mock.lockGetMinedTransactionsForBlock.RLock()
	calls = mock.calls.GetMinedTransactionsForBlock
	mock.lockGetMinedTransactionsForBlock.RUnlock()
	return calls
}

// GetTransactionBlock calls GetTransactionBlockFunc.
func (mock *ClientIMock) GetTransactionBlock(ctx context.Context, transaction *blocktx_api.Transaction) (*blocktx_api.RegisterTransactionResponse, error) {
	if mock.GetTransactionBlockFunc == nil {
		panic("ClientIMock.GetTransactionBlockFunc: method is nil but ClientI.GetTransactionBlock was just called")
	}
	callInfo := struct {
		Ctx         context.Context
		Transaction *blocktx_api.Transaction
	}{
		Ctx:         ctx,
		Transaction: transaction,
	}
	mock.lockGetTransactionBlock.Lock()
	mock.calls.GetTransactionBlock = append(mock.calls.GetTransactionBlock, callInfo)
	mock.lockGetTransactionBlock.Unlock()
	return mock.GetTransactionBlockFunc(ctx, transaction)
}

// GetTransactionBlockCalls gets all the calls that were made to GetTransactionBlock.
// Check the length with:
//     len(mockedClientI.GetTransactionBlockCalls())
func (mock *ClientIMock) GetTransactionBlockCalls() []struct {
	Ctx         context.Context
	Transaction *blocktx_api.Transaction
} {
	var calls []struct {
		Ctx         context.Context
		Transaction *blocktx_api.Transaction
	}
	mock.lockGetTransactionBlock.RLock()
	calls = mock.calls.GetTransactionBlock
	mock.lockGetTransactionBlock.RUnlock()
	return calls
}

// GetTransactionBlocks calls GetTransactionBlocksFunc.
func (mock *ClientIMock) GetTransactionBlocks(ctx context.Context, transaction *blocktx_api.Transactions) (*blocktx_api.TransactionBlocks, error) {
	if mock.GetTransactionBlocksFunc == nil {
		panic("ClientIMock.GetTransactionBlocksFunc: method is nil but ClientI.GetTransactionBlocks was just called")
	}
	callInfo := struct {
		Ctx         context.Context
		Transaction *blocktx_api.Transactions
	}{
		Ctx:         ctx,
		Transaction: transaction,
	}
	mock.lockGetTransactionBlocks.Lock()
	mock.calls.GetTransactionBlocks = append(mock.calls.GetTransactionBlocks, callInfo)
	mock.lockGetTransactionBlocks.Unlock()
	return mock.GetTransactionBlocksFunc(ctx, transaction)
}

// GetTransactionBlocksCalls gets all the calls that were made to GetTransactionBlocks.
// Check the length with:
//     len(mockedClientI.GetTransactionBlocksCalls())
func (mock *ClientIMock) GetTransactionBlocksCalls() []struct {
	Ctx         context.Context
	Transaction *blocktx_api.Transactions
} {
	var calls []struct {
		Ctx         context.Context
		Transaction *blocktx_api.Transactions
	}
	mock.lockGetTransactionBlocks.RLock()
	calls = mock.calls.GetTransactionBlocks
	mock.lockGetTransactionBlocks.RUnlock()
	return calls
}

// LocateTransaction calls LocateTransactionFunc.
func (mock *ClientIMock) LocateTransaction(ctx context.Context, transaction *blocktx_api.Transaction) (string, error) {
	if mock.LocateTransactionFunc == nil {
		panic("ClientIMock.LocateTransactionFunc: method is nil but ClientI.LocateTransaction was just called")
	}
	callInfo := struct {
		Ctx         context.Context
		Transaction *blocktx_api.Transaction
	}{
		Ctx:         ctx,
		Transaction: transaction,
	}
	mock.lockLocateTransaction.Lock()
	mock.calls.LocateTransaction = append(mock.calls.LocateTransaction, callInfo)
	mock.lockLocateTransaction.Unlock()
	return mock.LocateTransactionFunc(ctx, transaction)
}

// LocateTransactionCalls gets all the calls that were made to LocateTransaction.
// Check the length with:
//     len(mockedClientI.LocateTransactionCalls())
func (mock *ClientIMock) LocateTransactionCalls() []struct {
	Ctx         context.Context
	Transaction *blocktx_api.Transaction
} {
	var calls []struct {
		Ctx         context.Context
		Transaction *blocktx_api.Transaction
	}
	mock.lockLocateTransaction.RLock()
	calls = mock.calls.LocateTransaction
	mock.lockLocateTransaction.RUnlock()
	return calls
}

// RegisterTransaction calls RegisterTransactionFunc.
func (mock *ClientIMock) RegisterTransaction(ctx context.Context, transaction *blocktx_api.TransactionAndSource) (*blocktx_api.RegisterTransactionResponse, error) {
	if mock.RegisterTransactionFunc == nil {
		panic("ClientIMock.RegisterTransactionFunc: method is nil but ClientI.RegisterTransaction was just called")
	}
	callInfo := struct {
		Ctx         context.Context
		Transaction *blocktx_api.TransactionAndSource
	}{
		Ctx:         ctx,
		Transaction: transaction,
	}
	mock.lockRegisterTransaction.Lock()
	mock.calls.RegisterTransaction = append(mock.calls.RegisterTransaction, callInfo)
	mock.lockRegisterTransaction.Unlock()
	return mock.RegisterTransactionFunc(ctx, transaction)
}

// RegisterTransactionCalls gets all the calls that were made to RegisterTransaction.
// Check the length with:
//     len(mockedClientI.RegisterTransactionCalls())
func (mock *ClientIMock) RegisterTransactionCalls() []struct {
	Ctx         context.Context
	Transaction *blocktx_api.TransactionAndSource
} {
	var calls []struct {
		Ctx         context.Context
		Transaction *blocktx_api.TransactionAndSource
	}
	mock.lockRegisterTransaction.RLock()
	calls = mock.calls.RegisterTransaction
	mock.lockRegisterTransaction.RUnlock()
	return calls
}

// Start calls StartFunc.
func (mock *ClientIMock) Start(minedBlockChan chan *blocktx_api.Block) {
	if mock.StartFunc == nil {
		panic("ClientIMock.StartFunc: method is nil but ClientI.Start was just called")
	}
	callInfo := struct {
		MinedBlockChan chan *blocktx_api.Block
	}{
		MinedBlockChan: minedBlockChan,
	}
	mock.lockStart.Lock()
	mock.calls.Start = append(mock.calls.Start, callInfo)
	mock.lockStart.Unlock()
	mock.StartFunc(minedBlockChan)
}

// StartCalls gets all the calls that were made to Start.
// Check the length with:
//     len(mockedClientI.StartCalls())
func (mock *ClientIMock) StartCalls() []struct {
	MinedBlockChan chan *blocktx_api.Block
} {
	var calls []struct {
		MinedBlockChan chan *blocktx_api.Block
	}
	mock.lockStart.RLock()
	calls = mock.calls.Start
	mock.lockStart.RUnlock()
	return calls
}
