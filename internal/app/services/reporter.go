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
	compressor Compressor
	client     *http.Client
	updateURL  string
}

func NewHTTPReporter(client *http.Client, updateURL string, compressor Compressor) *HTTPReporter {
	return &HTTPReporter{
		client:     client,
		updateURL:  updateURL,
		compressor: compressor,
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

		compressedData, err := r.compressor.GetCompressedReader(body)
		if err != nil {
			log.Error().Err(err).Msg("error initializing compressed reader")
			continue
		}

		req, err := http.NewRequest(http.MethodPost, r.updateURL, compressedData)
		if err != nil {
			log.Error().Err(err).Msg("error initializing request for report")
			continue
		}

		req.Header.Set("Content-Type", "application/json")
		r.compressor.SetHeader(req)
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
