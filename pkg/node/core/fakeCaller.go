// This file was generated by counterfeiter
package core

import (
	"sync"

	"github.com/opspec-io/sdk-golang/pkg/model"
)

type fakeCaller struct {
	CallStub        func(nodeId string, args map[string]*model.Data, scg *model.Scg, opPkgRef string, rootOpId string) (outputs map[string]*model.Data, err error)
	callMutex       sync.RWMutex
	callArgsForCall []struct {
		nodeId   string
		args     map[string]*model.Data
		scg      *model.Scg
		opPkgRef string
		rootOpId string
	}
	callReturns struct {
		result1 map[string]*model.Data
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *fakeCaller) Call(nodeId string, args map[string]*model.Data, scg *model.Scg, opPkgRef string, rootOpId string) (outputs map[string]*model.Data, err error) {
	fake.callMutex.Lock()
	fake.callArgsForCall = append(fake.callArgsForCall, struct {
		nodeId   string
		args     map[string]*model.Data
		scg      *model.Scg
		opPkgRef string
		rootOpId string
	}{nodeId, args, scg, opPkgRef, rootOpId})
	fake.recordInvocation("Call", []interface{}{nodeId, args, scg, opPkgRef, rootOpId})
	fake.callMutex.Unlock()
	if fake.CallStub != nil {
		return fake.CallStub(nodeId, args, scg, opPkgRef, rootOpId)
	} else {
		return fake.callReturns.result1, fake.callReturns.result2
	}
}

func (fake *fakeCaller) CallCallCount() int {
	fake.callMutex.RLock()
	defer fake.callMutex.RUnlock()
	return len(fake.callArgsForCall)
}

func (fake *fakeCaller) CallArgsForCall(i int) (string, map[string]*model.Data, *model.Scg, string, string) {
	fake.callMutex.RLock()
	defer fake.callMutex.RUnlock()
	return fake.callArgsForCall[i].nodeId, fake.callArgsForCall[i].args, fake.callArgsForCall[i].scg, fake.callArgsForCall[i].opPkgRef, fake.callArgsForCall[i].rootOpId
}

func (fake *fakeCaller) CallReturns(result1 map[string]*model.Data, result2 error) {
	fake.CallStub = nil
	fake.callReturns = struct {
		result1 map[string]*model.Data
		result2 error
	}{result1, result2}
}

func (fake *fakeCaller) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.callMutex.RLock()
	defer fake.callMutex.RUnlock()
	return fake.invocations
}

func (fake *fakeCaller) recordInvocation(key string, args []interface{}) {
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
