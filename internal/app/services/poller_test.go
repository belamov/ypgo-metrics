package services

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoller_Poll(t *testing.T) {
	poller := NewPoller()

	poller.Poll()
	assert.Equal(t, uint64(1), poller.pollCount)
	poller.Poll()
	assert.Equal(t, uint64(2), poller.pollCount)
}

func TestPoller_GetMetricsToReport(t *testing.T) {
	poller := NewPoller()

	poller.Poll()
	metrics := poller.GetMetricsToReport()

	expectedGaugeMetrics := []string{
		"Alloc",
		"BuckHashSys",
		"Frees",
		"GCCPUFraction",
		"GCSys",
		"HeapAlloc",
		"HeapIdle",
		"HeapInuse",
		"HeapObjects",
		"HeapReleased",
		"HeapSys",
		"LastGC",
		"Lookups",
		"MCacheInuse",
		"MCacheSys",
		"MSpanInuse",
		"MSpanSys",
		"Mallocs",
		"NextGC",
		"NumForcedGC",
		"NumGC",
		"OtherSys",
		"PauseTotalNs",
		"StackInuse",
		"StackSys",
		"Sys",
		"TotalAlloc",
		"RandomValue",
	}

	expectedCounterMetrics := []string{
		"PollCount",
	}

	foundGaugeMetrics := make([]string, 0, len(expectedGaugeMetrics))
	foundCounterMetrics := make([]string, 0)
	for _, metric := range metrics {
		if metric.Type == "gauge" {
			foundGaugeMetrics = append(foundGaugeMetrics, metric.Name)
		}
		if metric.Type == "counter" {
			foundCounterMetrics = append(foundCounterMetrics, metric.Name)
		}
	}

	sort.Strings(expectedGaugeMetrics)
	sort.Strings(foundGaugeMetrics)
	assert.Equal(t, expectedGaugeMetrics, foundGaugeMetrics)

	sort.Strings(expectedCounterMetrics)
	sort.Strings(foundCounterMetrics)
	assert.Equal(t, expectedCounterMetrics, foundCounterMetrics)
}
