package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var _ RateLimitConfig = (*rateLimitConfig)(nil)

const (
	rtPeriod = "RATE_LIMIT_PERIOD"
	rtLimit  = "RATE_LIMIT_LIMIT"
)

type RateLimitConfig interface {
	Period() time.Duration
	Limit() int
}

type rateLimitConfig struct {
	period time.Duration
	limit  int
}

func NewRateLimitConfig() (*rateLimitConfig, error) {
	period, err := time.ParseDuration(os.Getenv(rtPeriod))
	if err != nil {
		return nil, fmt.Errorf("invalid period: %s", err)
	}

	limit, err := strconv.Atoi(os.Getenv(rtLimit))
	if err != nil {
		return nil, fmt.Errorf("invalid limit: %s", err)
	}

	return &rateLimitConfig{period: period, limit: limit}, nil
}

func (c *rateLimitConfig) Period() time.Duration {
	return c.period
}

func (c *rateLimitConfig) Limit() int {
	return c.limit
}
