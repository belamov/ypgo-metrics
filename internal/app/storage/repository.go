package storage

import (
	"context"
	"errors"

	"github.com/belamov/ypgo-metrics/internal/app/models"
)

type Repository interface {
	UpdateGaugeMetric(ctx context.Context, metric models.GaugeMetric) error
	UpdateCounterMetric(ctx context.Context, metric models.CounterMetric) error
	GetGaugeMetricByName(ctx context.Context, name string) (*models.GaugeMetric, error)
	GetCounterMetricByName(ctx context.Context, name string) (*models.CounterMetric, error)
	GetAllGaugeMetrics(ctx context.Context) ([]models.GaugeMetric, error)
	GetAllCounterMetrics(ctx context.Context) ([]models.CounterMetric, error)
}

var ErrMetricNotFound = errors.New("repository error: metric not found")
