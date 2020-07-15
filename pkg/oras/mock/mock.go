// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/oras/interface.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	model "github.com/kleveross/ormb/pkg/model"
	oci "github.com/kleveross/ormb/pkg/oci"
	reflect "reflect"
)

// MockInterface is a mock of Interface interface
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// Login mocks base method
func (m *MockInterface) Login(hostname, username, password string, insecure bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", hostname, username, password, insecure)
	ret0, _ := ret[0].(error)
	return ret0
}

// Login indicates an expected call of Login
func (mr *MockInterfaceMockRecorder) Login(hostname, username, password, insecure interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockInterface)(nil).Login), hostname, username, password, insecure)
}

// Logout mocks base method
func (m *MockInterface) Logout(hostname string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logout", hostname)
	ret0, _ := ret[0].(error)
	return ret0
}

// Logout indicates an expected call of Logout
func (mr *MockInterfaceMockRecorder) Logout(hostname interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockInterface)(nil).Logout), hostname)
}

// SaveModel mocks base method
func (m *MockInterface) SaveModel(ch *model.Model, ref *oci.Reference) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveModel", ch, ref)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveModel indicates an expected call of SaveModel
func (mr *MockInterfaceMockRecorder) SaveModel(ch, ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveModel", reflect.TypeOf((*MockInterface)(nil).SaveModel), ch, ref)
}

// PushModel mocks base method
func (m *MockInterface) PushModel(ref *oci.Reference) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PushModel", ref)
	ret0, _ := ret[0].(error)
	return ret0
}

// PushModel indicates an expected call of PushModel
func (mr *MockInterfaceMockRecorder) PushModel(ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushModel", reflect.TypeOf((*MockInterface)(nil).PushModel), ref)
}

// RemoveModel mocks base method
func (m *MockInterface) RemoveModel(ref *oci.Reference) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveModel", ref)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveModel indicates an expected call of RemoveModel
func (mr *MockInterfaceMockRecorder) RemoveModel(ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveModel", reflect.TypeOf((*MockInterface)(nil).RemoveModel), ref)
}

// PullModel mocks base method
func (m *MockInterface) PullModel(ref *oci.Reference) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PullModel", ref)
	ret0, _ := ret[0].(error)
	return ret0
}

// PullModel indicates an expected call of PullModel
func (mr *MockInterfaceMockRecorder) PullModel(ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PullModel", reflect.TypeOf((*MockInterface)(nil).PullModel), ref)
}

// LoadModel mocks base method
func (m *MockInterface) LoadModel(ref *oci.Reference) (*model.Model, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadModel", ref)
	ret0, _ := ret[0].(*model.Model)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadModel indicates an expected call of LoadModel
func (mr *MockInterfaceMockRecorder) LoadModel(ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadModel", reflect.TypeOf((*MockInterface)(nil).LoadModel), ref)
}
