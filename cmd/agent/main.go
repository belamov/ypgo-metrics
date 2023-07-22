package main

import (
	"net/http"
	"time"

	"github.com/belamov/ypgo-metrics/internal/app/agent"
	"github.com/belamov/ypgo-metrics/internal/app/services"
)

func main() {
	pollInterval := 2 * time.Second
	reportInterval := 10 * time.Second

	poller := services.NewPoller()
	reporter := services.NewHTTPReporter(http.DefaultClient, "http://localhost:8080/update")

	a := agent.NewAgent(
		pollInterval,
		reportInterval,
		poller,
		reporter,
	)

	a.Run()
}
