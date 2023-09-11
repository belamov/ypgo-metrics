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

func (repo *InMemoryRepository) GetAllGaugeMetrics(ctx context.Context) ([]models.GaugeMetric, error) {
	gaugeMetrics := make([]models.GaugeMetric, 0, len(repo.gaugeMetrics))
	for name, value := range repo.gaugeMetrics {
		gaugeMetrics = append(gaugeMetrics, models.GaugeMetric{
			Name:  name,
			Value: value,
		})
	}

	return gaugeMetrics, nil
}

func (repo *InMemoryRepository) GetAllCounterMetrics(ctx context.Context) ([]models.CounterMetric, error) {
	counterMetrics := make([]models.CounterMetric, 0, len(repo.counterMetrics))
	for name, value := range repo.counterMetrics {
		counterMetrics = append(counterMetrics, models.CounterMetric{
			Name:  name,
			Value: value,
		})
	}

	return counterMetrics, nil
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		gaugeMetrics:   make(map[string]float64),
		counterMetrics: make(map[string]int64),
		gaugeMutex:     sync.RWMutex{},
		counterMutex:   sync.RWMutex{},
	}
}

func (repo *InMemoryRepository) GetGaugeMetricByName(ctx context.Context, name string) (*models.GaugeMetric, error) {
	metric, ok := repo.gaugeMetrics[name]
	if !ok {
		return nil, ErrMetricNotFound
	}

	return &models.GaugeMetric{
		Name:  name,
		Value: metric,
	}, nil
}

func (repo *InMemoryRepository) GetCounterMetricByName(ctx context.Context, name string) (*models.CounterMetric, error) {
	metric, ok := repo.counterMetrics[name]
	if !ok {
		return nil, ErrMetricNotFound
	}

	return &models.CounterMetric{
		Name:  name,
		Value: metric,
	}, nil
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
