package user

import (
	"context"
	"github.com/Arkosh744/auth-service-api/internal/log"

	"github.com/Arkosh744/auth-service-api/internal/model"
)

func (s *service) Create(ctx context.Context, user *model.User) error {
	var userIdentifier model.UserIdentifier
	userIdentifier.Set(user.Username, user.Email)

	if err := s.validateNameEmail(ctx, &userIdentifier); err != nil {
		log.Errorf("error create user: %v", err)

		return err
	}

	if err := s.repository.Create(ctx, user); err != nil {
		log.Errorf("error create user: %v", err)

		return err
	}

	return nil
}
