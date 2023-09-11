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

// GetAllCounterMetrics mocks base method.
func (m *MockRepository) GetAllCounterMetrics(arg0 context.Context) ([]models.CounterMetric, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCounterMetrics", arg0)
	ret0, _ := ret[0].([]models.CounterMetric)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCounterMetrics indicates an expected call of GetAllCounterMetrics.
func (mr *MockRepositoryMockRecorder) GetAllCounterMetrics(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCounterMetrics", reflect.TypeOf((*MockRepository)(nil).GetAllCounterMetrics), arg0)
}

// GetAllGaugeMetrics mocks base method.
func (m *MockRepository) GetAllGaugeMetrics(arg0 context.Context) ([]models.GaugeMetric, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllGaugeMetrics", arg0)
	ret0, _ := ret[0].([]models.GaugeMetric)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllGaugeMetrics indicates an expected call of GetAllGaugeMetrics.
func (mr *MockRepositoryMockRecorder) GetAllGaugeMetrics(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllGaugeMetrics", reflect.TypeOf((*MockRepository)(nil).GetAllGaugeMetrics), arg0)
}

// GetCounterMetricByName mocks base method.
func (m *MockRepository) GetCounterMetricByName(arg0 context.Context, arg1 string) (*models.CounterMetric, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCounterMetricByName", arg0, arg1)
	ret0, _ := ret[0].(*models.CounterMetric)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCounterMetricByName indicates an expected call of GetCounterMetricByName.
func (mr *MockRepositoryMockRecorder) GetCounterMetricByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCounterMetricByName", reflect.TypeOf((*MockRepository)(nil).GetCounterMetricByName), arg0, arg1)
}

// GetGaugeMetricByName mocks base method.
func (m *MockRepository) GetGaugeMetricByName(arg0 context.Context, arg1 string) (*models.GaugeMetric, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGaugeMetricByName", arg0, arg1)
	ret0, _ := ret[0].(*models.GaugeMetric)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGaugeMetricByName indicates an expected call of GetGaugeMetricByName.
func (mr *MockRepositoryMockRecorder) GetGaugeMetricByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGaugeMetricByName", reflect.TypeOf((*MockRepository)(nil).GetGaugeMetricByName), arg0, arg1)
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
