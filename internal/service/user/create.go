package user

import (
	"context"

	"github.com/Arkosh744/auth-grpc/internal/model"
)

func (s *service) Create(ctx context.Context, user *model.User) error {
	if err := s.validateNameEmail(ctx, user); err != nil {
		return err
	}

	if err := s.repository.Create(ctx, user); err != nil {
		return err
	}

	return nil
}
