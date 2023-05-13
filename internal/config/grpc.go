package config

import (
	"fmt"
	"os"
)

var _ GRPCConfig = (*grpcConfig)(nil)

const (
	grpcEnvHost = "GRPC_HOST"
	grpcEnvPort = "GRPC_PORT"
)

type GRPCConfig interface {
	GetHost() string
}

type grpcConfig struct {
	host string
	port string
}

func NewGRPCConfig() (*grpcConfig, error) {
	host := os.Getenv(grpcEnvHost)
	port := os.Getenv(grpcEnvPort)
	if port == "" || host == "" {
		return nil, fmt.Errorf("grpc addr is not set")
	}

	return &grpcConfig{port: port}, nil
}

func (c *grpcConfig) GetHost() string {
	return fmt.Sprintf(c.host + ":" + c.port)
}
