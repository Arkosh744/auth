package config

import (
	"errors"
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type PGConfig interface {
	DSN() string
}

type pgConfig struct {
	dsn string
}

const pgEnvPrefix = "PG"

type pgEnv struct {
	Host     string `required:"true"`
	Port     string `required:"true"`
	DB       string `required:"true"`
	User     string `required:"true"`
	Password string `required:"true"`
	SSL      string `required:"true"`
}

func NewPGConfig() (*pgConfig, error) {
	var cfg pgEnv
	err := envconfig.Process(pgEnvPrefix, &cfg)
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.DB, cfg.User, cfg.Password, cfg.SSL)

	if cfg.Host == "" || cfg.Port == "" || cfg.DB == "" || cfg.User == "" || cfg.Password == "" || cfg.SSL == "" {
		return nil, errors.New("DSN is not set")
	}

	return &pgConfig{dsn: dsn}, nil
}

func (c *pgConfig) DSN() string {
	return c.dsn
}
