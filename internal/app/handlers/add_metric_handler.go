package handlers

import (
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) AddGaugeMetric(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	metricRawData, _ := strings.CutPrefix(r.URL.Path, "/update/gauge/")
	metricData := strings.Split(metricRawData, "/")
	if len(metricData) != 2 {
		http.Error(w, "cant identify metric and value", http.StatusNotFound)
		return
	}

	metricName := metricData[0]
	metricValue, err := strconv.ParseFloat(metricData[1], 64)
	if err != nil {
		http.Error(w, "metric value must be int", http.StatusBadRequest)
		return
	}

	err = h.metricsService.UpdateGaugeMetric(r.Context(), metricName, metricValue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) AddCounterMetric(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	metricRawData, _ := strings.CutPrefix(r.URL.Path, "/update/counter/")
	metricData := strings.Split(metricRawData, "/")
	if len(metricData) != 2 {
		http.Error(w, "cant identify metric and value", http.StatusNotFound)
		return
	}

	metricName := metricData[0]
	metricValue, err := strconv.ParseInt(metricData[1], 10, 64)
	if err != nil {
		http.Error(w, "metric value must be int", http.StatusBadRequest)
		return
	}

	err = h.metricsService.UpdateCounterMetric(r.Context(), metricName, metricValue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
