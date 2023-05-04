package user

import (
	"context"

	"github.com/Arkosh744/auth-grpc/internal/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) Create(ctx context.Context, user *model.User) error {
	if err := s.validateNameEmail(ctx, user); err != nil {
		return status.Errorf(codes.Internal, "validate data error: %v", err)
	}

	if err := s.repository.Create(ctx, user); err != nil {
		return err
	}

	return nil
}
