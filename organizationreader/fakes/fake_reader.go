// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/cloudfoundry-community/go-cfclient/v3/resource"
	"github.com/vmwarepivotallabs/cf-mgmt/organizationreader"
)

type FakeReader struct {
	AddOrgToListStub        func(*resource.Organization)
	addOrgToListMutex       sync.RWMutex
	addOrgToListArgsForCall []struct {
		arg1 *resource.Organization
	}
	ClearOrgListStub        func()
	clearOrgListMutex       sync.RWMutex
	clearOrgListArgsForCall []struct {
	}
	FindOrgStub        func(string) (*resource.Organization, error)
	findOrgMutex       sync.RWMutex
	findOrgArgsForCall []struct {
		arg1 string
	}
	findOrgReturns struct {
		result1 *resource.Organization
		result2 error
	}
	findOrgReturnsOnCall map[int]struct {
		result1 *resource.Organization
		result2 error
	}
	FindOrgByGUIDStub        func(string) (*resource.Organization, error)
	findOrgByGUIDMutex       sync.RWMutex
	findOrgByGUIDArgsForCall []struct {
		arg1 string
	}
	findOrgByGUIDReturns struct {
		result1 *resource.Organization
		result2 error
	}
	findOrgByGUIDReturnsOnCall map[int]struct {
		result1 *resource.Organization
		result2 error
	}
	GetDefaultIsolationSegmentStub        func(*resource.Organization) (string, error)
	getDefaultIsolationSegmentMutex       sync.RWMutex
	getDefaultIsolationSegmentArgsForCall []struct {
		arg1 *resource.Organization
	}
	getDefaultIsolationSegmentReturns struct {
		result1 string
		result2 error
	}
	getDefaultIsolationSegmentReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GetOrgGUIDStub        func(string) (string, error)
	getOrgGUIDMutex       sync.RWMutex
	getOrgGUIDArgsForCall []struct {
		arg1 string
	}
	getOrgGUIDReturns struct {
		result1 string
		result2 error
	}
	getOrgGUIDReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	ListOrgsStub        func() ([]*resource.Organization, error)
	listOrgsMutex       sync.RWMutex
	listOrgsArgsForCall []struct {
	}
	listOrgsReturns struct {
		result1 []*resource.Organization
		result2 error
	}
	listOrgsReturnsOnCall map[int]struct {
		result1 []*resource.Organization
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeReader) AddOrgToList(arg1 *resource.Organization) {
	fake.addOrgToListMutex.Lock()
	fake.addOrgToListArgsForCall = append(fake.addOrgToListArgsForCall, struct {
		arg1 *resource.Organization
	}{arg1})
	stub := fake.AddOrgToListStub
	fake.recordInvocation("AddOrgToList", []interface{}{arg1})
	fake.addOrgToListMutex.Unlock()
	if stub != nil {
		fake.AddOrgToListStub(arg1)
	}
}

func (fake *FakeReader) AddOrgToListCallCount() int {
	fake.addOrgToListMutex.RLock()
	defer fake.addOrgToListMutex.RUnlock()
	return len(fake.addOrgToListArgsForCall)
}

func (fake *FakeReader) AddOrgToListCalls(stub func(*resource.Organization)) {
	fake.addOrgToListMutex.Lock()
	defer fake.addOrgToListMutex.Unlock()
	fake.AddOrgToListStub = stub
}

func (fake *FakeReader) AddOrgToListArgsForCall(i int) *resource.Organization {
	fake.addOrgToListMutex.RLock()
	defer fake.addOrgToListMutex.RUnlock()
	argsForCall := fake.addOrgToListArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeReader) ClearOrgList() {
	fake.clearOrgListMutex.Lock()
	fake.clearOrgListArgsForCall = append(fake.clearOrgListArgsForCall, struct {
	}{})
	stub := fake.ClearOrgListStub
	fake.recordInvocation("ClearOrgList", []interface{}{})
	fake.clearOrgListMutex.Unlock()
	if stub != nil {
		fake.ClearOrgListStub()
	}
}

func (fake *FakeReader) ClearOrgListCallCount() int {
	fake.clearOrgListMutex.RLock()
	defer fake.clearOrgListMutex.RUnlock()
	return len(fake.clearOrgListArgsForCall)
}

func (fake *FakeReader) ClearOrgListCalls(stub func()) {
	fake.clearOrgListMutex.Lock()
	defer fake.clearOrgListMutex.Unlock()
	fake.ClearOrgListStub = stub
}

func (fake *FakeReader) FindOrg(arg1 string) (*resource.Organization, error) {
	fake.findOrgMutex.Lock()
	ret, specificReturn := fake.findOrgReturnsOnCall[len(fake.findOrgArgsForCall)]
	fake.findOrgArgsForCall = append(fake.findOrgArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.FindOrgStub
	fakeReturns := fake.findOrgReturns
	fake.recordInvocation("FindOrg", []interface{}{arg1})
	fake.findOrgMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeReader) FindOrgCallCount() int {
	fake.findOrgMutex.RLock()
	defer fake.findOrgMutex.RUnlock()
	return len(fake.findOrgArgsForCall)
}

func (fake *FakeReader) FindOrgCalls(stub func(string) (*resource.Organization, error)) {
	fake.findOrgMutex.Lock()
	defer fake.findOrgMutex.Unlock()
	fake.FindOrgStub = stub
}

func (fake *FakeReader) FindOrgArgsForCall(i int) string {
	fake.findOrgMutex.RLock()
	defer fake.findOrgMutex.RUnlock()
	argsForCall := fake.findOrgArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeReader) FindOrgReturns(result1 *resource.Organization, result2 error) {
	fake.findOrgMutex.Lock()
	defer fake.findOrgMutex.Unlock()
	fake.FindOrgStub = nil
	fake.findOrgReturns = struct {
		result1 *resource.Organization
		result2 error
	}{result1, result2}
}

func (fake *FakeReader) FindOrgReturnsOnCall(i int, result1 *resource.Organization, result2 error) {
	fake.findOrgMutex.Lock()
	defer fake.findOrgMutex.Unlock()
	fake.FindOrgStub = nil
	if fake.findOrgReturnsOnCall == nil {
		fake.findOrgReturnsOnCall = make(map[int]struct {
			result1 *resource.Organization
			result2 error
		})
	}
	fake.findOrgReturnsOnCall[i] = struct {
		result1 *resource.Organization
		result2 error
	}{result1, result2}
}

func (fake *FakeReader) FindOrgByGUID(arg1 string) (*resource.Organization, error) {
	fake.findOrgByGUIDMutex.Lock()
	ret, specificReturn := fake.findOrgByGUIDReturnsOnCall[len(fake.findOrgByGUIDArgsForCall)]
	fake.findOrgByGUIDArgsForCall = append(fake.findOrgByGUIDArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.FindOrgByGUIDStub
	fakeReturns := fake.findOrgByGUIDReturns
	fake.recordInvocation("FindOrgByGUID", []interface{}{arg1})
	fake.findOrgByGUIDMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeReader) FindOrgByGUIDCallCount() int {
	fake.findOrgByGUIDMutex.RLock()
	defer fake.findOrgByGUIDMutex.RUnlock()
	return len(fake.findOrgByGUIDArgsForCall)
}

func (fake *FakeReader) FindOrgByGUIDCalls(stub func(string) (*resource.Organization, error)) {
	fake.findOrgByGUIDMutex.Lock()
	defer fake.findOrgByGUIDMutex.Unlock()
	fake.FindOrgByGUIDStub = stub
}

func (fake *FakeReader) FindOrgByGUIDArgsForCall(i int) string {
	fake.findOrgByGUIDMutex.RLock()
	defer fake.findOrgByGUIDMutex.RUnlock()
	argsForCall := fake.findOrgByGUIDArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeReader) FindOrgByGUIDReturns(result1 *resource.Organization, result2 error) {
	fake.findOrgByGUIDMutex.Lock()
	defer fake.findOrgByGUIDMutex.Unlock()
	fake.FindOrgByGUIDStub = nil
	fake.findOrgByGUIDReturns = struct {
		result1 *resource.Organization
		result2 error
	}{result1, result2}
}

func (fake *FakeReader) FindOrgByGUIDReturnsOnCall(i int, result1 *resource.Organization, result2 error) {
	fake.findOrgByGUIDMutex.Lock()
	defer fake.findOrgByGUIDMutex.Unlock()
	fake.FindOrgByGUIDStub = nil
	if fake.findOrgByGUIDReturnsOnCall == nil {
		fake.findOrgByGUIDReturnsOnCall = make(map[int]struct {
			result1 *resource.Organization
			result2 error
		})
	}
	fake.findOrgByGUIDReturnsOnCall[i] = struct {
		result1 *resource.Organization
		result2 error
	}{result1, result2}
}

func (fake *FakeReader) GetDefaultIsolationSegment(arg1 *resource.Organization) (string, error) {
	fake.getDefaultIsolationSegmentMutex.Lock()
	ret, specificReturn := fake.getDefaultIsolationSegmentReturnsOnCall[len(fake.getDefaultIsolationSegmentArgsForCall)]
	fake.getDefaultIsolationSegmentArgsForCall = append(fake.getDefaultIsolationSegmentArgsForCall, struct {
		arg1 *resource.Organization
	}{arg1})
	stub := fake.GetDefaultIsolationSegmentStub
	fakeReturns := fake.getDefaultIsolationSegmentReturns
	fake.recordInvocation("GetDefaultIsolationSegment", []interface{}{arg1})
	fake.getDefaultIsolationSegmentMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeReader) GetDefaultIsolationSegmentCallCount() int {
	fake.getDefaultIsolationSegmentMutex.RLock()
	defer fake.getDefaultIsolationSegmentMutex.RUnlock()
	return len(fake.getDefaultIsolationSegmentArgsForCall)
}

func (fake *FakeReader) GetDefaultIsolationSegmentCalls(stub func(*resource.Organization) (string, error)) {
	fake.getDefaultIsolationSegmentMutex.Lock()
	defer fake.getDefaultIsolationSegmentMutex.Unlock()
	fake.GetDefaultIsolationSegmentStub = stub
}

func (fake *FakeReader) GetDefaultIsolationSegmentArgsForCall(i int) *resource.Organization {
	fake.getDefaultIsolationSegmentMutex.RLock()
	defer fake.getDefaultIsolationSegmentMutex.RUnlock()
	argsForCall := fake.getDefaultIsolationSegmentArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeReader) GetDefaultIsolationSegmentReturns(result1 string, result2 error) {
	fake.getDefaultIsolationSegmentMutex.Lock()
	defer fake.getDefaultIsolationSegmentMutex.Unlock()
	fake.GetDefaultIsolationSegmentStub = nil
	fake.getDefaultIsolationSegmentReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeReader) GetDefaultIsolationSegmentReturnsOnCall(i int, result1 string, result2 error) {
	fake.getDefaultIsolationSegmentMutex.Lock()
	defer fake.getDefaultIsolationSegmentMutex.Unlock()
	fake.GetDefaultIsolationSegmentStub = nil
	if fake.getDefaultIsolationSegmentReturnsOnCall == nil {
		fake.getDefaultIsolationSegmentReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getDefaultIsolationSegmentReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeReader) GetOrgGUID(arg1 string) (string, error) {
	fake.getOrgGUIDMutex.Lock()
	ret, specificReturn := fake.getOrgGUIDReturnsOnCall[len(fake.getOrgGUIDArgsForCall)]
	fake.getOrgGUIDArgsForCall = append(fake.getOrgGUIDArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetOrgGUIDStub
	fakeReturns := fake.getOrgGUIDReturns
	fake.recordInvocation("GetOrgGUID", []interface{}{arg1})
	fake.getOrgGUIDMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeReader) GetOrgGUIDCallCount() int {
	fake.getOrgGUIDMutex.RLock()
	defer fake.getOrgGUIDMutex.RUnlock()
	return len(fake.getOrgGUIDArgsForCall)
}

func (fake *FakeReader) GetOrgGUIDCalls(stub func(string) (string, error)) {
	fake.getOrgGUIDMutex.Lock()
	defer fake.getOrgGUIDMutex.Unlock()
	fake.GetOrgGUIDStub = stub
}

func (fake *FakeReader) GetOrgGUIDArgsForCall(i int) string {
	fake.getOrgGUIDMutex.RLock()
	defer fake.getOrgGUIDMutex.RUnlock()
	argsForCall := fake.getOrgGUIDArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeReader) GetOrgGUIDReturns(result1 string, result2 error) {
	fake.getOrgGUIDMutex.Lock()
	defer fake.getOrgGUIDMutex.Unlock()
	fake.GetOrgGUIDStub = nil
	fake.getOrgGUIDReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeReader) GetOrgGUIDReturnsOnCall(i int, result1 string, result2 error) {
	fake.getOrgGUIDMutex.Lock()
	defer fake.getOrgGUIDMutex.Unlock()
	fake.GetOrgGUIDStub = nil
	if fake.getOrgGUIDReturnsOnCall == nil {
		fake.getOrgGUIDReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getOrgGUIDReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeReader) ListOrgs() ([]*resource.Organization, error) {
	fake.listOrgsMutex.Lock()
	ret, specificReturn := fake.listOrgsReturnsOnCall[len(fake.listOrgsArgsForCall)]
	fake.listOrgsArgsForCall = append(fake.listOrgsArgsForCall, struct {
	}{})
	stub := fake.ListOrgsStub
	fakeReturns := fake.listOrgsReturns
	fake.recordInvocation("ListOrgs", []interface{}{})
	fake.listOrgsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeReader) ListOrgsCallCount() int {
	fake.listOrgsMutex.RLock()
	defer fake.listOrgsMutex.RUnlock()
	return len(fake.listOrgsArgsForCall)
}

func (fake *FakeReader) ListOrgsCalls(stub func() ([]*resource.Organization, error)) {
	fake.listOrgsMutex.Lock()
	defer fake.listOrgsMutex.Unlock()
	fake.ListOrgsStub = stub
}

func (fake *FakeReader) ListOrgsReturns(result1 []*resource.Organization, result2 error) {
	fake.listOrgsMutex.Lock()
	defer fake.listOrgsMutex.Unlock()
	fake.ListOrgsStub = nil
	fake.listOrgsReturns = struct {
		result1 []*resource.Organization
		result2 error
	}{result1, result2}
}

func (fake *FakeReader) ListOrgsReturnsOnCall(i int, result1 []*resource.Organization, result2 error) {
	fake.listOrgsMutex.Lock()
	defer fake.listOrgsMutex.Unlock()
	fake.ListOrgsStub = nil
	if fake.listOrgsReturnsOnCall == nil {
		fake.listOrgsReturnsOnCall = make(map[int]struct {
			result1 []*resource.Organization
			result2 error
		})
	}
	fake.listOrgsReturnsOnCall[i] = struct {
		result1 []*resource.Organization
		result2 error
	}{result1, result2}
}

func (fake *FakeReader) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addOrgToListMutex.RLock()
	defer fake.addOrgToListMutex.RUnlock()
	fake.clearOrgListMutex.RLock()
	defer fake.clearOrgListMutex.RUnlock()
	fake.findOrgMutex.RLock()
	defer fake.findOrgMutex.RUnlock()
	fake.findOrgByGUIDMutex.RLock()
	defer fake.findOrgByGUIDMutex.RUnlock()
	fake.getDefaultIsolationSegmentMutex.RLock()
	defer fake.getDefaultIsolationSegmentMutex.RUnlock()
	fake.getOrgGUIDMutex.RLock()
	defer fake.getOrgGUIDMutex.RUnlock()
	fake.listOrgsMutex.RLock()
	defer fake.listOrgsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeReader) recordInvocation(key string, args []interface{}) {
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

var _ organizationreader.Reader = new(FakeReader)
