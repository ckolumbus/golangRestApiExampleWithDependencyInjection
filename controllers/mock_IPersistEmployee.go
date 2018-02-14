// Package mock_persistence is a generated GoMock package.
package controllers

import (
	reflect "reflect"

	dto "github.com/ckolumbus/golangRestApiExampleWithDependencyInjection/dto"
	gomock "github.com/golang/mock/gomock"
)

// MockIEmployeePersist is a mock of IEmployeePersist interface
type MockIEmployeePersist struct {
	ctrl     *gomock.Controller
	recorder *MockIEmployeePersistMockRecorder
}

// MockIEmployeePersistMockRecorder is the mock recorder for MockIEmployeePersist
type MockIEmployeePersistMockRecorder struct {
	mock *MockIEmployeePersist
}

// NewMockIEmployeePersist creates a new mock instance
func NewMockIEmployeePersist(ctrl *gomock.Controller) *MockIEmployeePersist {
	mock := &MockIEmployeePersist{ctrl: ctrl}
	mock.recorder = &MockIEmployeePersistMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIEmployeePersist) EXPECT() *MockIEmployeePersistMockRecorder {
	return m.recorder
}

// Save mocks base method
func (m *MockIEmployeePersist) Save(arg0 *dto.Employee) (string, error) {
	ret := m.ctrl.Call(m, "Save", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save
func (mr *MockIEmployeePersistMockRecorder) Save(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockIEmployeePersist)(nil).Save), arg0)
}

// Delete mocks base method
func (m *MockIEmployeePersist) Delete(arg0 string) (string, error) {
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockIEmployeePersistMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIEmployeePersist)(nil).Delete), arg0)
}

// Get mocks base method
func (m *MockIEmployeePersist) Get(arg0 string) (*dto.Employee, error) {
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*dto.Employee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockIEmployeePersistMockRecorder) Get(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIEmployeePersist)(nil).Get), arg0)
}
