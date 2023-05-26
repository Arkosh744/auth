package auth

import (
	"context"

	"github.com/Arkosh744/auth-service-api/internal/config"
	userRepository "github.com/Arkosh744/auth-service-api/internal/repo/user"
)

type Service interface {
	GetAccessToken(ctx context.Context, refreshToken string) (string, error)
	GetRefreshToken(ctx context.Context, username, password string) (string, error)
}

type service struct {
	authConfig     config.AuthConfig
	userRepository userRepository.Repository
}

func NewService(authConfig config.AuthConfig, userRepository userRepository.Repository) *service {
	return &service{
		authConfig:     authConfig,
		userRepository: userRepository,
	}
}
