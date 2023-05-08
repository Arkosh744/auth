package user

import (
	"context"

	"github.com/Arkosh744/auth-grpc/internal/client/pg"
	"github.com/Arkosh744/auth-grpc/internal/model"
)

func (s *service) List(ctx context.Context) ([]*model.User, *pg.Records, error) {
	users, records, err := s.repository.List(ctx)
	if err != nil {
		return nil, nil, err
	}

	return users, records, nil
}
