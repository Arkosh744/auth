package user

import (
	"context"
)

func (s *service) Delete(ctx context.Context, username string) error {
	err := s.repository.Delete(ctx, username)
	if err != nil {
		return err
	}

	return nil
}
