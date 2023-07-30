package handlers

import (
	"fmt"
	"net/http"

	"github.com/belamov/ypgo-metrics/internal/app/services"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func (s *HandlersTestSuite) TestGetCounterMetric() {
	metricName := "metricName"
	metricValue := int64(10)

	s.mockService.EXPECT().GetCounterMetric(gomock.Any(), metricName).Times(1).Return(metricValue, nil)

	result, response := s.testRequest(
		http.MethodGet,
		fmt.Sprintf("/value/counter/%s", metricName),
		"",
		nil,
	)
	_ = result.Body.Close()

	assert.Equal(s.T(), http.StatusOK, result.StatusCode)
	assert.Equal(s.T(), "10", response)
}

func (s *HandlersTestSuite) TestGetUnknownCounterMetric() {
	metricName := "metricName"

	s.mockService.EXPECT().GetCounterMetric(gomock.Any(), metricName).Times(1).Return(int64(0), services.ErrMetricNotFound)

	result, _ := s.testRequest(
		http.MethodGet,
		fmt.Sprintf("/value/counter/%s", metricName),
		"",
		nil,
	)
	_ = result.Body.Close()

	assert.Equal(s.T(), http.StatusNotFound, result.StatusCode)
}

func (s *HandlersTestSuite) TestGetGaugeMetric() {
	metricName := "metricName"
	metricValue := float64(10)

	s.mockService.EXPECT().GetGaugeMetric(gomock.Any(), metricName).Times(1).Return(metricValue, nil)

	result, response := s.testRequest(
		http.MethodGet,
		fmt.Sprintf("/value/gauge/%s", metricName),
		"",
		nil,
	)
	_ = result.Body.Close()

	assert.Equal(s.T(), http.StatusOK, result.StatusCode)
	assert.Equal(s.T(), "10", response)
}

func (s *HandlersTestSuite) TestGetUnknownGaugeMetric() {
	metricName := "metricName"

	s.mockService.EXPECT().GetGaugeMetric(gomock.Any(), metricName).Times(1).Return(float64(0), services.ErrMetricNotFound)

	result, _ := s.testRequest(
		http.MethodGet,
		fmt.Sprintf("/value/gauge/%s", metricName),
		"",
		nil,
	)
	_ = result.Body.Close()

	assert.Equal(s.T(), http.StatusNotFound, result.StatusCode)
}

func (s *HandlersTestSuite) TestGetUnknownTypeMetric() {
	metricName := "metricName"

	s.mockService.EXPECT().GetGaugeMetric(gomock.Any(), metricName).Times(1).Return(float64(0), services.ErrMetricNotFound)

	result, _ := s.testRequest(
		http.MethodGet,
		fmt.Sprintf("/value/unkown/%s", metricName),
		"",
		nil,
	)
	_ = result.Body.Close()

	assert.Equal(s.T(), http.StatusBadRequest, result.StatusCode)
}
