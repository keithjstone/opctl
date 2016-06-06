// This file was generated by counterfeiter
package sdk

import (
  "sync"

  "github.com/opspec-io/sdk-golang/models"
)

type FakeApi struct {
  AddOpStub                     func(req models.AddOpReq) (err error)
  addOpMutex                    sync.RWMutex
  addOpArgsForCall              []struct {
    req models.AddOpReq
  }
  addOpReturns                  struct {
                                  result1 error
                                }
  SetDescriptionOfOpStub        func(req models.SetDescriptionOfOpReq) (err error)
  setDescriptionOfOpMutex       sync.RWMutex
  setDescriptionOfOpArgsForCall []struct {
    req models.SetDescriptionOfOpReq
  }
  setDescriptionOfOpReturns     struct {
                                  result1 error
                                }
  invocations                   map[string][][]interface{}
  invocationsMutex              sync.RWMutex
}

func (fake *FakeApi) AddOp(req models.AddOpReq) (err error) {
  fake.addOpMutex.Lock()
  fake.addOpArgsForCall = append(fake.addOpArgsForCall, struct {
    req models.AddOpReq
  }{req})
  fake.recordInvocation("AddOp", []interface{}{req})
  fake.addOpMutex.Unlock()
  if fake.AddOpStub != nil {
    return fake.AddOpStub(req)
  } else {
    return fake.addOpReturns.result1
  }
}

func (fake *FakeApi) AddOpCallCount() int {
  fake.addOpMutex.RLock()
  defer fake.addOpMutex.RUnlock()
  return len(fake.addOpArgsForCall)
}

func (fake *FakeApi) AddOpArgsForCall(i int) models.AddOpReq {
  fake.addOpMutex.RLock()
  defer fake.addOpMutex.RUnlock()
  return fake.addOpArgsForCall[i].req
}

func (fake *FakeApi) AddOpReturns(result1 error) {
  fake.AddOpStub = nil
  fake.addOpReturns = struct {
    result1 error
  }{result1}
}

func (fake *FakeApi) SetDescriptionOfOp(req models.SetDescriptionOfOpReq) (err error) {
  fake.setDescriptionOfOpMutex.Lock()
  fake.setDescriptionOfOpArgsForCall = append(fake.setDescriptionOfOpArgsForCall, struct {
    req models.SetDescriptionOfOpReq
  }{req})
  fake.recordInvocation("SetDescriptionOfOp", []interface{}{req})
  fake.setDescriptionOfOpMutex.Unlock()
  if fake.SetDescriptionOfOpStub != nil {
    return fake.SetDescriptionOfOpStub(req)
  } else {
    return fake.setDescriptionOfOpReturns.result1
  }
}

func (fake *FakeApi) SetDescriptionOfOpCallCount() int {
  fake.setDescriptionOfOpMutex.RLock()
  defer fake.setDescriptionOfOpMutex.RUnlock()
  return len(fake.setDescriptionOfOpArgsForCall)
}

func (fake *FakeApi) SetDescriptionOfOpArgsForCall(i int) models.SetDescriptionOfOpReq {
  fake.setDescriptionOfOpMutex.RLock()
  defer fake.setDescriptionOfOpMutex.RUnlock()
  return fake.setDescriptionOfOpArgsForCall[i].req
}

func (fake *FakeApi) SetDescriptionOfOpReturns(result1 error) {
  fake.SetDescriptionOfOpStub = nil
  fake.setDescriptionOfOpReturns = struct {
    result1 error
  }{result1}
}

func (fake *FakeApi) Invocations() map[string][][]interface{} {
  fake.invocationsMutex.RLock()
  defer fake.invocationsMutex.RUnlock()
  fake.addOpMutex.RLock()
  defer fake.addOpMutex.RUnlock()
  fake.setDescriptionOfOpMutex.RLock()
  defer fake.setDescriptionOfOpMutex.RUnlock()
  return fake.invocations
}

func (fake *FakeApi) recordInvocation(key string, args []interface{}) {
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

var _ Api = new(FakeApi)
