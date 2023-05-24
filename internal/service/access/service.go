package access

import (
	"context"
	"github.com/Arkosh744/auth-service-api/internal/config"
	accessRepo "github.com/Arkosh744/auth-service-api/internal/repo/access"
)

type Service interface {
	CheckAccess(ctx context.Context, endpointAddress string) (bool, error)
}

type service struct {
	authConfig config.AuthConfig

	repo accessRepo.Repository
}

func NewService(repo accessRepo.Repository, authConfig config.AuthConfig) *service {
	return &service{
		repo:       repo,
		authConfig: authConfig,
	}
}
