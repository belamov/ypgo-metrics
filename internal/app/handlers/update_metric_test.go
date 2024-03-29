package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func (s *HandlersTestSuite) TestCounterMetricAdd() {
	metricName := "metricName"
	metricValue := int64(10)

	s.mockService.EXPECT().UpdateCounterMetric(gomock.Any(), metricName, metricValue).Times(1)

	result, _ := s.testRequest(
		http.MethodPost,
		fmt.Sprintf("/update/counter/%s/%d", metricName, metricValue),
		"",
		nil,
	)
	_ = result.Body.Close()

	assert.Equal(s.T(), http.StatusOK, result.StatusCode)
}

func (s *HandlersTestSuite) TestCounterMetricAddServiceError() {
	metricName := "metricName"
	metricValue := int64(10)

	s.mockService.EXPECT().UpdateCounterMetric(gomock.Any(), metricName, metricValue).
		Times(1).
		Return(errors.New("unexpected error"))

	result, _ := s.testRequest(
		http.MethodPost,
		fmt.Sprintf("/update/counter/%s/%d", metricName, metricValue),
		"",
		nil,
	)
	_ = result.Body.Close()

	assert.Equal(s.T(), http.StatusInternalServerError, result.StatusCode)
}

func (s *HandlersTestSuite) TestCounterMetricAddNotFloat() {
	metricName := "metricName"
	metricValue := "not float"

	result, _ := s.testRequest(
		http.MethodPost,
		fmt.Sprintf("/update/counter/%s/%s", metricName, metricValue),
		"",
		nil,
	)
	_ = result.Body.Close()

	assert.Equal(s.T(), http.StatusBadRequest, result.StatusCode)
}

func (s *HandlersTestSuite) TestCounterMetricWrongFormat() {
	metricName := "metricName"

	result, _ := s.testRequest(
		http.MethodPost,
		fmt.Sprintf("/update/counter/%s", metricName),
		"",
		nil,
	)
	_ = result.Body.Close()

	assert.Equal(s.T(), http.StatusNotFound, result.StatusCode)
}

func (s *HandlersTestSuite) TestCounterMetricAddWrongMethod() {
	metricName := "metricName"
	metricValue := 10.0

	result, _ := s.testRequest(
		http.MethodGet,
		fmt.Sprintf("/update/counter/%s/%f", metricName, metricValue),
		"",
		nil,
	)
	_ = result.Body.Close()
	assert.Equal(s.T(), http.StatusMethodNotAllowed, result.StatusCode)

	result, _ = s.testRequest(
		http.MethodPut,
		fmt.Sprintf("/update/counter/%s/%f", metricName, metricValue),
		"",
		nil,
	)
	_ = result.Body.Close()
	assert.Equal(s.T(), http.StatusMethodNotAllowed, result.StatusCode)
}

func (s *HandlersTestSuite) TestGaugeMetricAdd() {
	metricName := "metricName"
	metricValue := 10.5

	s.mockService.EXPECT().UpdateGaugeMetric(gomock.Any(), metricName, metricValue).Times(1)

	result, _ := s.testRequest(
		http.MethodPost,
		fmt.Sprintf("/update/gauge/%s/%f", metricName, metricValue),
		"",
		nil,
	)
	_ = result.Body.Close()

	assert.Equal(s.T(), http.StatusOK, result.StatusCode)
}

func (s *HandlersTestSuite) TestGaugeMetricAddServiceError() {
	metricName := "metricName"
	metricValue := float64(10)

	s.mockService.EXPECT().UpdateGaugeMetric(gomock.Any(), metricName, metricValue).
		Times(1).
		Return(errors.New("unexpected error"))

	result, _ := s.testRequest(
		http.MethodPost,
		fmt.Sprintf("/update/gauge/%s/%f", metricName, metricValue),
		"",
		nil,
	)
	_ = result.Body.Close()

	assert.Equal(s.T(), http.StatusInternalServerError, result.StatusCode)
}

func (s *HandlersTestSuite) TestGaugeMetricAddNotFloat() {
	metricName := "metricName"
	metricValue := "not float"

	result, _ := s.testRequest(
		http.MethodPost,
		fmt.Sprintf("/update/gauge/%s/%s", metricName, metricValue),
		"",
		nil,
	)
	_ = result.Body.Close()

	assert.Equal(s.T(), http.StatusBadRequest, result.StatusCode)
}

func (s *HandlersTestSuite) TestGaugeMetricWrongFormat() {
	metricName := "metricName"

	result, _ := s.testRequest(
		http.MethodPost,
		fmt.Sprintf("/update/gauge/%s", metricName),
		"",
		nil,
	)
	_ = result.Body.Close()

	assert.Equal(s.T(), http.StatusNotFound, result.StatusCode)
}

func (s *HandlersTestSuite) TestGaugeMetricAddWrongMethod() {
	metricName := "metricName"
	metricValue := 10.0

	result, _ := s.testRequest(
		http.MethodGet,
		fmt.Sprintf("/update/gauge/%s/%f", metricName, metricValue),
		"",
		nil,
	)
	_ = result.Body.Close()
	assert.Equal(s.T(), http.StatusMethodNotAllowed, result.StatusCode)

	result, _ = s.testRequest(
		http.MethodPut,
		fmt.Sprintf("/update/gauge/%s/%f", metricName, metricValue),
		"",
		nil,
	)
	_ = result.Body.Close()
	assert.Equal(s.T(), http.StatusMethodNotAllowed, result.StatusCode)
}

func (s *HandlersTestSuite) TestUnknownMetricType() {
	metricName := "metricName"
	metricValue := int64(10)

	result, _ := s.testRequest(
		http.MethodPost,
		fmt.Sprintf("/update/unknown/%s/%d", metricName, metricValue),
		"",
		nil,
	)
	_ = result.Body.Close()

	assert.Equal(s.T(), http.StatusBadRequest, result.StatusCode)
}
