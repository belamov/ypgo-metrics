package storage

import (
	"context"
	"sync"
	"testing"

	"github.com/belamov/ypgo-metrics/internal/app/models"
	"github.com/stretchr/testify/assert"
)

func TestInMemoryRepository_UpdateGaugeMetric(t *testing.T) {
	repo := NewInMemoryRepository()

	metric := models.GaugeMetric{
		Name:  "gauge",
		Value: 1,
	}
	err := repo.UpdateGaugeMetric(context.Background(), metric)

	assert.NoError(t, err)
	assert.Equal(t, metric.Value, repo.gaugeMetrics[metric.Name])
}

func TestInMemoryRepository_UpdateGaugeMetricParallel(t *testing.T) {
	repo := NewInMemoryRepository()

	lastValue := 10000
	metricName := "gauge"

	wg := sync.WaitGroup{}
	for i := 0; i <= lastValue; i++ {
		metric := models.GaugeMetric{
			Name:  metricName,
			Value: float64(i),
		}
		wg.Add(1)
		go func() {
			err := repo.UpdateGaugeMetric(context.Background(), metric)
			assert.NoError(t, err)
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestInMemoryRepository_UpdateCounterMetric(t *testing.T) {
	repo := NewInMemoryRepository()

	metric := models.CounterMetric{
		Name:  "gauge",
		Value: 1,
	}
	err := repo.UpdateCounterMetric(context.Background(), metric)

	assert.NoError(t, err)
	assert.Equal(t, metric.Value, repo.counterMetrics[metric.Name])
}

func TestInMemoryRepository_UpdateCounterMetricParallel(t *testing.T) {
	repo := NewInMemoryRepository()

	lastValue := 10000
	metricName := "counter"

	wg := sync.WaitGroup{}
	for i := 0; i <= lastValue; i++ {
		metric := models.CounterMetric{
			Name:  metricName,
			Value: int64(i),
		}
		wg.Add(1)
		go func() {
			err := repo.UpdateCounterMetric(context.Background(), metric)
			assert.NoError(t, err)
			wg.Done()
		}()
	}
	wg.Wait()
}
