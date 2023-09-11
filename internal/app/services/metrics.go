package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/belamov/ypgo-metrics/internal/app/models"
	"github.com/belamov/ypgo-metrics/internal/app/storage"
)

type MetricServiceInterface interface {
	UpdateCounterMetric(ctx context.Context, name string, value int64) error
	GetCounterMetric(ctx context.Context, name string) (int64, error)
	UpdateGaugeMetric(ctx context.Context, name string, value float64) error
	GetGaugeMetric(ctx context.Context, name string) (float64, error)
	GetAllMetrics(ctx context.Context) ([]models.CounterMetric, []models.GaugeMetric, error)
}

type MetricService struct {
	repo storage.Repository
}

var ErrMetricNotFound = errors.New("service error: metric not found")

func NewMetricService(storage storage.Repository) *MetricService {
	return &MetricService{
		repo: storage,
	}
}

func (service *MetricService) GetCounterMetric(ctx context.Context, name string) (int64, error) {
	metric, err := service.repo.GetCounterMetricByName(ctx, name)
	if errors.Is(err, storage.ErrMetricNotFound) {
		return 0, ErrMetricNotFound
	}
	if err != nil {
		return 0, err
	}
	return metric.Value, nil
}

func (service *MetricService) GetGaugeMetric(ctx context.Context, name string) (float64, error) {
	metric, err := service.repo.GetGaugeMetricByName(ctx, name)
	if errors.Is(err, storage.ErrMetricNotFound) {
		return 0, ErrMetricNotFound
	}
	if err != nil {
		return 0, err
	}

	return metric.Value, nil
}

func (service *MetricService) UpdateCounterMetric(ctx context.Context, name string, value int64) error {
	metric := models.CounterMetric{
		Name:  name,
		Value: value,
	}

	err := service.repo.UpdateCounterMetric(ctx, metric)
	if err != nil {
		return fmt.Errorf("cant update metric in storage: %w", err)
	}

	return nil
}

func (service *MetricService) UpdateGaugeMetric(ctx context.Context, name string, value float64) error {
	metric := models.GaugeMetric{
		Name:  name,
		Value: value,
	}
	return service.repo.UpdateGaugeMetric(ctx, metric)
}

func (service *MetricService) GetAllMetrics(ctx context.Context) ([]models.CounterMetric, []models.GaugeMetric, error) {
	counterMetrics, err := service.repo.GetAllCounterMetrics(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("error getting counter metrics from repo: %w", err)
	}

	gaugeMetrics, err := service.repo.GetAllGaugeMetrics(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("error getting gauge metrics from repo: %w", err)
	}

	return counterMetrics, gaugeMetrics, nil
}
