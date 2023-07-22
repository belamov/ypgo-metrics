package services

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/belamov/ypgo-metrics/internal/app/models"
	"github.com/stretchr/testify/assert"
)

type recordingTransport struct {
	req *http.Request
}

func (t *recordingTransport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	t.req = req
	return &http.Response{StatusCode: http.StatusOK}, nil
}

func TestHttpReporter_Report(t *testing.T) {
	tr := recordingTransport{}

	client := &http.Client{Transport: &tr}
	updateUrl := "update_url"

	reporter := NewHttpReporter(client, updateUrl)

	metrics := []models.MetricForReport{{
		Type:  "type",
		Name:  "name",
		Value: "value",
	}}

	reporter.Report(metrics)

	assert.Equal(t,
		fmt.Sprintf("%s/%s/%s/%s", updateUrl, metrics[0].Type, metrics[0].Name, metrics[0].Value),
		tr.req.URL.String(),
	)
}
