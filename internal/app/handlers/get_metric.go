package handlers

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/belamov/ypgo-metrics/internal/app/services"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

func (h *Handler) GetMetric(w http.ResponseWriter, r *http.Request) {
	metricType := chi.URLParam(r, "metricType")

	switch metricType {
	case TypeCounter:
		h.getCounterMetric(r.Context(), w, r)
	case TypeGauge:
		h.getGaugeMetric(r.Context(), w, r)
	default:
		http.Error(w, "unknown metric type", http.StatusBadRequest)
		return
	}
}

func (h *Handler) getCounterMetric(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	metricName := chi.URLParam(r, "metricName")

	metricValue, err := h.metricsService.GetCounterMetric(ctx, metricName)
	if errors.Is(err, services.ErrMetricNotFound) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		log.Error().Err(err).Msg("error getting counter value from service")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte(strconv.FormatInt(metricValue, 10)))
	if err != nil {
		log.Error().Err(err).Msg("error writing counter value")
		http.Error(w, "cant write metric", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) getGaugeMetric(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	metricName := chi.URLParam(r, "metricName")

	metricValue, err := h.metricsService.GetGaugeMetric(ctx, metricName)
	if errors.Is(err, services.ErrMetricNotFound) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		log.Error().Err(err).Msg("error getting gauge value from service")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte(strconv.FormatFloat(metricValue, 'f', -1, 64)))
	if err != nil {
		log.Error().Err(err).Msg("error writing gauge value")
		http.Error(w, "cant write metric", http.StatusInternalServerError)
		return
	}
}
