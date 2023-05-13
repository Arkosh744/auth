package user

import (
	"context"
	"github.com/Arkosh744/auth-service-api/internal/log"

	"github.com/Arkosh744/auth-service-api/internal/model"
)

func (s *service) List(ctx context.Context) ([]*model.User, error) {
	users, err := s.repository.List(ctx)
	if err != nil {
		log.Errorf("error list users: %v", err)

		return nil, err
	}

	return users, nil
}
