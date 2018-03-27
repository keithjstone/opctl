// Code generated by counterfeiter. DO NOT EDIT.
package ref

import (
	"net/http"
	"sync"
)

type FakeHandler struct {
	HandleStub        func(res http.ResponseWriter, req *http.Request)
	handleMutex       sync.RWMutex
	handleArgsForCall []struct {
		res http.ResponseWriter
		req *http.Request
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeHandler) Handle(res http.ResponseWriter, req *http.Request) {
	fake.handleMutex.Lock()
	fake.handleArgsForCall = append(fake.handleArgsForCall, struct {
		res http.ResponseWriter
		req *http.Request
	}{res, req})
	fake.recordInvocation("Handle", []interface{}{res, req})
	fake.handleMutex.Unlock()
	if fake.HandleStub != nil {
		fake.HandleStub(res, req)
	}
}

func (fake *FakeHandler) HandleCallCount() int {
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	return len(fake.handleArgsForCall)
}

func (fake *FakeHandler) HandleArgsForCall(i int) (http.ResponseWriter, *http.Request) {
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	return fake.handleArgsForCall[i].res, fake.handleArgsForCall[i].req
}

func (fake *FakeHandler) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeHandler) recordInvocation(key string, args []interface{}) {
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

var _ Handler = new(FakeHandler)
