// This file was generated by counterfeiter
package fake_app_examiner

import (
	"sync"

	"github.com/pivotal-cf-experimental/lattice-cli/app_examiner"
)

type FakeAppExaminer struct {
	ListAppsStub        func() ([]app_examiner.AppInfo, error)
	listAppsMutex       sync.RWMutex
	listAppsArgsForCall []struct{}
	listAppsReturns     struct {
		result1 []app_examiner.AppInfo
		result2 error
	}
	ListCellsStub        func() ([]app_examiner.CellInfo, error)
	listCellsMutex       sync.RWMutex
	listCellsArgsForCall []struct{}
	listCellsReturns     struct {
		result1 []app_examiner.CellInfo
		result2 error
	}
}

func (fake *FakeAppExaminer) ListApps() ([]app_examiner.AppInfo, error) {
	fake.listAppsMutex.Lock()
	fake.listAppsArgsForCall = append(fake.listAppsArgsForCall, struct{}{})
	fake.listAppsMutex.Unlock()
	if fake.ListAppsStub != nil {
		return fake.ListAppsStub()
	} else {
		return fake.listAppsReturns.result1, fake.listAppsReturns.result2
	}
}

func (fake *FakeAppExaminer) ListAppsCallCount() int {
	fake.listAppsMutex.RLock()
	defer fake.listAppsMutex.RUnlock()
	return len(fake.listAppsArgsForCall)
}

func (fake *FakeAppExaminer) ListAppsReturns(result1 []app_examiner.AppInfo, result2 error) {
	fake.ListAppsStub = nil
	fake.listAppsReturns = struct {
		result1 []app_examiner.AppInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeAppExaminer) ListCells() ([]app_examiner.CellInfo, error) {
	fake.listCellsMutex.Lock()
	fake.listCellsArgsForCall = append(fake.listCellsArgsForCall, struct{}{})
	fake.listCellsMutex.Unlock()
	if fake.ListCellsStub != nil {
		return fake.ListCellsStub()
	} else {
		return fake.listCellsReturns.result1, fake.listCellsReturns.result2
	}
}

func (fake *FakeAppExaminer) ListCellsCallCount() int {
	fake.listCellsMutex.RLock()
	defer fake.listCellsMutex.RUnlock()
	return len(fake.listCellsArgsForCall)
}

func (fake *FakeAppExaminer) ListCellsReturns(result1 []app_examiner.CellInfo, result2 error) {
	fake.ListCellsStub = nil
	fake.listCellsReturns = struct {
		result1 []app_examiner.CellInfo
		result2 error
	}{result1, result2}
}

var _ app_examiner.AppExaminer = new(FakeAppExaminer)