package services

import (
	"context"
	"errors"
	"testing"

	"github.com/belamov/ypgo-metrics/internal/app/mocks"
	"github.com/belamov/ypgo-metrics/internal/app/models"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestMetricService_UpdateGaugeMetric(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStorage := mocks.NewMockRepository(ctrl)
	service := NewMetricService(mockStorage)

	ctx := context.Background()
	metricName := "metricName"
	metricValue := 3.14
	mockStorage.EXPECT().UpdateGaugeMetric(ctx, models.GaugeMetric{
		Name:  metricName,
		Value: metricValue,
	}).Times(1).Return(nil)

	err := service.UpdateGaugeMetric(ctx, metricName, metricValue)
	assert.NoError(t, err)
}

func TestMetricService_UpdateGaugeMetricErrorFromStorage(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStorage := mocks.NewMockRepository(ctrl)
	service := NewMetricService(mockStorage)

	storageErr := errors.New("storage error")
	ctx := context.Background()
	metricName := "metricName"
	metricValue := 3.14
	mockStorage.EXPECT().UpdateGaugeMetric(ctx, models.GaugeMetric{
		Name:  metricName,
		Value: metricValue,
	}).Times(1).Return(storageErr)

	err := service.UpdateGaugeMetric(ctx, metricName, metricValue)
	assert.ErrorIs(t, err, storageErr)
}

func TestMetricService_UpdateCounterMetric(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStorage := mocks.NewMockRepository(ctrl)
	service := NewMetricService(mockStorage)

	ctx := context.Background()
	metricName := "metricName"
	metricValue := int64(3)
	mockStorage.EXPECT().UpdateCounterMetric(ctx, models.CounterMetric{
		Name:  metricName,
		Value: metricValue,
	}).Times(1).Return(nil)

	err := service.UpdateCounterMetric(ctx, metricName, metricValue)
	assert.NoError(t, err)
}

func TestMetricService_UpdateCounterMetricErrorFromStorage(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStorage := mocks.NewMockRepository(ctrl)
	service := NewMetricService(mockStorage)

	storageErr := errors.New("storage error")
	ctx := context.Background()
	metricName := "metricName"
	metricValue := int64(3)
	mockStorage.EXPECT().UpdateCounterMetric(ctx, models.CounterMetric{
		Name:  metricName,
		Value: metricValue,
	}).Times(1).Return(storageErr)

	err := service.UpdateCounterMetric(ctx, metricName, metricValue)
	assert.ErrorIs(t, err, storageErr)
}
