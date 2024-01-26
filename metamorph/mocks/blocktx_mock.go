// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"context"
	"github.com/bitcoin-sv/arc/blocktx"
	"github.com/bitcoin-sv/arc/blocktx/blocktx_api"
	"sync"
)

// Ensure, that ClientIMock does implement blocktx.ClientI.
// If this is not the case, regenerate this file with moq.
var _ blocktx.ClientI = &ClientIMock{}

// ClientIMock is a mock implementation of blocktx.ClientI.
//
//	func TestSomethingThatUsesClientI(t *testing.T) {
//
//		// make and configure a mocked blocktx.ClientI
//		mockedClientI := &ClientIMock{
//			GetTransactionBlocksFunc: func(ctx context.Context, transaction *blocktx_api.Transactions) (*blocktx_api.TransactionBlocks, error) {
//				panic("mock out the GetTransactionBlocks method")
//			},
//			GetTransactionMerklePathFunc: func(ctx context.Context, transaction *blocktx_api.Transaction) (string, error) {
//				panic("mock out the GetTransactionMerklePath method")
//			},
//			HealthFunc: func(ctx context.Context) error {
//				panic("mock out the Health method")
//			},
//			RegisterTransactionFunc: func(ctx context.Context, transaction *blocktx_api.TransactionAndSource) error {
//				panic("mock out the RegisterTransaction method")
//			},
//		}
//
//		// use mockedClientI in code that requires blocktx.ClientI
//		// and then make assertions.
//
//	}
type ClientIMock struct {
	// GetTransactionBlocksFunc mocks the GetTransactionBlocks method.
	GetTransactionBlocksFunc func(ctx context.Context, transaction *blocktx_api.Transactions) (*blocktx_api.TransactionBlocks, error)

	// GetTransactionMerklePathFunc mocks the GetTransactionMerklePath method.
	GetTransactionMerklePathFunc func(ctx context.Context, transaction *blocktx_api.Transaction) (string, error)

	// HealthFunc mocks the Health method.
	HealthFunc func(ctx context.Context) error

	// RegisterTransactionFunc mocks the RegisterTransaction method.
	RegisterTransactionFunc func(ctx context.Context, transaction *blocktx_api.TransactionAndSource) error

	// calls tracks calls to the methods.
	calls struct {
		// GetTransactionBlocks holds details about calls to the GetTransactionBlocks method.
		GetTransactionBlocks []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Transaction is the transaction argument value.
			Transaction *blocktx_api.Transactions
		}
		// GetTransactionMerklePath holds details about calls to the GetTransactionMerklePath method.
		GetTransactionMerklePath []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Transaction is the transaction argument value.
			Transaction *blocktx_api.Transaction
		}
		// Health holds details about calls to the Health method.
		Health []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// RegisterTransaction holds details about calls to the RegisterTransaction method.
		RegisterTransaction []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Transaction is the transaction argument value.
			Transaction *blocktx_api.TransactionAndSource
		}
	}
	lockGetTransactionBlocks     sync.RWMutex
	lockGetTransactionMerklePath sync.RWMutex
	lockHealth                   sync.RWMutex
	lockRegisterTransaction      sync.RWMutex
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
//
//	len(mockedClientI.GetTransactionBlocksCalls())
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

// GetTransactionMerklePath calls GetTransactionMerklePathFunc.
func (mock *ClientIMock) GetTransactionMerklePath(ctx context.Context, transaction *blocktx_api.Transaction) (string, error) {
	if mock.GetTransactionMerklePathFunc == nil {
		panic("ClientIMock.GetTransactionMerklePathFunc: method is nil but ClientI.GetTransactionMerklePath was just called")
	}
	callInfo := struct {
		Ctx         context.Context
		Transaction *blocktx_api.Transaction
	}{
		Ctx:         ctx,
		Transaction: transaction,
	}
	mock.lockGetTransactionMerklePath.Lock()
	mock.calls.GetTransactionMerklePath = append(mock.calls.GetTransactionMerklePath, callInfo)
	mock.lockGetTransactionMerklePath.Unlock()
	return mock.GetTransactionMerklePathFunc(ctx, transaction)
}

// GetTransactionMerklePathCalls gets all the calls that were made to GetTransactionMerklePath.
// Check the length with:
//
//	len(mockedClientI.GetTransactionMerklePathCalls())
func (mock *ClientIMock) GetTransactionMerklePathCalls() []struct {
	Ctx         context.Context
	Transaction *blocktx_api.Transaction
} {
	var calls []struct {
		Ctx         context.Context
		Transaction *blocktx_api.Transaction
	}
	mock.lockGetTransactionMerklePath.RLock()
	calls = mock.calls.GetTransactionMerklePath
	mock.lockGetTransactionMerklePath.RUnlock()
	return calls
}

// Health calls HealthFunc.
func (mock *ClientIMock) Health(ctx context.Context) error {
	if mock.HealthFunc == nil {
		panic("ClientIMock.HealthFunc: method is nil but ClientI.Health was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockHealth.Lock()
	mock.calls.Health = append(mock.calls.Health, callInfo)
	mock.lockHealth.Unlock()
	return mock.HealthFunc(ctx)
}

// HealthCalls gets all the calls that were made to Health.
// Check the length with:
//
//	len(mockedClientI.HealthCalls())
func (mock *ClientIMock) HealthCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockHealth.RLock()
	calls = mock.calls.Health
	mock.lockHealth.RUnlock()
	return calls
}

// RegisterTransaction calls RegisterTransactionFunc.
func (mock *ClientIMock) RegisterTransaction(ctx context.Context, transaction *blocktx_api.TransactionAndSource) error {
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
//
//	len(mockedClientI.RegisterTransactionCalls())
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
