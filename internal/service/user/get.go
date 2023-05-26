package user

import (
	"context"

	"github.com/Arkosh744/auth-service-api/internal/log"
	"github.com/Arkosh744/auth-service-api/internal/model"
	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) Get(ctx context.Context, username string) (*model.UserSpec, error) {
	userRaw, err := s.repository.Get(ctx, username)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "error: %v", ErrNotFound)
		}
		log.Errorf("error get user: %v", err)

		return nil, err
	}

	user, err := userRaw.ToUserSpec()
	if err != nil {
		log.Errorf("error get user: %v", err)

		return nil, err
	}

	return user, nil
}
