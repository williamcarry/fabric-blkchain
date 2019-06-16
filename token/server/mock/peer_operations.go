// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/hyperledger/fabric/common/channelconfig"
	"github.com/hyperledger/fabric/common/policies"
	"github.com/hyperledger/fabric/core/ledger"
	"github.com/hyperledger/fabric/protos/common"
)

type PeerOperations struct {
	GetApplicationConfigStub        func(string) (channelconfig.Application, bool)
	getApplicationConfigMutex       sync.RWMutex
	getApplicationConfigArgsForCall []struct {
		arg1 string
	}
	getApplicationConfigReturns struct {
		result1 channelconfig.Application
		result2 bool
	}
	getApplicationConfigReturnsOnCall map[int]struct {
		result1 channelconfig.Application
		result2 bool
	}
	GetChannelConfigStub        func(string) channelconfig.Resources
	getChannelConfigMutex       sync.RWMutex
	getChannelConfigArgsForCall []struct {
		arg1 string
	}
	getChannelConfigReturns struct {
		result1 channelconfig.Resources
	}
	getChannelConfigReturnsOnCall map[int]struct {
		result1 channelconfig.Resources
	}
	GetCurrConfigBlockStub        func(string) *common.Block
	getCurrConfigBlockMutex       sync.RWMutex
	getCurrConfigBlockArgsForCall []struct {
		arg1 string
	}
	getCurrConfigBlockReturns struct {
		result1 *common.Block
	}
	getCurrConfigBlockReturnsOnCall map[int]struct {
		result1 *common.Block
	}
	GetLedgerStub        func(string) ledger.PeerLedger
	getLedgerMutex       sync.RWMutex
	getLedgerArgsForCall []struct {
		arg1 string
	}
	getLedgerReturns struct {
		result1 ledger.PeerLedger
	}
	getLedgerReturnsOnCall map[int]struct {
		result1 ledger.PeerLedger
	}
	GetMSPIDsStub        func(string) []string
	getMSPIDsMutex       sync.RWMutex
	getMSPIDsArgsForCall []struct {
		arg1 string
	}
	getMSPIDsReturns struct {
		result1 []string
	}
	getMSPIDsReturnsOnCall map[int]struct {
		result1 []string
	}
	GetPolicyManagerStub        func(string) policies.Manager
	getPolicyManagerMutex       sync.RWMutex
	getPolicyManagerArgsForCall []struct {
		arg1 string
	}
	getPolicyManagerReturns struct {
		result1 policies.Manager
	}
	getPolicyManagerReturnsOnCall map[int]struct {
		result1 policies.Manager
	}
	GetStableChannelConfigStub        func(string) channelconfig.Resources
	getStableChannelConfigMutex       sync.RWMutex
	getStableChannelConfigArgsForCall []struct {
		arg1 string
	}
	getStableChannelConfigReturns struct {
		result1 channelconfig.Resources
	}
	getStableChannelConfigReturnsOnCall map[int]struct {
		result1 channelconfig.Resources
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *PeerOperations) GetApplicationConfig(arg1 string) (channelconfig.Application, bool) {
	fake.getApplicationConfigMutex.Lock()
	ret, specificReturn := fake.getApplicationConfigReturnsOnCall[len(fake.getApplicationConfigArgsForCall)]
	fake.getApplicationConfigArgsForCall = append(fake.getApplicationConfigArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetApplicationConfig", []interface{}{arg1})
	fake.getApplicationConfigMutex.Unlock()
	if fake.GetApplicationConfigStub != nil {
		return fake.GetApplicationConfigStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getApplicationConfigReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *PeerOperations) GetApplicationConfigCallCount() int {
	fake.getApplicationConfigMutex.RLock()
	defer fake.getApplicationConfigMutex.RUnlock()
	return len(fake.getApplicationConfigArgsForCall)
}

func (fake *PeerOperations) GetApplicationConfigCalls(stub func(string) (channelconfig.Application, bool)) {
	fake.getApplicationConfigMutex.Lock()
	defer fake.getApplicationConfigMutex.Unlock()
	fake.GetApplicationConfigStub = stub
}

func (fake *PeerOperations) GetApplicationConfigArgsForCall(i int) string {
	fake.getApplicationConfigMutex.RLock()
	defer fake.getApplicationConfigMutex.RUnlock()
	argsForCall := fake.getApplicationConfigArgsForCall[i]
	return argsForCall.arg1
}

func (fake *PeerOperations) GetApplicationConfigReturns(result1 channelconfig.Application, result2 bool) {
	fake.getApplicationConfigMutex.Lock()
	defer fake.getApplicationConfigMutex.Unlock()
	fake.GetApplicationConfigStub = nil
	fake.getApplicationConfigReturns = struct {
		result1 channelconfig.Application
		result2 bool
	}{result1, result2}
}

func (fake *PeerOperations) GetApplicationConfigReturnsOnCall(i int, result1 channelconfig.Application, result2 bool) {
	fake.getApplicationConfigMutex.Lock()
	defer fake.getApplicationConfigMutex.Unlock()
	fake.GetApplicationConfigStub = nil
	if fake.getApplicationConfigReturnsOnCall == nil {
		fake.getApplicationConfigReturnsOnCall = make(map[int]struct {
			result1 channelconfig.Application
			result2 bool
		})
	}
	fake.getApplicationConfigReturnsOnCall[i] = struct {
		result1 channelconfig.Application
		result2 bool
	}{result1, result2}
}

func (fake *PeerOperations) GetChannelConfig(arg1 string) channelconfig.Resources {
	fake.getChannelConfigMutex.Lock()
	ret, specificReturn := fake.getChannelConfigReturnsOnCall[len(fake.getChannelConfigArgsForCall)]
	fake.getChannelConfigArgsForCall = append(fake.getChannelConfigArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetChannelConfig", []interface{}{arg1})
	fake.getChannelConfigMutex.Unlock()
	if fake.GetChannelConfigStub != nil {
		return fake.GetChannelConfigStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getChannelConfigReturns
	return fakeReturns.result1
}

func (fake *PeerOperations) GetChannelConfigCallCount() int {
	fake.getChannelConfigMutex.RLock()
	defer fake.getChannelConfigMutex.RUnlock()
	return len(fake.getChannelConfigArgsForCall)
}

func (fake *PeerOperations) GetChannelConfigCalls(stub func(string) channelconfig.Resources) {
	fake.getChannelConfigMutex.Lock()
	defer fake.getChannelConfigMutex.Unlock()
	fake.GetChannelConfigStub = stub
}

func (fake *PeerOperations) GetChannelConfigArgsForCall(i int) string {
	fake.getChannelConfigMutex.RLock()
	defer fake.getChannelConfigMutex.RUnlock()
	argsForCall := fake.getChannelConfigArgsForCall[i]
	return argsForCall.arg1
}

func (fake *PeerOperations) GetChannelConfigReturns(result1 channelconfig.Resources) {
	fake.getChannelConfigMutex.Lock()
	defer fake.getChannelConfigMutex.Unlock()
	fake.GetChannelConfigStub = nil
	fake.getChannelConfigReturns = struct {
		result1 channelconfig.Resources
	}{result1}
}

func (fake *PeerOperations) GetChannelConfigReturnsOnCall(i int, result1 channelconfig.Resources) {
	fake.getChannelConfigMutex.Lock()
	defer fake.getChannelConfigMutex.Unlock()
	fake.GetChannelConfigStub = nil
	if fake.getChannelConfigReturnsOnCall == nil {
		fake.getChannelConfigReturnsOnCall = make(map[int]struct {
			result1 channelconfig.Resources
		})
	}
	fake.getChannelConfigReturnsOnCall[i] = struct {
		result1 channelconfig.Resources
	}{result1}
}

func (fake *PeerOperations) GetCurrConfigBlock(arg1 string) *common.Block {
	fake.getCurrConfigBlockMutex.Lock()
	ret, specificReturn := fake.getCurrConfigBlockReturnsOnCall[len(fake.getCurrConfigBlockArgsForCall)]
	fake.getCurrConfigBlockArgsForCall = append(fake.getCurrConfigBlockArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetCurrConfigBlock", []interface{}{arg1})
	fake.getCurrConfigBlockMutex.Unlock()
	if fake.GetCurrConfigBlockStub != nil {
		return fake.GetCurrConfigBlockStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getCurrConfigBlockReturns
	return fakeReturns.result1
}

func (fake *PeerOperations) GetCurrConfigBlockCallCount() int {
	fake.getCurrConfigBlockMutex.RLock()
	defer fake.getCurrConfigBlockMutex.RUnlock()
	return len(fake.getCurrConfigBlockArgsForCall)
}

func (fake *PeerOperations) GetCurrConfigBlockCalls(stub func(string) *common.Block) {
	fake.getCurrConfigBlockMutex.Lock()
	defer fake.getCurrConfigBlockMutex.Unlock()
	fake.GetCurrConfigBlockStub = stub
}

func (fake *PeerOperations) GetCurrConfigBlockArgsForCall(i int) string {
	fake.getCurrConfigBlockMutex.RLock()
	defer fake.getCurrConfigBlockMutex.RUnlock()
	argsForCall := fake.getCurrConfigBlockArgsForCall[i]
	return argsForCall.arg1
}

func (fake *PeerOperations) GetCurrConfigBlockReturns(result1 *common.Block) {
	fake.getCurrConfigBlockMutex.Lock()
	defer fake.getCurrConfigBlockMutex.Unlock()
	fake.GetCurrConfigBlockStub = nil
	fake.getCurrConfigBlockReturns = struct {
		result1 *common.Block
	}{result1}
}

func (fake *PeerOperations) GetCurrConfigBlockReturnsOnCall(i int, result1 *common.Block) {
	fake.getCurrConfigBlockMutex.Lock()
	defer fake.getCurrConfigBlockMutex.Unlock()
	fake.GetCurrConfigBlockStub = nil
	if fake.getCurrConfigBlockReturnsOnCall == nil {
		fake.getCurrConfigBlockReturnsOnCall = make(map[int]struct {
			result1 *common.Block
		})
	}
	fake.getCurrConfigBlockReturnsOnCall[i] = struct {
		result1 *common.Block
	}{result1}
}

func (fake *PeerOperations) GetLedger(arg1 string) ledger.PeerLedger {
	fake.getLedgerMutex.Lock()
	ret, specificReturn := fake.getLedgerReturnsOnCall[len(fake.getLedgerArgsForCall)]
	fake.getLedgerArgsForCall = append(fake.getLedgerArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetLedger", []interface{}{arg1})
	fake.getLedgerMutex.Unlock()
	if fake.GetLedgerStub != nil {
		return fake.GetLedgerStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getLedgerReturns
	return fakeReturns.result1
}

func (fake *PeerOperations) GetLedgerCallCount() int {
	fake.getLedgerMutex.RLock()
	defer fake.getLedgerMutex.RUnlock()
	return len(fake.getLedgerArgsForCall)
}

func (fake *PeerOperations) GetLedgerCalls(stub func(string) ledger.PeerLedger) {
	fake.getLedgerMutex.Lock()
	defer fake.getLedgerMutex.Unlock()
	fake.GetLedgerStub = stub
}

func (fake *PeerOperations) GetLedgerArgsForCall(i int) string {
	fake.getLedgerMutex.RLock()
	defer fake.getLedgerMutex.RUnlock()
	argsForCall := fake.getLedgerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *PeerOperations) GetLedgerReturns(result1 ledger.PeerLedger) {
	fake.getLedgerMutex.Lock()
	defer fake.getLedgerMutex.Unlock()
	fake.GetLedgerStub = nil
	fake.getLedgerReturns = struct {
		result1 ledger.PeerLedger
	}{result1}
}

func (fake *PeerOperations) GetLedgerReturnsOnCall(i int, result1 ledger.PeerLedger) {
	fake.getLedgerMutex.Lock()
	defer fake.getLedgerMutex.Unlock()
	fake.GetLedgerStub = nil
	if fake.getLedgerReturnsOnCall == nil {
		fake.getLedgerReturnsOnCall = make(map[int]struct {
			result1 ledger.PeerLedger
		})
	}
	fake.getLedgerReturnsOnCall[i] = struct {
		result1 ledger.PeerLedger
	}{result1}
}

func (fake *PeerOperations) GetMSPIDs(arg1 string) []string {
	fake.getMSPIDsMutex.Lock()
	ret, specificReturn := fake.getMSPIDsReturnsOnCall[len(fake.getMSPIDsArgsForCall)]
	fake.getMSPIDsArgsForCall = append(fake.getMSPIDsArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetMSPIDs", []interface{}{arg1})
	fake.getMSPIDsMutex.Unlock()
	if fake.GetMSPIDsStub != nil {
		return fake.GetMSPIDsStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getMSPIDsReturns
	return fakeReturns.result1
}

func (fake *PeerOperations) GetMSPIDsCallCount() int {
	fake.getMSPIDsMutex.RLock()
	defer fake.getMSPIDsMutex.RUnlock()
	return len(fake.getMSPIDsArgsForCall)
}

func (fake *PeerOperations) GetMSPIDsCalls(stub func(string) []string) {
	fake.getMSPIDsMutex.Lock()
	defer fake.getMSPIDsMutex.Unlock()
	fake.GetMSPIDsStub = stub
}

func (fake *PeerOperations) GetMSPIDsArgsForCall(i int) string {
	fake.getMSPIDsMutex.RLock()
	defer fake.getMSPIDsMutex.RUnlock()
	argsForCall := fake.getMSPIDsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *PeerOperations) GetMSPIDsReturns(result1 []string) {
	fake.getMSPIDsMutex.Lock()
	defer fake.getMSPIDsMutex.Unlock()
	fake.GetMSPIDsStub = nil
	fake.getMSPIDsReturns = struct {
		result1 []string
	}{result1}
}

func (fake *PeerOperations) GetMSPIDsReturnsOnCall(i int, result1 []string) {
	fake.getMSPIDsMutex.Lock()
	defer fake.getMSPIDsMutex.Unlock()
	fake.GetMSPIDsStub = nil
	if fake.getMSPIDsReturnsOnCall == nil {
		fake.getMSPIDsReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.getMSPIDsReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *PeerOperations) GetPolicyManager(arg1 string) policies.Manager {
	fake.getPolicyManagerMutex.Lock()
	ret, specificReturn := fake.getPolicyManagerReturnsOnCall[len(fake.getPolicyManagerArgsForCall)]
	fake.getPolicyManagerArgsForCall = append(fake.getPolicyManagerArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetPolicyManager", []interface{}{arg1})
	fake.getPolicyManagerMutex.Unlock()
	if fake.GetPolicyManagerStub != nil {
		return fake.GetPolicyManagerStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getPolicyManagerReturns
	return fakeReturns.result1
}

func (fake *PeerOperations) GetPolicyManagerCallCount() int {
	fake.getPolicyManagerMutex.RLock()
	defer fake.getPolicyManagerMutex.RUnlock()
	return len(fake.getPolicyManagerArgsForCall)
}

func (fake *PeerOperations) GetPolicyManagerCalls(stub func(string) policies.Manager) {
	fake.getPolicyManagerMutex.Lock()
	defer fake.getPolicyManagerMutex.Unlock()
	fake.GetPolicyManagerStub = stub
}

func (fake *PeerOperations) GetPolicyManagerArgsForCall(i int) string {
	fake.getPolicyManagerMutex.RLock()
	defer fake.getPolicyManagerMutex.RUnlock()
	argsForCall := fake.getPolicyManagerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *PeerOperations) GetPolicyManagerReturns(result1 policies.Manager) {
	fake.getPolicyManagerMutex.Lock()
	defer fake.getPolicyManagerMutex.Unlock()
	fake.GetPolicyManagerStub = nil
	fake.getPolicyManagerReturns = struct {
		result1 policies.Manager
	}{result1}
}

func (fake *PeerOperations) GetPolicyManagerReturnsOnCall(i int, result1 policies.Manager) {
	fake.getPolicyManagerMutex.Lock()
	defer fake.getPolicyManagerMutex.Unlock()
	fake.GetPolicyManagerStub = nil
	if fake.getPolicyManagerReturnsOnCall == nil {
		fake.getPolicyManagerReturnsOnCall = make(map[int]struct {
			result1 policies.Manager
		})
	}
	fake.getPolicyManagerReturnsOnCall[i] = struct {
		result1 policies.Manager
	}{result1}
}

func (fake *PeerOperations) GetStableChannelConfig(arg1 string) channelconfig.Resources {
	fake.getStableChannelConfigMutex.Lock()
	ret, specificReturn := fake.getStableChannelConfigReturnsOnCall[len(fake.getStableChannelConfigArgsForCall)]
	fake.getStableChannelConfigArgsForCall = append(fake.getStableChannelConfigArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetStableChannelConfig", []interface{}{arg1})
	fake.getStableChannelConfigMutex.Unlock()
	if fake.GetStableChannelConfigStub != nil {
		return fake.GetStableChannelConfigStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getStableChannelConfigReturns
	return fakeReturns.result1
}

func (fake *PeerOperations) GetStableChannelConfigCallCount() int {
	fake.getStableChannelConfigMutex.RLock()
	defer fake.getStableChannelConfigMutex.RUnlock()
	return len(fake.getStableChannelConfigArgsForCall)
}

func (fake *PeerOperations) GetStableChannelConfigCalls(stub func(string) channelconfig.Resources) {
	fake.getStableChannelConfigMutex.Lock()
	defer fake.getStableChannelConfigMutex.Unlock()
	fake.GetStableChannelConfigStub = stub
}

func (fake *PeerOperations) GetStableChannelConfigArgsForCall(i int) string {
	fake.getStableChannelConfigMutex.RLock()
	defer fake.getStableChannelConfigMutex.RUnlock()
	argsForCall := fake.getStableChannelConfigArgsForCall[i]
	return argsForCall.arg1
}

func (fake *PeerOperations) GetStableChannelConfigReturns(result1 channelconfig.Resources) {
	fake.getStableChannelConfigMutex.Lock()
	defer fake.getStableChannelConfigMutex.Unlock()
	fake.GetStableChannelConfigStub = nil
	fake.getStableChannelConfigReturns = struct {
		result1 channelconfig.Resources
	}{result1}
}

func (fake *PeerOperations) GetStableChannelConfigReturnsOnCall(i int, result1 channelconfig.Resources) {
	fake.getStableChannelConfigMutex.Lock()
	defer fake.getStableChannelConfigMutex.Unlock()
	fake.GetStableChannelConfigStub = nil
	if fake.getStableChannelConfigReturnsOnCall == nil {
		fake.getStableChannelConfigReturnsOnCall = make(map[int]struct {
			result1 channelconfig.Resources
		})
	}
	fake.getStableChannelConfigReturnsOnCall[i] = struct {
		result1 channelconfig.Resources
	}{result1}
}

func (fake *PeerOperations) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getApplicationConfigMutex.RLock()
	defer fake.getApplicationConfigMutex.RUnlock()
	fake.getChannelConfigMutex.RLock()
	defer fake.getChannelConfigMutex.RUnlock()
	fake.getCurrConfigBlockMutex.RLock()
	defer fake.getCurrConfigBlockMutex.RUnlock()
	fake.getLedgerMutex.RLock()
	defer fake.getLedgerMutex.RUnlock()
	fake.getMSPIDsMutex.RLock()
	defer fake.getMSPIDsMutex.RUnlock()
	fake.getPolicyManagerMutex.RLock()
	defer fake.getPolicyManagerMutex.RUnlock()
	fake.getStableChannelConfigMutex.RLock()
	defer fake.getStableChannelConfigMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *PeerOperations) recordInvocation(key string, args []interface{}) {
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