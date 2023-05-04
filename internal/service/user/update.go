package user

import (
	"context"

	"github.com/Arkosh744/auth-grpc/internal/model"
)

func (s *service) Update(ctx context.Context, username string, user *model.User) error {
	if err := s.validateNameEmail(ctx, user); err != nil {
		return err
	}

	if err := s.repository.Update(ctx, username, user); err != nil {
		return err
	}

	return nil
}
