package agent

import (
	"os"
	"os/signal"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/belamov/ypgo-metrics/internal/app/services"
)

type Agent struct {
	reportTicker *time.Ticker
	pollTicker   *time.Ticker
	poller       services.PollerInterface
	reporter     services.ReporterInterface
}

func NewAgent(
	pollInterval time.Duration,
	reportInterval time.Duration,
	poller services.PollerInterface,
	reporter services.ReporterInterface,
) *Agent {
	return &Agent{
		pollTicker:   time.NewTicker(pollInterval),
		reportTicker: time.NewTicker(reportInterval),
		poller:       poller,
		reporter:     reporter,
	}
}

func (a *Agent) Run() {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)

	go func() {
		log.Info().Msg("Agent started")
		for {
			select {
			case <-sigint:
				log.Info().Msg("Stopping Agent")
				a.shutdown()
				return
			case <-a.pollTicker.C:
				go a.poller.Poll()
			case <-a.reportTicker.C:
				go a.reporter.Report(a.poller.GetMetricsToReport())
			}
		}
	}()

	<-sigint
}

func (a *Agent) shutdown() {
	a.pollTicker.Stop()
	a.reportTicker.Stop()
	log.Info().Msg("Agent Stopped")
}
