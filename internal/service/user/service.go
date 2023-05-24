package user

import (
	"context"

	"github.com/Arkosh744/auth-service-api/internal/model"
	userRepo "github.com/Arkosh744/auth-service-api/internal/repo/user"
)

var _ Service = (*service)(nil)

type Service interface {
	Create(ctx context.Context, user *model.UserSpec) error
	Get(ctx context.Context, username string) (*model.UserSpec, error)
	List(ctx context.Context) ([]*model.UserSpec, error)
	Update(ctx context.Context, username string, user *model.UpdateUser) error
	Delete(ctx context.Context, username string) error
}

type service struct {
	repository userRepo.Repository
}

func NewService(repo userRepo.Repository) *service {
	return &service{
		repository: repo,
	}
}
