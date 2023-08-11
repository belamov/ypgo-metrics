package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/belamov/ypgo-metrics/internal/app"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/belamov/ypgo-metrics/internal/app/agent"
	"github.com/belamov/ypgo-metrics/internal/app/services"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Caller().Logger()

	config := app.BuildClientConfig()

	poller := services.NewPoller()
	httpClient := &http.Client{
		Timeout: config.ReportInterval / 2,
	}
	reporter := services.NewHTTPReporter(
		httpClient,
		fmt.Sprintf("http://%s/update", config.ServerAddress),
	)

	a := agent.NewAgent(
		config.PollInterval,
		config.ReportInterval,
		poller,
		reporter,
	)

	a.Run()
}
