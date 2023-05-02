package user

import (
	"context"

	"github.com/Arkosh744/auth-grpc/internal/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) Create(ctx context.Context, user *model.User) error {
	if err := s.validateNameEmail(ctx, user.Username, user.Email); err != nil {
		return err
	}

	if err := s.repository.Create(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *service) validateNameEmail(ctx context.Context, username, email string) error {
	exists, err := s.repository.ExistsName(ctx, username)
	if err != nil {
		return err
	}
	if !exists {
		return status.Errorf(codes.AlreadyExists, "Error: %v", ErrUsernameExists)
	}

	exists, err = s.repository.ExistsEmail(ctx, email)
	if err != nil {
		return err
	}
	if !exists {
		return status.Errorf(codes.AlreadyExists, "Error: %v", ErrEmailExists)
	}

	return nil
}
