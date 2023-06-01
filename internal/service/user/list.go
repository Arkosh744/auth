package user

import (
	"context"

	"github.com/Arkosh744/auth-service-api/internal/log"
	"github.com/Arkosh744/auth-service-api/internal/model"
	"github.com/Arkosh744/auth-service-api/internal/sys"
	"github.com/Arkosh744/auth-service-api/internal/sys/codes"
)

func (s *service) List(ctx context.Context) ([]*model.UserSpec, error) {
	usersRaw, err := s.repository.List(ctx)
	if err != nil {
		log.Errorf("error list users: %v", err)

		return nil, sys.NewCommonError("error list users", codes.Internal)
	}

	users := make([]*model.UserSpec, 0, len(usersRaw))
	for _, userRaw := range usersRaw {
		var user *model.UserSpec
		user, err = userRaw.ToUserSpec()
		if err != nil {
			log.Errorf("error list users: %v", err)

			return nil, sys.NewCommonError("error list users", codes.Internal)
		}

		users = append(users, user)
	}

	return users, nil
}
