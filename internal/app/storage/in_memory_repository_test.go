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

func TestInMemoryRepository_GetCounterMetric(t *testing.T) {
	repo := NewInMemoryRepository()

	metricName := "metricName"
	metricValue := int64(1)
	repo.counterMetrics[metricName] = metricValue
	metric, err := repo.GetCounterMetricByName(context.Background(), metricName)

	assert.NoError(t, err)
	assert.Equal(t, metricValue, metric.Value)
	assert.Equal(t, metricName, metric.Name)
}

func TestInMemoryRepository_GetGaugeMetric(t *testing.T) {
	repo := NewInMemoryRepository()

	metricName := "metricName"
	metricValue := float64(1)
	repo.gaugeMetrics[metricName] = metricValue
	metric, err := repo.GetGaugeMetricByName(context.Background(), metricName)

	assert.NoError(t, err)
	assert.Equal(t, metricValue, metric.Value)
	assert.Equal(t, metricName, metric.Name)
}

func TestInMemoryRepository_GetAllGaugeMetrics(t *testing.T) {
	repo := NewInMemoryRepository()

	metricName := "metricName"
	metricValue := float64(1)
	repo.gaugeMetrics[metricName] = metricValue
	metrics, err := repo.GetAllGaugeMetrics(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, 1, len(metrics))
	assert.Equal(t, metricValue, metrics[0].Value)
	assert.Equal(t, metricName, metrics[0].Name)
}

func TestInMemoryRepository_GetUnexistingCounterMetric(t *testing.T) {
	repo := NewInMemoryRepository()

	metricName := "metricName"
	metric, err := repo.GetCounterMetricByName(context.Background(), metricName)

	assert.ErrorIs(t, err, ErrMetricNotFound)
	assert.Nil(t, metric)
}

func TestInMemoryRepository_GetUnexistingGaugeMetric(t *testing.T) {
	repo := NewInMemoryRepository()

	metricName := "metricName"
	metric, err := repo.GetGaugeMetricByName(context.Background(), metricName)

	assert.ErrorIs(t, err, ErrMetricNotFound)
	assert.Nil(t, metric)
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
