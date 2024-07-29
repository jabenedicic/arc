// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"github.com/bitcoin-sv/arc/internal/blocktx"
	"google.golang.org/protobuf/reflect/protoreflect"
	"sync"
)

// Ensure, that MessageQueueClientMock does implement blocktx.MessageQueueClient.
// If this is not the case, regenerate this file with moq.
var _ blocktx.MessageQueueClient = &MessageQueueClientMock{}

// MessageQueueClientMock is a mock implementation of blocktx.MessageQueueClient.
//
//	func TestSomethingThatUsesMessageQueueClient(t *testing.T) {
//
//		// make and configure a mocked blocktx.MessageQueueClient
//		mockedMessageQueueClient := &MessageQueueClientMock{
//			PublishMarshalFunc: func(topic string, m protoreflect.ProtoMessage) error {
//				panic("mock out the PublishMarshal method")
//			},
//			SubscribeFunc: func(topic string, msgFunc func([]byte) error) error {
//				panic("mock out the Subscribe method")
//			},
//		}
//
//		// use mockedMessageQueueClient in code that requires blocktx.MessageQueueClient
//		// and then make assertions.
//
//	}
type MessageQueueClientMock struct {
	// PublishMarshalFunc mocks the PublishMarshal method.
	PublishMarshalFunc func(topic string, m protoreflect.ProtoMessage) error

	// SubscribeFunc mocks the Subscribe method.
	SubscribeFunc func(topic string, msgFunc func([]byte) error) error

	// calls tracks calls to the methods.
	calls struct {
		// PublishMarshal holds details about calls to the PublishMarshal method.
		PublishMarshal []struct {
			// Topic is the topic argument value.
			Topic string
			// M is the m argument value.
			M protoreflect.ProtoMessage
		}
		// Subscribe holds details about calls to the Subscribe method.
		Subscribe []struct {
			// Topic is the topic argument value.
			Topic string
			// MsgFunc is the msgFunc argument value.
			MsgFunc func([]byte) error
		}
	}
	lockPublishMarshal sync.RWMutex
	lockSubscribe      sync.RWMutex
}

// PublishMarshal calls PublishMarshalFunc.
func (mock *MessageQueueClientMock) PublishMarshal(topic string, m protoreflect.ProtoMessage) error {
	if mock.PublishMarshalFunc == nil {
		panic("MessageQueueClientMock.PublishMarshalFunc: method is nil but MessageQueueClient.PublishMarshal was just called")
	}
	callInfo := struct {
		Topic string
		M     protoreflect.ProtoMessage
	}{
		Topic: topic,
		M:     m,
	}
	mock.lockPublishMarshal.Lock()
	mock.calls.PublishMarshal = append(mock.calls.PublishMarshal, callInfo)
	mock.lockPublishMarshal.Unlock()
	return mock.PublishMarshalFunc(topic, m)
}

// PublishMarshalCalls gets all the calls that were made to PublishMarshal.
// Check the length with:
//
//	len(mockedMessageQueueClient.PublishMarshalCalls())
func (mock *MessageQueueClientMock) PublishMarshalCalls() []struct {
	Topic string
	M     protoreflect.ProtoMessage
} {
	var calls []struct {
		Topic string
		M     protoreflect.ProtoMessage
	}
	mock.lockPublishMarshal.RLock()
	calls = mock.calls.PublishMarshal
	mock.lockPublishMarshal.RUnlock()
	return calls
}

// Subscribe calls SubscribeFunc.
func (mock *MessageQueueClientMock) Subscribe(topic string, msgFunc func([]byte) error) error {
	if mock.SubscribeFunc == nil {
		panic("MessageQueueClientMock.SubscribeFunc: method is nil but MessageQueueClient.Subscribe was just called")
	}
	callInfo := struct {
		Topic   string
		MsgFunc func([]byte) error
	}{
		Topic:   topic,
		MsgFunc: msgFunc,
	}
	mock.lockSubscribe.Lock()
	mock.calls.Subscribe = append(mock.calls.Subscribe, callInfo)
	mock.lockSubscribe.Unlock()
	return mock.SubscribeFunc(topic, msgFunc)
}

// SubscribeCalls gets all the calls that were made to Subscribe.
// Check the length with:
//
//	len(mockedMessageQueueClient.SubscribeCalls())
func (mock *MessageQueueClientMock) SubscribeCalls() []struct {
	Topic   string
	MsgFunc func([]byte) error
} {
	var calls []struct {
		Topic   string
		MsgFunc func([]byte) error
	}
	mock.lockSubscribe.RLock()
	calls = mock.calls.Subscribe
	mock.lockSubscribe.RUnlock()
	return calls
}
