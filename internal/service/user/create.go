package user

import (
	"context"

	"github.com/Arkosh744/auth-grpc/internal/model"
)

func (s *service) Create(ctx context.Context, info *model.User) error {
	err := s.repository.Create(ctx, info)
	if err != nil {
		return err
	}

	return nil
}
