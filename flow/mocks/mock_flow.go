// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/larkox/mattermost-plugin-utils/flow (interfaces: Flow)

// Package mock_flow is a generated GoMock package.
package mock_flow

import (
	gomock "github.com/golang/mock/gomock"
	steps "github.com/larkox/mattermost-plugin-utils/flow/steps"
	reflect "reflect"
)

// MockFlow is a mock of Flow interface
type MockFlow struct {
	ctrl     *gomock.Controller
	recorder *MockFlowMockRecorder
}

// MockFlowMockRecorder is the mock recorder for MockFlow
type MockFlowMockRecorder struct {
	mock *MockFlow
}

// NewMockFlow creates a new mock instance
func NewMockFlow(ctrl *gomock.Controller) *MockFlow {
	mock := &MockFlow{ctrl: ctrl}
	mock.recorder = &MockFlowMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFlow) EXPECT() *MockFlowMockRecorder {
	return m.recorder
}

// FlowDone mocks base method
func (m *MockFlow) FlowDone(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FlowDone", arg0)
}

// FlowDone indicates an expected call of FlowDone
func (mr *MockFlowMockRecorder) FlowDone(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlowDone", reflect.TypeOf((*MockFlow)(nil).FlowDone), arg0)
}

// Length mocks base method
func (m *MockFlow) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length
func (mr *MockFlowMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockFlow)(nil).Length))
}

// Step mocks base method
func (m *MockFlow) Step(arg0 int) steps.Step {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Step", arg0)
	ret0, _ := ret[0].(steps.Step)
	return ret0
}

// Step indicates an expected call of Step
func (mr *MockFlowMockRecorder) Step(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Step", reflect.TypeOf((*MockFlow)(nil).Step), arg0)
}

// StepDone mocks base method
func (m *MockFlow) StepDone(arg0 string, arg1 int, arg2 interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StepDone", arg0, arg1, arg2)
}

// StepDone indicates an expected call of StepDone
func (mr *MockFlowMockRecorder) StepDone(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StepDone", reflect.TypeOf((*MockFlow)(nil).StepDone), arg0, arg1, arg2)
}

// Steps mocks base method
func (m *MockFlow) Steps() []steps.Step {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Steps")
	ret0, _ := ret[0].([]steps.Step)
	return ret0
}

// Steps indicates an expected call of Steps
func (mr *MockFlowMockRecorder) Steps() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Steps", reflect.TypeOf((*MockFlow)(nil).Steps))
}

// URL mocks base method
func (m *MockFlow) URL() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "URL")
	ret0, _ := ret[0].(string)
	return ret0
}

// URL indicates an expected call of URL
func (mr *MockFlowMockRecorder) URL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "URL", reflect.TypeOf((*MockFlow)(nil).URL))
}
