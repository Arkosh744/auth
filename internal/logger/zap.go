package logger

import (
	"context"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
	"log"
)

var (
	Log *zap.SugaredLogger

	_ LogConfig = (*logConfig)(nil)
)

const (
	loggerEnvPrefix = "LOG"
)

type LogConfig interface {
	GetPreset() string
}

type logConfig struct {
	Preset string `default:"dev"`
}

func InitLogger(_ context.Context) error {
	zapLog, err := SelectLogger()
	if err != nil {
		log.Fatalf("failed to get logger: %s", err.Error())

		return err
	}

	Log = zapLog.Sugar()

	return nil
}

func SelectLogger() (*zap.Logger, error) {
	switch getLogConfig().GetPreset() {
	case "prod":
		return zap.NewProduction()
	case "dev":
		return zap.NewDevelopment()
	default:
		log.Println("unknown logger preset, using development preset")
		return zap.NewDevelopment()
	}
}

func NewLogConfig() (*logConfig, error) {
	var cfg logConfig
	err := envconfig.Process(loggerEnvPrefix, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (c *logConfig) GetPreset() string {
	return c.Preset
}

func getLogConfig() LogConfig {
	cfg, err := NewLogConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %s", err.Error())
	}

	return cfg
}
