package config

import (
	"fmt"
	"os"
)

var _ GRPCConfig = (*grpcConfig)(nil)

const (
	grpcEnvPrefix = "GRPC_PORT"
)

type GRPCConfig interface {
	GetPort() string
}

type grpcConfig struct {
	port string
}

func NewGRPCConfig() (*grpcConfig, error) {
	port := os.Getenv(grpcEnvPrefix)
	if port == "" {
		return nil, fmt.Errorf("grpc port is not set")
	}

	return &grpcConfig{port: port}, nil
}

func (c *grpcConfig) GetPort() string {
	return c.port
}
