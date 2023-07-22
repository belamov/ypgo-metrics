package services

import (
	"math/rand"
	"runtime"
	"strconv"

	"github.com/belamov/ypgo-metrics/internal/app/models"
)

type PollerInterface interface {
	Poll()
	GetMetricsToReport() []models.MetricForReport
}

type Poller struct {
	ms          *runtime.MemStats
	pollCount   uint64
	randomValue uint64
}

func NewPoller() *Poller {
	return &Poller{
		ms: &runtime.MemStats{},
	}
}

func (p *Poller) Poll() {
	runtime.ReadMemStats(p.ms)
	p.pollCount++
	p.randomValue = rand.Uint64()
}

func (p *Poller) GetMetricsToReport() []models.MetricForReport {
	metrics := make([]models.MetricForReport, 29)
	metrics[0] = uintGaugeMetric("Alloc", p.ms.Alloc)
	metrics[1] = uintGaugeMetric("BuckHashSys", p.ms.BuckHashSys)
	metrics[2] = uintGaugeMetric("Frees", p.ms.Frees)
	metrics[3] = uintGaugeMetric("GCSys", p.ms.GCSys)
	metrics[4] = uintGaugeMetric("HeapAlloc", p.ms.HeapAlloc)
	metrics[5] = uintGaugeMetric("HeapIdle", p.ms.HeapIdle)
	metrics[6] = uintGaugeMetric("HeapInuse", p.ms.HeapInuse)
	metrics[7] = uintGaugeMetric("HeapObjects", p.ms.HeapObjects)
	metrics[8] = uintGaugeMetric("HeapReleased", p.ms.HeapReleased)
	metrics[9] = uintGaugeMetric("HeapSys", p.ms.HeapSys)
	metrics[10] = uintGaugeMetric("LastGC", p.ms.LastGC)
	metrics[11] = uintGaugeMetric("Lookups", p.ms.Lookups)
	metrics[12] = uintGaugeMetric("MCacheInuse", p.ms.MCacheInuse)
	metrics[13] = uintGaugeMetric("MCacheSys", p.ms.MCacheSys)
	metrics[14] = uintGaugeMetric("MSpanInuse", p.ms.MSpanInuse)
	metrics[15] = uintGaugeMetric("MSpanSys", p.ms.MSpanSys)
	metrics[16] = uintGaugeMetric("Mallocs", p.ms.Mallocs)
	metrics[17] = uintGaugeMetric("NextGC", p.ms.NextGC)
	metrics[18] = uintGaugeMetric("NumForcedGC", uint64(p.ms.NumForcedGC))
	metrics[19] = uintGaugeMetric("NumGC", uint64(p.ms.NumGC))
	metrics[20] = uintGaugeMetric("OtherSys", p.ms.OtherSys)
	metrics[21] = uintGaugeMetric("PauseTotalNs", p.ms.PauseTotalNs)
	metrics[22] = uintGaugeMetric("StackInuse", p.ms.StackInuse)
	metrics[23] = uintGaugeMetric("StackSys", p.ms.StackSys)
	metrics[24] = uintGaugeMetric("Sys", p.ms.Sys)
	metrics[25] = uintGaugeMetric("TotalAlloc", p.ms.TotalAlloc)
	metrics[26] = floatGaugeMetric("GCCPUFraction", p.ms.GCCPUFraction)
	metrics[27] = uintGaugeMetric("RandomValue", p.randomValue)
	metrics[28] = counterMetric(p.pollCount)

	return metrics
}

func uintGaugeMetric(name string, value uint64) models.MetricForReport {
	return models.MetricForReport{
		Type:  "gauge",
		Name:  name,
		Value: strconv.FormatUint(value, 10),
	}
}

func floatGaugeMetric(name string, value float64) models.MetricForReport {
	return models.MetricForReport{
		Type:  "gauge",
		Name:  name,
		Value: strconv.FormatFloat(value, 'e', -1, 64),
	}
}

func counterMetric(value uint64) models.MetricForReport {
	return models.MetricForReport{
		Type:  "counter",
		Name:  "PollCount",
		Value: strconv.FormatUint(value, 10),
	}
}
