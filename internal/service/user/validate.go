package user

import (
	"context"

	"github.com/Arkosh744/auth-grpc/internal/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) validateEmail(ctx context.Context, email string) error {
	exists, err := s.repository.ExistsEmail(ctx, email)
	if err != nil {
		return err
	}

	if exists {
		return status.Errorf(codes.AlreadyExists, "Error: %v", ErrEmailExists)
	}

	return nil
}

func (s *service) validateName(ctx context.Context, username string) error {
	exists, err := s.repository.ExistsName(ctx, username)
	if err != nil {
		return err
	}

	if exists {
		return status.Errorf(codes.AlreadyExists, "Error: %v", ErrUsernameExists)
	}

	return nil
}

func (s *service) validateNameEmail(ctx context.Context, user *model.User) error {
	if err := s.validateName(ctx, user.Username); err != nil {
		return err
	}

	if err := s.validateEmail(ctx, user.Email); err != nil {
		return err
	}

	return nil
}
