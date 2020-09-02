// This file was generated by counterfeiter
package ioutil_fake

import (
	"os"
	"sync"
	"time"
)

type FakeFileInfo struct {
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct{}
	nameReturns struct {
		result1 string
	}
	SizeStub        func() int64
	sizeMutex       sync.RWMutex
	sizeArgsForCall []struct{}
	sizeReturns struct {
		result1 int64
	}
	ModeStub        func() os.FileMode
	modeMutex       sync.RWMutex
	modeArgsForCall []struct{}
	modeReturns struct {
		result1 os.FileMode
	}
	ModTimeStub        func() time.Time
	modTimeMutex       sync.RWMutex
	modTimeArgsForCall []struct{}
	modTimeReturns struct {
		result1 time.Time
	}
	IsDirStub        func() bool
	isDirMutex       sync.RWMutex
	isDirArgsForCall []struct{}
	isDirReturns struct {
		result1 bool
	}
	SysStub        func() interface{}
	sysMutex       sync.RWMutex
	sysArgsForCall []struct{}
	sysReturns struct {
		result1 interface{}
	}
}

func (fake *FakeFileInfo) Name() string {
	fake.nameMutex.Lock()
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct{}{})
	fake.nameMutex.Unlock()
	if fake.NameStub != nil {
		return fake.NameStub()
	} else {
		return fake.nameReturns.result1
	}
}

func (fake *FakeFileInfo) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeFileInfo) NameReturns(result1 string) {
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeFileInfo) Size() int64 {
	fake.sizeMutex.Lock()
	fake.sizeArgsForCall = append(fake.sizeArgsForCall, struct{}{})
	fake.sizeMutex.Unlock()
	if fake.SizeStub != nil {
		return fake.SizeStub()
	} else {
		return fake.sizeReturns.result1
	}
}

func (fake *FakeFileInfo) SizeCallCount() int {
	fake.sizeMutex.RLock()
	defer fake.sizeMutex.RUnlock()
	return len(fake.sizeArgsForCall)
}

func (fake *FakeFileInfo) SizeReturns(result1 int64) {
	fake.SizeStub = nil
	fake.sizeReturns = struct {
		result1 int64
	}{result1}
}

func (fake *FakeFileInfo) Mode() os.FileMode {
	fake.modeMutex.Lock()
	fake.modeArgsForCall = append(fake.modeArgsForCall, struct{}{})
	fake.modeMutex.Unlock()
	if fake.ModeStub != nil {
		return fake.ModeStub()
	} else {
		return fake.modeReturns.result1
	}
}

func (fake *FakeFileInfo) ModeCallCount() int {
	fake.modeMutex.RLock()
	defer fake.modeMutex.RUnlock()
	return len(fake.modeArgsForCall)
}

func (fake *FakeFileInfo) ModeReturns(result1 os.FileMode) {
	fake.ModeStub = nil
	fake.modeReturns = struct {
		result1 os.FileMode
	}{result1}
}

func (fake *FakeFileInfo) ModTime() time.Time {
	fake.modTimeMutex.Lock()
	fake.modTimeArgsForCall = append(fake.modTimeArgsForCall, struct{}{})
	fake.modTimeMutex.Unlock()
	if fake.ModTimeStub != nil {
		return fake.ModTimeStub()
	} else {
		return fake.modTimeReturns.result1
	}
}

func (fake *FakeFileInfo) ModTimeCallCount() int {
	fake.modTimeMutex.RLock()
	defer fake.modTimeMutex.RUnlock()
	return len(fake.modTimeArgsForCall)
}

func (fake *FakeFileInfo) ModTimeReturns(result1 time.Time) {
	fake.ModTimeStub = nil
	fake.modTimeReturns = struct {
		result1 time.Time
	}{result1}
}

func (fake *FakeFileInfo) IsDir() bool {
	fake.isDirMutex.Lock()
	fake.isDirArgsForCall = append(fake.isDirArgsForCall, struct{}{})
	fake.isDirMutex.Unlock()
	if fake.IsDirStub != nil {
		return fake.IsDirStub()
	} else {
		return fake.isDirReturns.result1
	}
}

func (fake *FakeFileInfo) IsDirCallCount() int {
	fake.isDirMutex.RLock()
	defer fake.isDirMutex.RUnlock()
	return len(fake.isDirArgsForCall)
}

func (fake *FakeFileInfo) IsDirReturns(result1 bool) {
	fake.IsDirStub = nil
	fake.isDirReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeFileInfo) Sys() interface{} {
	fake.sysMutex.Lock()
	fake.sysArgsForCall = append(fake.sysArgsForCall, struct{}{})
	fake.sysMutex.Unlock()
	if fake.SysStub != nil {
		return fake.SysStub()
	} else {
		return fake.sysReturns.result1
	}
}

func (fake *FakeFileInfo) SysCallCount() int {
	fake.sysMutex.RLock()
	defer fake.sysMutex.RUnlock()
	return len(fake.sysArgsForCall)
}

func (fake *FakeFileInfo) SysReturns(result1 interface{}) {
	fake.SysStub = nil
	fake.sysReturns = struct {
		result1 interface{}
	}{result1}
}

var _ os.FileInfo = new(FakeFileInfo)