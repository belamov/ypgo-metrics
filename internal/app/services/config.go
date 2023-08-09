package services

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

type ServerConfig struct {
	Address string
}

func BuildServerConfig() *ServerConfig {
	var addressFlag string

	flag.StringVar(&addressFlag, "a", "0.0.0.0:8080", "server address")

	flag.Parse()

	cfg := &ServerConfig{
		Address: coalesceStrings(os.Getenv("ADDRESS"), addressFlag),
	}

	return cfg
}

type ClientConfig struct {
	ServerAddress  string
	ReportInterval time.Duration
	PollInterval   time.Duration
}

func BuildClientConfig() *ClientConfig {
	var serverAddressFlag string
	var pollIntervalFlag uint64
	var reportIntervalFlag uint64

	flag.StringVar(&serverAddressFlag, "a", "localhost:8080", "server address")
	flag.Uint64Var(&reportIntervalFlag, "r", 10, "report interval in seconds")
	flag.Uint64Var(&pollIntervalFlag, "p", 2, "poll interval in seconds")

	flag.Parse()

	reportIntervalSeconds := coalesceUints(uintFromEnv("REPORT_INTERVAL"), reportIntervalFlag)
	pollIntervalSeconds := coalesceUints(uintFromEnv("POLL_INTERVAL"), pollIntervalFlag)
	cfg := &ClientConfig{
		ServerAddress:  coalesceStrings(os.Getenv("ADDRESS"), serverAddressFlag),
		ReportInterval: time.Duration(reportIntervalSeconds) * time.Second,
		PollInterval:   time.Duration(pollIntervalSeconds) * time.Second,
	}

	return cfg
}

func uintFromEnv(key string) uint64 {
	s := os.Getenv(key)
	v, err := strconv.ParseUint(s, 10, 64)
	if err != nil && s != "" {
		panic(fmt.Sprintf("invalid env value: %s. %s", s, err))
	}
	return v
}

func coalesceStrings(strings ...string) string {
	for _, str := range strings {
		if str != "" {
			return str
		}
	}
	return ""
}

func coalesceUints(uints ...uint64) uint64 {
	for _, v := range uints {
		if v != 0 {
			return v
		}
	}
	return 0
}
