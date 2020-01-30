// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	sync "sync"

	cfclient "github.com/cloudfoundry-community/go-cfclient"
	organization "github.com/pivotalservices/cf-mgmt/organization"
)

type FakeManager struct {
	CreateOrgsStub        func() error
	createOrgsMutex       sync.RWMutex
	createOrgsArgsForCall []struct {
	}
	createOrgsReturns struct {
		result1 error
	}
	createOrgsReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteOrgsStub        func() error
	deleteOrgsMutex       sync.RWMutex
	deleteOrgsArgsForCall []struct {
	}
	deleteOrgsReturns struct {
		result1 error
	}
	deleteOrgsReturnsOnCall map[int]struct {
		result1 error
	}
	RenameOrgStub        func(string, string) error
	renameOrgMutex       sync.RWMutex
	renameOrgArgsForCall []struct {
		arg1 string
		arg2 string
	}
	renameOrgReturns struct {
		result1 error
	}
	renameOrgReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateOrgStub        func(string, cfclient.OrgRequest) (cfclient.Org, error)
	updateOrgMutex       sync.RWMutex
	updateOrgArgsForCall []struct {
		arg1 string
		arg2 cfclient.OrgRequest
	}
	updateOrgReturns struct {
		result1 cfclient.Org
		result2 error
	}
	updateOrgReturnsOnCall map[int]struct {
		result1 cfclient.Org
		result2 error
	}
	UpdateOrgsMetadataStub        func() error
	updateOrgsMetadataMutex       sync.RWMutex
	updateOrgsMetadataArgsForCall []struct {
	}
	updateOrgsMetadataReturns struct {
		result1 error
	}
	updateOrgsMetadataReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeManager) CreateOrgs() error {
	fake.createOrgsMutex.Lock()
	ret, specificReturn := fake.createOrgsReturnsOnCall[len(fake.createOrgsArgsForCall)]
	fake.createOrgsArgsForCall = append(fake.createOrgsArgsForCall, struct {
	}{})
	fake.recordInvocation("CreateOrgs", []interface{}{})
	fake.createOrgsMutex.Unlock()
	if fake.CreateOrgsStub != nil {
		return fake.CreateOrgsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.createOrgsReturns
	return fakeReturns.result1
}

func (fake *FakeManager) CreateOrgsCallCount() int {
	fake.createOrgsMutex.RLock()
	defer fake.createOrgsMutex.RUnlock()
	return len(fake.createOrgsArgsForCall)
}

func (fake *FakeManager) CreateOrgsCalls(stub func() error) {
	fake.createOrgsMutex.Lock()
	defer fake.createOrgsMutex.Unlock()
	fake.CreateOrgsStub = stub
}

func (fake *FakeManager) CreateOrgsReturns(result1 error) {
	fake.createOrgsMutex.Lock()
	defer fake.createOrgsMutex.Unlock()
	fake.CreateOrgsStub = nil
	fake.createOrgsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) CreateOrgsReturnsOnCall(i int, result1 error) {
	fake.createOrgsMutex.Lock()
	defer fake.createOrgsMutex.Unlock()
	fake.CreateOrgsStub = nil
	if fake.createOrgsReturnsOnCall == nil {
		fake.createOrgsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createOrgsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) DeleteOrgs() error {
	fake.deleteOrgsMutex.Lock()
	ret, specificReturn := fake.deleteOrgsReturnsOnCall[len(fake.deleteOrgsArgsForCall)]
	fake.deleteOrgsArgsForCall = append(fake.deleteOrgsArgsForCall, struct {
	}{})
	fake.recordInvocation("DeleteOrgs", []interface{}{})
	fake.deleteOrgsMutex.Unlock()
	if fake.DeleteOrgsStub != nil {
		return fake.DeleteOrgsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteOrgsReturns
	return fakeReturns.result1
}

func (fake *FakeManager) DeleteOrgsCallCount() int {
	fake.deleteOrgsMutex.RLock()
	defer fake.deleteOrgsMutex.RUnlock()
	return len(fake.deleteOrgsArgsForCall)
}

func (fake *FakeManager) DeleteOrgsCalls(stub func() error) {
	fake.deleteOrgsMutex.Lock()
	defer fake.deleteOrgsMutex.Unlock()
	fake.DeleteOrgsStub = stub
}

func (fake *FakeManager) DeleteOrgsReturns(result1 error) {
	fake.deleteOrgsMutex.Lock()
	defer fake.deleteOrgsMutex.Unlock()
	fake.DeleteOrgsStub = nil
	fake.deleteOrgsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) DeleteOrgsReturnsOnCall(i int, result1 error) {
	fake.deleteOrgsMutex.Lock()
	defer fake.deleteOrgsMutex.Unlock()
	fake.DeleteOrgsStub = nil
	if fake.deleteOrgsReturnsOnCall == nil {
		fake.deleteOrgsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteOrgsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) RenameOrg(arg1 string, arg2 string) error {
	fake.renameOrgMutex.Lock()
	ret, specificReturn := fake.renameOrgReturnsOnCall[len(fake.renameOrgArgsForCall)]
	fake.renameOrgArgsForCall = append(fake.renameOrgArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("RenameOrg", []interface{}{arg1, arg2})
	fake.renameOrgMutex.Unlock()
	if fake.RenameOrgStub != nil {
		return fake.RenameOrgStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.renameOrgReturns
	return fakeReturns.result1
}

func (fake *FakeManager) RenameOrgCallCount() int {
	fake.renameOrgMutex.RLock()
	defer fake.renameOrgMutex.RUnlock()
	return len(fake.renameOrgArgsForCall)
}

func (fake *FakeManager) RenameOrgCalls(stub func(string, string) error) {
	fake.renameOrgMutex.Lock()
	defer fake.renameOrgMutex.Unlock()
	fake.RenameOrgStub = stub
}

func (fake *FakeManager) RenameOrgArgsForCall(i int) (string, string) {
	fake.renameOrgMutex.RLock()
	defer fake.renameOrgMutex.RUnlock()
	argsForCall := fake.renameOrgArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeManager) RenameOrgReturns(result1 error) {
	fake.renameOrgMutex.Lock()
	defer fake.renameOrgMutex.Unlock()
	fake.RenameOrgStub = nil
	fake.renameOrgReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) RenameOrgReturnsOnCall(i int, result1 error) {
	fake.renameOrgMutex.Lock()
	defer fake.renameOrgMutex.Unlock()
	fake.RenameOrgStub = nil
	if fake.renameOrgReturnsOnCall == nil {
		fake.renameOrgReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.renameOrgReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) UpdateOrg(arg1 string, arg2 cfclient.OrgRequest) (cfclient.Org, error) {
	fake.updateOrgMutex.Lock()
	ret, specificReturn := fake.updateOrgReturnsOnCall[len(fake.updateOrgArgsForCall)]
	fake.updateOrgArgsForCall = append(fake.updateOrgArgsForCall, struct {
		arg1 string
		arg2 cfclient.OrgRequest
	}{arg1, arg2})
	fake.recordInvocation("UpdateOrg", []interface{}{arg1, arg2})
	fake.updateOrgMutex.Unlock()
	if fake.UpdateOrgStub != nil {
		return fake.UpdateOrgStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.updateOrgReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeManager) UpdateOrgCallCount() int {
	fake.updateOrgMutex.RLock()
	defer fake.updateOrgMutex.RUnlock()
	return len(fake.updateOrgArgsForCall)
}

func (fake *FakeManager) UpdateOrgCalls(stub func(string, cfclient.OrgRequest) (cfclient.Org, error)) {
	fake.updateOrgMutex.Lock()
	defer fake.updateOrgMutex.Unlock()
	fake.UpdateOrgStub = stub
}

func (fake *FakeManager) UpdateOrgArgsForCall(i int) (string, cfclient.OrgRequest) {
	fake.updateOrgMutex.RLock()
	defer fake.updateOrgMutex.RUnlock()
	argsForCall := fake.updateOrgArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeManager) UpdateOrgReturns(result1 cfclient.Org, result2 error) {
	fake.updateOrgMutex.Lock()
	defer fake.updateOrgMutex.Unlock()
	fake.UpdateOrgStub = nil
	fake.updateOrgReturns = struct {
		result1 cfclient.Org
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) UpdateOrgReturnsOnCall(i int, result1 cfclient.Org, result2 error) {
	fake.updateOrgMutex.Lock()
	defer fake.updateOrgMutex.Unlock()
	fake.UpdateOrgStub = nil
	if fake.updateOrgReturnsOnCall == nil {
		fake.updateOrgReturnsOnCall = make(map[int]struct {
			result1 cfclient.Org
			result2 error
		})
	}
	fake.updateOrgReturnsOnCall[i] = struct {
		result1 cfclient.Org
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) UpdateOrgsMetadata() error {
	fake.updateOrgsMetadataMutex.Lock()
	ret, specificReturn := fake.updateOrgsMetadataReturnsOnCall[len(fake.updateOrgsMetadataArgsForCall)]
	fake.updateOrgsMetadataArgsForCall = append(fake.updateOrgsMetadataArgsForCall, struct {
	}{})
	fake.recordInvocation("UpdateOrgsMetadata", []interface{}{})
	fake.updateOrgsMetadataMutex.Unlock()
	if fake.UpdateOrgsMetadataStub != nil {
		return fake.UpdateOrgsMetadataStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.updateOrgsMetadataReturns
	return fakeReturns.result1
}

func (fake *FakeManager) UpdateOrgsMetadataCallCount() int {
	fake.updateOrgsMetadataMutex.RLock()
	defer fake.updateOrgsMetadataMutex.RUnlock()
	return len(fake.updateOrgsMetadataArgsForCall)
}

func (fake *FakeManager) UpdateOrgsMetadataCalls(stub func() error) {
	fake.updateOrgsMetadataMutex.Lock()
	defer fake.updateOrgsMetadataMutex.Unlock()
	fake.UpdateOrgsMetadataStub = stub
}

func (fake *FakeManager) UpdateOrgsMetadataReturns(result1 error) {
	fake.updateOrgsMetadataMutex.Lock()
	defer fake.updateOrgsMetadataMutex.Unlock()
	fake.UpdateOrgsMetadataStub = nil
	fake.updateOrgsMetadataReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) UpdateOrgsMetadataReturnsOnCall(i int, result1 error) {
	fake.updateOrgsMetadataMutex.Lock()
	defer fake.updateOrgsMetadataMutex.Unlock()
	fake.UpdateOrgsMetadataStub = nil
	if fake.updateOrgsMetadataReturnsOnCall == nil {
		fake.updateOrgsMetadataReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateOrgsMetadataReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createOrgsMutex.RLock()
	defer fake.createOrgsMutex.RUnlock()
	fake.deleteOrgsMutex.RLock()
	defer fake.deleteOrgsMutex.RUnlock()
	fake.renameOrgMutex.RLock()
	defer fake.renameOrgMutex.RUnlock()
	fake.updateOrgMutex.RLock()
	defer fake.updateOrgMutex.RUnlock()
	fake.updateOrgsMetadataMutex.RLock()
	defer fake.updateOrgsMetadataMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeManager) recordInvocation(key string, args []interface{}) {
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

var _ organization.Manager = new(FakeManager)
