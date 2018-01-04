// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Scalingo/go-etcd-lock/lock (interfaces: Locker)

// Package lockmock is a generated GoMock package.
package lockmock

import (
	lock "github.com/Scalingo/go-etcd-lock/lock"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockLocker is a mock of Locker interface
type MockLocker struct {
	ctrl     *gomock.Controller
	recorder *MockLockerMockRecorder
}

// MockLockerMockRecorder is the mock recorder for MockLocker
type MockLockerMockRecorder struct {
	mock *MockLocker
}

// NewMockLocker creates a new mock instance
func NewMockLocker(ctrl *gomock.Controller) *MockLocker {
	mock := &MockLocker{ctrl: ctrl}
	mock.recorder = &MockLockerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLocker) EXPECT() *MockLockerMockRecorder {
	return m.recorder
}

// Acquire mocks base method
func (m *MockLocker) Acquire(arg0 string, arg1 int) (lock.Lock, error) {
	ret := m.ctrl.Call(m, "Acquire", arg0, arg1)
	ret0, _ := ret[0].(lock.Lock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Acquire indicates an expected call of Acquire
func (mr *MockLockerMockRecorder) Acquire(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Acquire", reflect.TypeOf((*MockLocker)(nil).Acquire), arg0, arg1)
}

// Wait mocks base method
func (m *MockLocker) Wait(arg0 string) error {
	ret := m.ctrl.Call(m, "Wait", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Wait indicates an expected call of Wait
func (mr *MockLockerMockRecorder) Wait(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockLocker)(nil).Wait), arg0)
}

// WaitAcquire mocks base method
func (m *MockLocker) WaitAcquire(arg0 string, arg1 int) (lock.Lock, error) {
	ret := m.ctrl.Call(m, "WaitAcquire", arg0, arg1)
	ret0, _ := ret[0].(lock.Lock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitAcquire indicates an expected call of WaitAcquire
func (mr *MockLockerMockRecorder) WaitAcquire(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitAcquire", reflect.TypeOf((*MockLocker)(nil).WaitAcquire), arg0, arg1)
}
