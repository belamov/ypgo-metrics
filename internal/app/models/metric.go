package models

type GaugeMetric struct {
	Name  string
	Value float64
}

type CounterMetric struct {
	Name  string
	Value int64
}

type MetricForReport struct {
	Type  string
	Name  string
	Value string
}
