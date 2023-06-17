package user

import (
	"context"

	"github.com/Arkosh744/auth-service-api/internal/log"
	"github.com/Arkosh744/auth-service-api/internal/model"
	"github.com/Arkosh744/auth-service-api/internal/sys"
	"github.com/Arkosh744/auth-service-api/internal/sys/codes"
	"github.com/Arkosh744/auth-service-api/internal/sys/validate"
	"github.com/jackc/pgx/v4"
)

func (s *service) Update(ctx context.Context, username string, user *model.UpdateUser) error {
	exists, err := s.repository.ExistsNameEmail(ctx, &user.UserIdentifier)
	if err != nil {
		return sys.NewCommonError("error checking if user credentials exists", codes.Internal)
	}

	if err = validate.Validate(
		ctx,
		checkExists(exists),
	); err != nil {
		log.Errorf("error update user: %v", err)

		return sys.NewCommonError(err.Error(), codes.Internal)
	}

	if err = s.repository.Update(ctx, username, user); err != nil {
		if err == pgx.ErrNoRows {
			return sys.NewCommonError(ErrNotFound, codes.NotFound)
		}
		log.Errorf("error update user: %v", err)

		return sys.NewCommonError("error update user", codes.Internal)
	}

	return nil
}
