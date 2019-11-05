// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/keyval/keyval.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockKeyVal is a mock of KeyVal interface
type MockKeyVal struct {
	ctrl     *gomock.Controller
	recorder *MockKeyValMockRecorder
}

// MockKeyValMockRecorder is the mock recorder for MockKeyVal
type MockKeyValMockRecorder struct {
	mock *MockKeyVal
}

// NewMockKeyVal creates a new mock instance
func NewMockKeyVal(ctrl *gomock.Controller) *MockKeyVal {
	mock := &MockKeyVal{ctrl: ctrl}
	mock.recorder = &MockKeyValMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockKeyVal) EXPECT() *MockKeyValMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockKeyVal) Get(key string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockKeyValMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockKeyVal)(nil).Get), key)
}

// Set mocks base method
func (m *MockKeyVal) Set(key string, value []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set
func (mr *MockKeyValMockRecorder) Set(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockKeyVal)(nil).Set), key, value)
}

// Delete mocks base method
func (m *MockKeyVal) Delete(key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockKeyValMockRecorder) Delete(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockKeyVal)(nil).Delete), key)
}