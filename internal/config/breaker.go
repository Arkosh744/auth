package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var _ BreakerConfig = (*breakerConfig)(nil)

const (
	brRequests = "BREAKER_REQUESTS"
	brInterval = "BREAKER_INTERVAL"
	brTimeout  = "BREAKER_TIMEOUT"
)

type BreakerConfig interface {
	Requests() int
	Interval() time.Duration
	Timeout() time.Duration
}

type breakerConfig struct {
	requests int
	interval time.Duration
	timeout  time.Duration
}

func NewBreakerConfig() (*breakerConfig, error) {
	requests, err := strconv.Atoi(os.Getenv(brRequests))
	if err != nil {
		return nil, fmt.Errorf("invalid requests number: %s", err)
	}

	interval, err := time.ParseDuration(os.Getenv(brInterval))
	if err != nil {
		return nil, fmt.Errorf("invalid interval: %s", err)
	}

	timeout, err := time.ParseDuration(os.Getenv(brTimeout))
	if err != nil {
		return nil, fmt.Errorf("invalid timeout: %s", err)
	}

	return &breakerConfig{requests: requests, interval: interval, timeout: timeout}, nil
}

func (c *breakerConfig) Requests() int {
	return c.requests
}

func (c *breakerConfig) Interval() time.Duration {
	return c.interval
}

func (c *breakerConfig) Timeout() time.Duration {
	return c.timeout
}
