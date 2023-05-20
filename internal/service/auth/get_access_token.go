package auth

import (
	"context"
	"github.com/Arkosh744/auth-service-api/internal/pkg/token"
)

func (s *service) GetAccessToken(ctx context.Context, refreshToken string) (string, error) {
	claims, err := token.VerifyToken(refreshToken, s.authConfig.RefreshTokenSecretKey())
	if err != nil {
		return "", err
	}

	userInfo, err := s.userRepository.Get(ctx, claims.Username)
	if err != nil {
		return "", err
	}

	accessToken, err := token.GenerateToken(&userInfo.User, s.authConfig.AccessTokenSecretKey(), s.authConfig.AccessTokenExpirationMinutes())
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
