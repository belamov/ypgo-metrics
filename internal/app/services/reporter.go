package services

import (
	"fmt"
	"net/http"

	"github.com/belamov/ypgo-metrics/internal/app/models"
)

type ReporterInterface interface {
	Report([]models.MetricForReport)
}

type HttpReporter struct {
	client    *http.Client
	updateUrl string
}

func NewHttpReporter(client *http.Client, updateUrl string) *HttpReporter {
	return &HttpReporter{
		client:    client,
		updateUrl: updateUrl,
	}
}

func (r *HttpReporter) Report(metrics []models.MetricForReport) {
	for _, metric := range metrics {
		// TODO: client throttle
		response, err := r.client.Post(
			fmt.Sprintf("%s/%s/%s/%s", r.updateUrl, metric.Type, metric.Name, metric.Value),
			"text/plain",
			nil,
		)
		if err != nil {
			fmt.Println(fmt.Errorf("update metric error: %w", err))
			continue
		}

		if response.StatusCode != http.StatusOK {
			fmt.Println(fmt.Errorf("unexpected update response: %v", response))
		}
	}
}
