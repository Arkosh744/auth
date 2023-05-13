package user

import (
	"context"
	"github.com/Arkosh744/auth-service-api/internal/log"

	"github.com/Arkosh744/auth-service-api/internal/model"
	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) Update(ctx context.Context, username string, user *model.UpdateUser) error {
	if err := s.validateNameEmail(ctx, &user.UserIdentifier); err != nil {
		return err
	}

	if err := s.repository.Update(ctx, username, user); err != nil {
		if err == pgx.ErrNoRows {
			return status.Errorf(codes.NotFound, "error: %v", ErrNotFound)
		}
		log.Errorf("error update user: %v", err)

		return err
	}

	return nil
}
