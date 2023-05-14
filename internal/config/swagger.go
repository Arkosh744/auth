package config

import (
	"fmt"
	"os"
)

var _ SwaggerConfig = (*swaggerConfig)(nil)

const (
	SwaggerEnvHost = "SWAG_HOST"
	SwaggerEnvPort = "SWAG_PORT"
)

type SwaggerConfig interface {
	GetHost() string
}

type swaggerConfig struct {
	port string
	host string
}

func NewSwaggerConfig() (*swaggerConfig, error) {
	host := os.Getenv(SwaggerEnvHost)
	port := os.Getenv(SwaggerEnvPort)
	if port == "" || host == "" {
		return nil, fmt.Errorf("http addr is not set")
	}

	return &swaggerConfig{host: host, port: port}, nil
}

func (c *swaggerConfig) GetHost() string {
	return fmt.Sprintf(c.host + ":" + c.port)
}
