package user

import (
	"context"

	"github.com/Arkosh744/auth-grpc/internal/model"
)

func (s *service) Get(ctx context.Context, username string) (user *model.User, err error) {
	user, err = s.repository.Get(ctx, username)
	if err != nil {
		return nil, err
	}

	return user, nil
}
