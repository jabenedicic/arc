// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"github.com/bitcoin-sv/arc/internal/callbacker"
	"sync"
)

// Ensure, that SenderIMock does implement callbacker.SenderI.
// If this is not the case, regenerate this file with moq.
var _ callbacker.SenderI = &SenderIMock{}

// SenderIMock is a mock implementation of callbacker.SenderI.
//
//	func TestSomethingThatUsesSenderI(t *testing.T) {
//
//		// make and configure a mocked callbacker.SenderI
//		mockedSenderI := &SenderIMock{
//			SendFunc: func(url string, token string, callback *callbacker.Callback) (bool, bool) {
//				panic("mock out the Send method")
//			},
//			SendBatchFunc: func(url string, token string, callbacks []*callbacker.Callback) (bool, bool) {
//				panic("mock out the SendBatch method")
//			},
//		}
//
//		// use mockedSenderI in code that requires callbacker.SenderI
//		// and then make assertions.
//
//	}
type SenderIMock struct {
	// SendFunc mocks the Send method.
	SendFunc func(url string, token string, callback *callbacker.Callback) (bool, bool)

	// SendBatchFunc mocks the SendBatch method.
	SendBatchFunc func(url string, token string, callbacks []*callbacker.Callback) (bool, bool)

	// calls tracks calls to the methods.
	calls struct {
		// Send holds details about calls to the Send method.
		Send []struct {
			// URL is the url argument value.
			URL string
			// Token is the token argument value.
			Token string
			// Callback is the callback argument value.
			Callback *callbacker.Callback
		}
		// SendBatch holds details about calls to the SendBatch method.
		SendBatch []struct {
			// URL is the url argument value.
			URL string
			// Token is the token argument value.
			Token string
			// Callbacks is the callbacks argument value.
			Callbacks []*callbacker.Callback
		}
	}
	lockSend      sync.RWMutex
	lockSendBatch sync.RWMutex
}

// Send calls SendFunc.
func (mock *SenderIMock) Send(url string, token string, callback *callbacker.Callback) (bool, bool) {
	if mock.SendFunc == nil {
		panic("SenderIMock.SendFunc: method is nil but SenderI.Send was just called")
	}
	callInfo := struct {
		URL      string
		Token    string
		Callback *callbacker.Callback
	}{
		URL:      url,
		Token:    token,
		Callback: callback,
	}
	mock.lockSend.Lock()
	mock.calls.Send = append(mock.calls.Send, callInfo)
	mock.lockSend.Unlock()
	return mock.SendFunc(url, token, callback)
}

// SendCalls gets all the calls that were made to Send.
// Check the length with:
//
//	len(mockedSenderI.SendCalls())
func (mock *SenderIMock) SendCalls() []struct {
	URL      string
	Token    string
	Callback *callbacker.Callback
} {
	var calls []struct {
		URL      string
		Token    string
		Callback *callbacker.Callback
	}
	mock.lockSend.RLock()
	calls = mock.calls.Send
	mock.lockSend.RUnlock()
	return calls
}

// SendBatch calls SendBatchFunc.
func (mock *SenderIMock) SendBatch(url string, token string, callbacks []*callbacker.Callback) (bool, bool) {
	if mock.SendBatchFunc == nil {
		panic("SenderIMock.SendBatchFunc: method is nil but SenderI.SendBatch was just called")
	}
	callInfo := struct {
		URL       string
		Token     string
		Callbacks []*callbacker.Callback
	}{
		URL:       url,
		Token:     token,
		Callbacks: callbacks,
	}
	mock.lockSendBatch.Lock()
	mock.calls.SendBatch = append(mock.calls.SendBatch, callInfo)
	mock.lockSendBatch.Unlock()
	return mock.SendBatchFunc(url, token, callbacks)
}

// SendBatchCalls gets all the calls that were made to SendBatch.
// Check the length with:
//
//	len(mockedSenderI.SendBatchCalls())
func (mock *SenderIMock) SendBatchCalls() []struct {
	URL       string
	Token     string
	Callbacks []*callbacker.Callback
} {
	var calls []struct {
		URL       string
		Token     string
		Callbacks []*callbacker.Callback
	}
	mock.lockSendBatch.RLock()
	calls = mock.calls.SendBatch
	mock.lockSendBatch.RUnlock()
	return calls
}
