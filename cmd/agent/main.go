package main

import (
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/belamov/ypgo-metrics/internal/app/agent"
	"github.com/belamov/ypgo-metrics/internal/app/services"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Caller().Logger()

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
