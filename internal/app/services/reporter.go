package services

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/belamov/ypgo-metrics/internal/app/resources"

	"github.com/rs/zerolog/log"
)

type ReporterInterface interface {
	Report([]resources.Metric)
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

func (r *HTTPReporter) Report(metrics []resources.Metric) {
	for _, metric := range metrics {
		// TODO: client throttle
		body, err := json.Marshal(metric)
		if err != nil {
			log.Error().Err(err).Msg("error marshalling metric for report")
			continue
		}
		response, err := r.client.Post(
			r.updateURL,
			"application/json",
			bytes.NewReader(body),
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
