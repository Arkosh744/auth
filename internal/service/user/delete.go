package user

import (
	"context"

	"github.com/Arkosh744/auth-service-api/internal/log"
	"github.com/Arkosh744/auth-service-api/internal/sys"
	"github.com/Arkosh744/auth-service-api/internal/sys/codes"

	"github.com/jackc/pgx/v4"
)

func (s *service) Delete(ctx context.Context, username string) error {
	err := s.repository.Delete(ctx, username)
	if err != nil {
		if err == pgx.ErrNoRows {
			return sys.NewCommonError(ErrNotFound, codes.NotFound)
		}
		log.Errorf("error delete user: %v", err)

		return sys.NewCommonError("error deleting user", codes.Internal)
	}

	return nil
}
