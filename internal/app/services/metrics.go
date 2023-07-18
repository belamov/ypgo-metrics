package services

import (
	"context"

	"github.com/belamov/ypgo-metrics/internal/app/models"
	"github.com/belamov/ypgo-metrics/internal/app/storage"
)

type MetricServiceInterface interface {
	UpdateCounterMetric(ctx context.Context, name string, value int64) error
	UpdateGaugeMetric(ctx context.Context, name string, value float64) error
}

type MetricService struct {
	repo storage.Repository
}

func NewMetricService(storage storage.Repository) *MetricService {
	return &MetricService{
		repo: storage,
	}
}

func (service *MetricService) UpdateCounterMetric(ctx context.Context, name string, value int64) error {
	metric := models.CounterMetric{
		Name:  name,
		Value: value,
	}
	return service.repo.UpdateCounterMetric(ctx, metric)
}

func (service *MetricService) UpdateGaugeMetric(ctx context.Context, name string, value float64) error {
	metric := models.GaugeMetric{
		Name:  name,
		Value: value,
	}
	return service.repo.UpdateGaugeMetric(ctx, metric)
}
