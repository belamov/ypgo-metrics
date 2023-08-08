package services

import (
	"flag"
	"time"
)

type ServerConfig struct {
	Address string
}

func BuildServerConfig() *ServerConfig {
	cfg := &ServerConfig{}

	flag.StringVar(&cfg.Address, "a", "0.0.0.0:8080", "server address")

	flag.Parse()

	return cfg
}

type ClientConfig struct {
	ServerAddress  string
	ReportInterval time.Duration
	PollInterval   time.Duration
}

func BuildClientConfig() *ClientConfig {
	cfg := &ClientConfig{}
	var pollIntervalSeconds uint64
	var reportIntervalSeconds uint64

	flag.StringVar(&cfg.ServerAddress, "a", "localhost:8080", "server address")
	flag.Uint64Var(&reportIntervalSeconds, "r", 10, "report interval in seconds")
	flag.Uint64Var(&pollIntervalSeconds, "p", 2, "poll interval in seconds")

	flag.Parse()

	cfg.ReportInterval = time.Duration(reportIntervalSeconds) * time.Second
	cfg.PollInterval = time.Duration(pollIntervalSeconds) * time.Second

	return cfg
}
