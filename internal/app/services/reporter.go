package services

import (
	"fmt"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"

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

		// for immediate tcp connection reuse
		if response != nil {
			_, _ = io.Copy(io.Discard, response.Body)
			_ = response.Body.Close()
		}

		if err != nil {
			log.Error().Err(err).Msg("update metric error")
			continue
		}

		if response.StatusCode != http.StatusOK {
			log.Error().Any("response", response).Msg("unexpected update response")
		}
	}
}
