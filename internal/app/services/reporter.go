package services

import (
	"fmt"
	"net/http"

	"github.com/belamov/ypgo-metrics/internal/app/models"
)

type ReporterInterface interface {
	Report([]models.MetricForReport)
}

type HTTPReporter struct {
	client    *http.Client
	updateURL string
}

func NewHTTPReporter(client *http.Client, updateURL string) *HTTPReporter {
	return &HTTPReporter{
		client:    client,
		updateURL: updateURL,
	}
}

func (r *HTTPReporter) Report(metrics []models.MetricForReport) {
	for _, metric := range metrics {
		// TODO: client throttle
		response, err := r.client.Post(
			fmt.Sprintf("%s/%s/%s/%s", r.updateURL, metric.Type, metric.Name, metric.Value),
			"text/plain",
			nil,
		)
		if err != nil {
			fmt.Println(fmt.Errorf("update metric error: %w", err))
			_ = response.Body.Close()
			continue
		}

		if response.StatusCode != http.StatusOK {
			fmt.Println(fmt.Errorf("unexpected update response: %v", response))
		}
		_ = response.Body.Close()
	}
}
