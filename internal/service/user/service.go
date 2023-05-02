package user

import (
	"context"

	"github.com/Arkosh744/auth-grpc/internal/model"
	userRepo "github.com/Arkosh744/auth-grpc/internal/repo/user"
)

var _ Service = (*service)(nil)

type Service interface {
	Create(ctx context.Context, info *model.User) error
	Get(ctx context.Context, username string) (user *model.User, err error)
}

type service struct {
	repository userRepo.Repository
}

func NewService(repo userRepo.Repository) *service {
	return &service{
		repository: repo,
	}
}
