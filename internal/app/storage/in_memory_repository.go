package storage

import (
	"context"
	"sync"

	"github.com/belamov/ypgo-metrics/internal/app/models"
)

type InMemoryRepository struct {
	gaugeMetrics   map[string]float64
	counterMetrics map[string]int64
	gaugeMutex     sync.RWMutex
	counterMutex   sync.RWMutex
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		gaugeMetrics:   make(map[string]float64),
		counterMetrics: make(map[string]int64),
		gaugeMutex:     sync.RWMutex{},
		counterMutex:   sync.RWMutex{},
	}
}

func (repo *InMemoryRepository) UpdateGaugeMetric(ctx context.Context, metric models.GaugeMetric) error {
	repo.gaugeMutex.Lock()
	defer repo.gaugeMutex.Unlock()

	repo.gaugeMetrics[metric.Name] = metric.Value

	return nil
}

func (repo *InMemoryRepository) UpdateCounterMetric(ctx context.Context, metric models.CounterMetric) error {
	repo.counterMutex.Lock()
	defer repo.counterMutex.Unlock()

	repo.counterMetrics[metric.Name] += metric.Value

	return nil
}
