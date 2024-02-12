// Code generated by MockGen. DO NOT EDIT.
// Source: internal/transport/handlers.go
//
// Generated by this command:
//
//	mockgen -source=internal/transport/handlers.go -destination=internal/transport/mock_handlers.go -package=transport
//
// Package transport is a generated GoMock package.
package transport

import (
	reflect "reflect"
	models "github.com/tahanasir/service-catalog/internal/models"

	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockdbStorage is a mock of dbStorage interface.
type MockdbStorage struct {
	ctrl     *gomock.Controller
	recorder *MockdbStorageMockRecorder
}

// MockdbStorageMockRecorder is the mock recorder for MockdbStorage.
type MockdbStorageMockRecorder struct {
	mock *MockdbStorage
}

// NewMockdbStorage creates a new mock instance.
func NewMockdbStorage(ctrl *gomock.Controller) *MockdbStorage {
	mock := &MockdbStorage{ctrl: ctrl}
	mock.recorder = &MockdbStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockdbStorage) EXPECT() *MockdbStorageMockRecorder {
	return m.recorder
}

// GetAllServices mocks base method.
func (m *MockdbStorage) GetAllServices(arg0, arg1 int, arg2 string) ([]models.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllServices", arg0, arg1, arg2)
	ret0, _ := ret[0].([]models.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllServices indicates an expected call of GetAllServices.
func (mr *MockdbStorageMockRecorder) GetAllServices(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllServices", reflect.TypeOf((*MockdbStorage)(nil).GetAllServices), arg0, arg1, arg2)
}

// GetService mocks base method.
func (m *MockdbStorage) GetService(arg0 uuid.UUID) (models.SingleService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetService", arg0)
	ret0, _ := ret[0].(models.SingleService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetService indicates an expected call of GetService.
func (mr *MockdbStorageMockRecorder) GetService(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetService", reflect.TypeOf((*MockdbStorage)(nil).GetService), arg0)
}

// GetTotalServicesCount mocks base method.
func (m *MockdbStorage) GetTotalServicesCount(arg0 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalServicesCount", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTotalServicesCount indicates an expected call of GetTotalServicesCount.
func (mr *MockdbStorageMockRecorder) GetTotalServicesCount(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalServicesCount", reflect.TypeOf((*MockdbStorage)(nil).GetTotalServicesCount), arg0)
}
