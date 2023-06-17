package user

import (
	"context"

	"github.com/Arkosh744/auth-service-api/internal/log"
	"github.com/Arkosh744/auth-service-api/internal/model"
	"github.com/Arkosh744/auth-service-api/internal/sys"
	"github.com/Arkosh744/auth-service-api/internal/sys/codes"
	"github.com/jackc/pgx/v4"
)

func (s *service) Get(ctx context.Context, username string) (*model.UserSpec, error) {
	userRaw, err := s.repository.Get(ctx, username)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, sys.NewCommonError(ErrNotFound, codes.NotFound)
		}
		log.Errorf("error getting user: %v", err)

		return nil, sys.NewCommonError("error getting user", codes.Internal)
	}

	user, err := userRaw.ToUserSpec()
	if err != nil {
		log.Errorf("error getting user: %v", err)

		return nil, sys.NewCommonError("error getting user", codes.Internal)
	}

	return user, nil
}
