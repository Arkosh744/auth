package user

import (
	"context"

	"github.com/Arkosh744/auth-grpc/internal/client/pg"
	"github.com/Arkosh744/auth-grpc/internal/model"
	userRepo "github.com/Arkosh744/auth-grpc/internal/repo/user"
	"go.uber.org/zap"
)

var _ Service = (*service)(nil)

type Service interface {
	Create(ctx context.Context, user *model.User) error
	Get(ctx context.Context, username string) (*model.User, error)
	List(ctx context.Context) ([]*model.User, *pg.Records, error)
	Update(ctx context.Context, username string, user *model.UpdateUser) error
	Delete(ctx context.Context, username string) error
}

type service struct {
	log        *zap.SugaredLogger
	repository userRepo.Repository
}

func NewService(repo userRepo.Repository, log *zap.SugaredLogger) *service {
	return &service{
		repository: repo,
		log:        log,
	}
}
