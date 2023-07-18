// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/belamov/ypgo-metrics/internal/app/storage (interfaces: Repository)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	models "github.com/belamov/ypgo-metrics/internal/app/models"
	gomock "go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// UpdateCounterMetric mocks base method.
func (m *MockRepository) UpdateCounterMetric(arg0 context.Context, arg1 models.CounterMetric) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCounterMetric", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCounterMetric indicates an expected call of UpdateCounterMetric.
func (mr *MockRepositoryMockRecorder) UpdateCounterMetric(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCounterMetric", reflect.TypeOf((*MockRepository)(nil).UpdateCounterMetric), arg0, arg1)
}

// UpdateGaugeMetric mocks base method.
func (m *MockRepository) UpdateGaugeMetric(arg0 context.Context, arg1 models.GaugeMetric) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGaugeMetric", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateGaugeMetric indicates an expected call of UpdateGaugeMetric.
func (mr *MockRepositoryMockRecorder) UpdateGaugeMetric(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGaugeMetric", reflect.TypeOf((*MockRepository)(nil).UpdateGaugeMetric), arg0, arg1)
}
