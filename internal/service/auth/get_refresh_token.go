package auth

import (
	"context"

	"github.com/Arkosh744/auth-service-api/internal/pkg/encrypt"
	"github.com/Arkosh744/auth-service-api/internal/pkg/token"
	"github.com/Arkosh744/auth-service-api/internal/sys"
	"github.com/Arkosh744/auth-service-api/internal/sys/codes"
)

func (s *service) GetRefreshToken(ctx context.Context, username, password string) (string, error) {
	// TODO: add cache?
	userInfo, err := s.userRepository.Get(ctx, username)
	if err != nil {
		return "", err
	}

	if !encrypt.VerifyPassword(userInfo.Password, password) {
		return "", sys.NewCommonError("invalid password", codes.Aborted)
	}

	refresh, err := token.GenerateToken(&userInfo.User, s.authConfig.RefreshTokenSecretKey(), s.authConfig.RefreshTokenExpirationMinutes())
	if err != nil {
		return "", sys.NewCommonError("failed to generate token", codes.Internal)
	}

	return refresh, nil
}
