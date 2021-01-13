// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"context"
	"sync"

	"github.com/opctl/opctl/sdks/go/model"
	"github.com/opctl/opctl/sdks/go/pubsub"
)

type FakeEventSubscriber struct {
	SubscribeStub        func(context.Context, model.EventFilter) (<-chan model.Event, error)
	subscribeMutex       sync.RWMutex
	subscribeArgsForCall []struct {
		arg1 context.Context
		arg2 model.EventFilter
	}
	subscribeReturns struct {
		result1 <-chan model.Event
		result2 error
	}
	subscribeReturnsOnCall map[int]struct {
		result1 <-chan model.Event
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEventSubscriber) Subscribe(arg1 context.Context, arg2 model.EventFilter) (<-chan model.Event, error) {
	fake.subscribeMutex.Lock()
	ret, specificReturn := fake.subscribeReturnsOnCall[len(fake.subscribeArgsForCall)]
	fake.subscribeArgsForCall = append(fake.subscribeArgsForCall, struct {
		arg1 context.Context
		arg2 model.EventFilter
	}{arg1, arg2})
	fake.recordInvocation("Subscribe", []interface{}{arg1, arg2})
	fake.subscribeMutex.Unlock()
	if fake.SubscribeStub != nil {
		return fake.SubscribeStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.subscribeReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeEventSubscriber) SubscribeCallCount() int {
	fake.subscribeMutex.RLock()
	defer fake.subscribeMutex.RUnlock()
	return len(fake.subscribeArgsForCall)
}

func (fake *FakeEventSubscriber) SubscribeCalls(stub func(context.Context, model.EventFilter) (<-chan model.Event, error)) {
	fake.subscribeMutex.Lock()
	defer fake.subscribeMutex.Unlock()
	fake.SubscribeStub = stub
}

func (fake *FakeEventSubscriber) SubscribeArgsForCall(i int) (context.Context, model.EventFilter) {
	fake.subscribeMutex.RLock()
	defer fake.subscribeMutex.RUnlock()
	argsForCall := fake.subscribeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeEventSubscriber) SubscribeReturns(result1 <-chan model.Event, result2 error) {
	fake.subscribeMutex.Lock()
	defer fake.subscribeMutex.Unlock()
	fake.SubscribeStub = nil
	fake.subscribeReturns = struct {
		result1 <-chan model.Event
		result2 error
	}{result1, result2}
}

func (fake *FakeEventSubscriber) SubscribeReturnsOnCall(i int, result1 <-chan model.Event, result2 error) {
	fake.subscribeMutex.Lock()
	defer fake.subscribeMutex.Unlock()
	fake.SubscribeStub = nil
	if fake.subscribeReturnsOnCall == nil {
		fake.subscribeReturnsOnCall = make(map[int]struct {
			result1 <-chan model.Event
			result2 error
		})
	}
	fake.subscribeReturnsOnCall[i] = struct {
		result1 <-chan model.Event
		result2 error
	}{result1, result2}
}

func (fake *FakeEventSubscriber) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.subscribeMutex.RLock()
	defer fake.subscribeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEventSubscriber) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ pubsub.EventSubscriber = new(FakeEventSubscriber)
