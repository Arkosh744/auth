package user

import (
	"context"

	"github.com/Arkosh744/auth-service-api/internal/log"
	"github.com/Arkosh744/auth-service-api/internal/model"
	"github.com/Arkosh744/auth-service-api/internal/sys"
	"github.com/Arkosh744/auth-service-api/internal/sys/codes"
	"github.com/Arkosh744/auth-service-api/internal/sys/validate"
)

func (s *service) Create(ctx context.Context, user *model.UserSpec) error {
	var userIdentifier model.UserIdentifier
	userIdentifier.Set(user.Username, user.Email)

	exists, err := s.repository.ExistsNameEmail(ctx, &userIdentifier)
	if err != nil {
		return sys.NewCommonError("error checking if user credentials exists", codes.Internal)
	}

	err = validate.Validate(
		ctx,
		checkExists(exists),
	)
	if err != nil {
		return err
	}

	userSerialized, err := user.ToUserSpecSerialized()
	if err != nil {
		log.Errorf("error create user: %v", err)

		return sys.NewCommonError("error creating user", codes.Internal)
	}

	if err = s.repository.Create(ctx, userSerialized); err != nil {
		log.Errorf("error create user: %v", err)

		return sys.NewCommonError("error creating user", codes.Internal)
	}

	return nil
}

func checkExists(s model.ExistsStatus) validate.Condition {
	return func(ctx context.Context) error {
		switch s {
		case model.StatusUsernameExists:
			return validate.NewValidationErrors(ErrUsernameExists)
		case model.StatusEmailExists:
			return validate.NewValidationErrors(ErrEmailExists)
		case model.StatusBothExist:
			return validate.NewValidationErrors(ErrBothExists)
		default:
			return nil
		}
	}
}
