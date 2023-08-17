package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/belamov/ypgo-metrics/internal/app/resources"
	"github.com/belamov/ypgo-metrics/internal/app/services"
	"github.com/rs/zerolog/log"
)

func (h *Handler) GetMetricJSON(w http.ResponseWriter, r *http.Request) {
	var req resources.Metric
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&req); err != nil {
		log.Err(err).Msg("cannot decode request JSON body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var getErr error
	var counterMetric *int64
	var gaugeMetric *float64

	switch req.MType {
	case TypeCounter:
		counterMetric = new(int64)
		*counterMetric, getErr = h.metricsService.GetCounterMetric(r.Context(), req.ID)
	case TypeGauge:
		gaugeMetric = new(float64)
		*gaugeMetric, getErr = h.metricsService.GetGaugeMetric(r.Context(), req.ID)
	default:
		http.Error(w, "unknown metric type", http.StatusBadRequest)
		return
	}

	if errors.Is(getErr, services.ErrMetricNotFound) {
		http.Error(w, getErr.Error(), http.StatusNotFound)
		return
	}
	if getErr != nil {
		log.Error().Err(getErr).Msg("error getting metric from service")
		http.Error(w, getErr.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	resp := resources.Metric{
		ID:    req.ID,
		MType: req.MType,
		Delta: counterMetric,
		Value: gaugeMetric,
	}
	enc := json.NewEncoder(w)
	if err := enc.Encode(resp); err != nil {
		log.Err(err).Msg("error encoding response")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
