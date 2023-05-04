package user

import (
	"context"

	"github.com/Arkosh744/auth-grpc/internal/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) validateNameEmail(ctx context.Context, user *model.User) error {
	name, email, err := s.repository.ExistsNameEmail(ctx, user)
	if err != nil {
		return status.Errorf(codes.Internal, "validate data error: %v", err)
	}

	if name {
		return status.Errorf(codes.AlreadyExists, "error: %v", ErrUsernameExists)
	}
	if email {
		return status.Errorf(codes.AlreadyExists, "error: %v", ErrEmailExists)
	}

	return nil
}
