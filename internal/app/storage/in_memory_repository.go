package storage

import (
	"context"
	"sync"

	"github.com/belamov/ypgo-metrics/internal/app/models"
)

type InMemoryRepository struct {
	gaugeMetrics   map[string]float64
	counterMetrics map[string]int64
	mutex          sync.RWMutex
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		gaugeMetrics:   make(map[string]float64),
		counterMetrics: make(map[string]int64),
		mutex:          sync.RWMutex{},
	}
}

func (repo *InMemoryRepository) UpdateGaugeMetric(ctx context.Context, metric models.GaugeMetric) error {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	repo.gaugeMetrics[metric.Name] = metric.Value

	return nil
}

func (repo *InMemoryRepository) UpdateCounterMetric(ctx context.Context, metric models.CounterMetric) error {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	repo.counterMetrics[metric.Name] += metric.Value

	return nil
}
