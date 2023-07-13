package storage

import (
	"context"

	"github.com/belamov/ypgo-metrics/internal/app/models"
)

type Repository interface {
	UpdateGaugeMetric(ctx context.Context, metric models.GaugeMetric) error
	UpdateCounterMetric(ctx context.Context, metric models.CounterMetric) error
}
