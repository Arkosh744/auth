package user

import (
	"fmt"
	"strings"

	"github.com/Arkosh744/auth-grpc/internal/model"
	desc "github.com/Arkosh744/auth-grpc/pkg/user_v1"
)

func ToUserRole(role string) (userRole model.UserRole, err error) {
	switch role {
	case model.UserRoleAdmin.String():
		return model.UserRoleAdmin, nil
	case model.UserRoleUser.String():
		return model.UserRoleUser, nil
	default:
		return userRole, fmt.Errorf("failed to convert user role: %s", role)
	}
}

func ToUser(user *desc.CreateRequest) (*model.User, error) {
	userRole, err := ToUserRole(user.GetRole())
	if err != nil {
		return nil, err
	}

	return &model.User{
		Username: user.GetUsername(),
		Email:    strings.ToLower(strings.TrimSpace(user.GetEmail())),
		Password: user.GetPassword(),
		Role:     userRole,
	}, nil
}
