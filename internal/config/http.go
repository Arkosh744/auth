package config

import (
	"fmt"
	"os"
)

var _ HTTPConfig = (*httpConfig)(nil)

const (
	httpEnvHost = "HTTP_HOST"
	httpEnvPort = "HTTP_PORT"
)

type HTTPConfig interface {
	GetHost() string
}

type httpConfig struct {
	port string
	host string
}

func NewHTTPConfig() (*httpConfig, error) {
	host := os.Getenv(httpEnvHost)
	port := os.Getenv(httpEnvPort)
	if port == "" || host == "" {
		return nil, fmt.Errorf("http addr is not set")
	}

	return &httpConfig{port: port}, nil
}

func (c *httpConfig) GetHost() string {
	return fmt.Sprintf(c.host + ":" + c.port)
}
