package config

import (
	"fmt"
	"net"
	"os"
)

var _ PromConfig = (*promConfig)(nil)

const (
	promEnvHost = "PROM_HOST"
	promEnvPort = "PROM_PORT"
)

type PromConfig interface {
	GetHost() string
}

type promConfig struct {
	port string
	host string
}

func NewPromConfig() (*promConfig, error) {
	host := os.Getenv(promEnvHost)
	port := os.Getenv(promEnvPort)
	if port == "" || host == "" {
		return nil, fmt.Errorf("prometheus addr is not set")
	}

	return &promConfig{host: host, port: port}, nil
}

func (c *promConfig) GetHost() string {
	return net.JoinHostPort(c.host, c.port)
}
