package auth

import (
	"context"

	"github.com/Arkosh744/auth-service-api/internal/pkg/encrypt"
	"github.com/Arkosh744/auth-service-api/internal/pkg/token"
	"github.com/pkg/errors"
)

func (s *service) GetRefreshToken(ctx context.Context, username, password string) (string, error) {
	// TODO: add cache?
	userInfo, err := s.userRepository.Get(ctx, username)
	if err != nil {
		return "", err
	}

	if !encrypt.VerifyPassword(userInfo.Password, password) {
		return "", errors.New("invalid password")
	}

	refresh, err := token.GenerateToken(&userInfo.User, s.authConfig.RefreshTokenSecretKey(), s.authConfig.RefreshTokenExpirationMinutes())
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return refresh, nil
}
