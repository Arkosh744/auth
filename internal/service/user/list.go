package user

import (
	"context"
	"github.com/Arkosh744/auth-service-api/internal/logger"

	"github.com/Arkosh744/auth-service-api/internal/model"
)

func (s *service) List(ctx context.Context) ([]*model.User, error) {
	users, err := s.repository.List(ctx)
	if err != nil {
		logger.Log.Error("error list users: %v", err)

		return nil, err
	}

	return users, nil
}
