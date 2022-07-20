// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/remiphilippe/go-etcd-lock/lock (interfaces: Lock)

// Package lockmock is a generated GoMock package.
package lockmock

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockLock is a mock of Lock interface
type MockLock struct {
	ctrl     *gomock.Controller
	recorder *MockLockMockRecorder
}

// MockLockMockRecorder is the mock recorder for MockLock
type MockLockMockRecorder struct {
	mock *MockLock
}

// NewMockLock creates a new mock instance
func NewMockLock(ctrl *gomock.Controller) *MockLock {
	mock := &MockLock{ctrl: ctrl}
	mock.recorder = &MockLockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLock) EXPECT() *MockLockMockRecorder {
	return m.recorder
}

// Release mocks base method
func (m *MockLock) Release() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Release")
	ret0, _ := ret[0].(error)
	return ret0
}

// Release indicates an expected call of Release
func (mr *MockLockMockRecorder) Release() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Release", reflect.TypeOf((*MockLock)(nil).Release))
}
