package user

import (
	"context"
	"github.com/Arkosh744/auth-service-api/internal/logger"

	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) Delete(ctx context.Context, username string) error {
	err := s.repository.Delete(ctx, username)
	if err != nil {
		if err == pgx.ErrNoRows {
			return status.Errorf(codes.NotFound, "error: %v", ErrNotFound)
		}
		logger.Log.Error("error delete user: %v", err)

		return err
	}

	return nil
}
