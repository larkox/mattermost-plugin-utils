// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/larkox/mattermost-plugin-utils/bot (interfaces: Admin)

// Package mock_bot is a generated GoMock package.
package mock_bot

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAdmin is a mock of Admin interface
type MockAdmin struct {
	ctrl     *gomock.Controller
	recorder *MockAdminMockRecorder
}

// MockAdminMockRecorder is the mock recorder for MockAdmin
type MockAdminMockRecorder struct {
	mock *MockAdmin
}

// NewMockAdmin creates a new mock instance
func NewMockAdmin(ctrl *gomock.Controller) *MockAdmin {
	mock := &MockAdmin{ctrl: ctrl}
	mock.recorder = &MockAdminMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAdmin) EXPECT() *MockAdminMockRecorder {
	return m.recorder
}

// DMAdmins mocks base method
func (m *MockAdmin) DMAdmins(arg0 string, arg1 ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DMAdmins", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DMAdmins indicates an expected call of DMAdmins
func (mr *MockAdminMockRecorder) DMAdmins(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DMAdmins", reflect.TypeOf((*MockAdmin)(nil).DMAdmins), varargs...)
}

// IsUserAdmin mocks base method
func (m *MockAdmin) IsUserAdmin(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUserAdmin", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsUserAdmin indicates an expected call of IsUserAdmin
func (mr *MockAdminMockRecorder) IsUserAdmin(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUserAdmin", reflect.TypeOf((*MockAdmin)(nil).IsUserAdmin), arg0)
}
