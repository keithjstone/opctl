// This file was generated by counterfeiter
package fake

import (
  "sync"

  "github.com/opspec-io/sdk-golang/pkg/engineprovider"
)

type EngineProvider struct {
  EnsureEngineRunningStub                     func() (err error)
  ensureEngineRunningMutex                    sync.RWMutex
  ensureEngineRunningArgsForCall              []struct{}
  ensureEngineRunningReturns                  struct {
                                                result1 error
                                              }
  GetEngineProtocolRelativeBaseUrlStub        func() (protocolRelativeBaseUrl string, err error)
  getEngineProtocolRelativeBaseUrlMutex       sync.RWMutex
  getEngineProtocolRelativeBaseUrlArgsForCall []struct{}
  getEngineProtocolRelativeBaseUrlReturns     struct {
                                                result1 string
                                                result2 error
                                              }
}

func (fake *EngineProvider) EnsureEngineRunning() (err error) {
  fake.ensureEngineRunningMutex.Lock()
  fake.ensureEngineRunningArgsForCall = append(fake.ensureEngineRunningArgsForCall, struct{}{})
  fake.ensureEngineRunningMutex.Unlock()
  if fake.EnsureEngineRunningStub != nil {
    return fake.EnsureEngineRunningStub()
  } else {
    return fake.ensureEngineRunningReturns.result1
  }
}

func (fake *EngineProvider) EnsureEngineRunningCallCount() int {
  fake.ensureEngineRunningMutex.RLock()
  defer fake.ensureEngineRunningMutex.RUnlock()
  return len(fake.ensureEngineRunningArgsForCall)
}

func (fake *EngineProvider) EnsureEngineRunningReturns(result1 error) {
  fake.EnsureEngineRunningStub = nil
  fake.ensureEngineRunningReturns = struct {
    result1 error
  }{result1}
}

func (fake *EngineProvider) GetEngineProtocolRelativeBaseUrl() (protocolRelativeBaseUrl string, err error) {
  fake.getEngineProtocolRelativeBaseUrlMutex.Lock()
  fake.getEngineProtocolRelativeBaseUrlArgsForCall = append(fake.getEngineProtocolRelativeBaseUrlArgsForCall, struct{}{})
  fake.getEngineProtocolRelativeBaseUrlMutex.Unlock()
  if fake.GetEngineProtocolRelativeBaseUrlStub != nil {
    return fake.GetEngineProtocolRelativeBaseUrlStub()
  } else {
    return fake.getEngineProtocolRelativeBaseUrlReturns.result1, fake.getEngineProtocolRelativeBaseUrlReturns.result2
  }
}

func (fake *EngineProvider) GetEngineProtocolRelativeBaseUrlCallCount() int {
  fake.getEngineProtocolRelativeBaseUrlMutex.RLock()
  defer fake.getEngineProtocolRelativeBaseUrlMutex.RUnlock()
  return len(fake.getEngineProtocolRelativeBaseUrlArgsForCall)
}

func (fake *EngineProvider) GetEngineProtocolRelativeBaseUrlReturns(result1 string, result2 error) {
  fake.GetEngineProtocolRelativeBaseUrlStub = nil
  fake.getEngineProtocolRelativeBaseUrlReturns = struct {
    result1 string
    result2 error
  }{result1, result2}
}

var _ engineprovider.EngineProvider = new(EngineProvider)
