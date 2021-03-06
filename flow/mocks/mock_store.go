// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/larkox/mattermost-plugin-utils/flow (interfaces: FlowStore)

// Package mock_flow is a generated GoMock package.
package mock_flow

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockFlowStore is a mock of FlowStore interface
type MockFlowStore struct {
	ctrl     *gomock.Controller
	recorder *MockFlowStoreMockRecorder
}

// MockFlowStoreMockRecorder is the mock recorder for MockFlowStore
type MockFlowStoreMockRecorder struct {
	mock *MockFlowStore
}

// NewMockFlowStore creates a new mock instance
func NewMockFlowStore(ctrl *gomock.Controller) *MockFlowStore {
	mock := &MockFlowStore{ctrl: ctrl}
	mock.recorder = &MockFlowStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFlowStore) EXPECT() *MockFlowStoreMockRecorder {
	return m.recorder
}

// DeleteCurrentStep mocks base method
func (m *MockFlowStore) DeleteCurrentStep(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCurrentStep", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCurrentStep indicates an expected call of DeleteCurrentStep
func (mr *MockFlowStoreMockRecorder) DeleteCurrentStep(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCurrentStep", reflect.TypeOf((*MockFlowStore)(nil).DeleteCurrentStep), arg0)
}

// GetCurrentStep mocks base method
func (m *MockFlowStore) GetCurrentStep(arg0 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentStep", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentStep indicates an expected call of GetCurrentStep
func (mr *MockFlowStoreMockRecorder) GetCurrentStep(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentStep", reflect.TypeOf((*MockFlowStore)(nil).GetCurrentStep), arg0)
}

// GetPostID mocks base method
func (m *MockFlowStore) GetPostID(arg0, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostID", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostID indicates an expected call of GetPostID
func (mr *MockFlowStoreMockRecorder) GetPostID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostID", reflect.TypeOf((*MockFlowStore)(nil).GetPostID), arg0, arg1)
}

// RemovePostID mocks base method
func (m *MockFlowStore) RemovePostID(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemovePostID", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemovePostID indicates an expected call of RemovePostID
func (mr *MockFlowStoreMockRecorder) RemovePostID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemovePostID", reflect.TypeOf((*MockFlowStore)(nil).RemovePostID), arg0, arg1)
}

// SetCurrentStep mocks base method
func (m *MockFlowStore) SetCurrentStep(arg0 string, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCurrentStep", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetCurrentStep indicates an expected call of SetCurrentStep
func (mr *MockFlowStoreMockRecorder) SetCurrentStep(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCurrentStep", reflect.TypeOf((*MockFlowStore)(nil).SetCurrentStep), arg0, arg1)
}

// SetPostID mocks base method
func (m *MockFlowStore) SetPostID(arg0, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPostID", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPostID indicates an expected call of SetPostID
func (mr *MockFlowStoreMockRecorder) SetPostID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPostID", reflect.TypeOf((*MockFlowStore)(nil).SetPostID), arg0, arg1, arg2)
}

// SetProperty mocks base method
func (m *MockFlowStore) SetProperty(arg0, arg1 string, arg2 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetProperty", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetProperty indicates an expected call of SetProperty
func (mr *MockFlowStoreMockRecorder) SetProperty(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetProperty", reflect.TypeOf((*MockFlowStore)(nil).SetProperty), arg0, arg1, arg2)
}
