package config

import (
	"github.com/kelseyhightower/envconfig"
)

var _ GRPCConfig = (*grpcConfig)(nil)

const (
	grpcEnvPrefix = "GRPC"
)

type GRPCConfig interface {
	GetPort() string
}

type grpcConfig struct {
	Port string `required:"true"`
}

func NewGRPCConfig() (*grpcConfig, error) {
	var cfg grpcConfig
	err := envconfig.Process(grpcEnvPrefix, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (c *grpcConfig) GetPort() string {
	return c.Port
}
