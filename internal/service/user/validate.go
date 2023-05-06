package user

import (
	"context"

	"github.com/Arkosh744/auth-grpc/internal/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) validateNameEmail(ctx context.Context, user *model.UserIdentifier) error {
	exists, err := s.repository.ExistsNameEmail(ctx, user)
	if err != nil {
		return status.Errorf(codes.Internal, "validate data error: %v", err)
	}

	return checkExists(exists)
}

func checkExists(s model.ExistsStatus) error {
	switch s {
	case model.StatusUsernameExists:
		return status.Errorf(codes.AlreadyExists, "error: %v", ErrUsernameExists)
	case model.StatusEmailExists:
		return status.Errorf(codes.AlreadyExists, "error: %v", ErrEmailExists)
	case model.StatusBothExist:
		return status.Errorf(codes.AlreadyExists, "error: %v", ErrBothExists)
	default:
		return nil
	}
}
