package user

import (
	"strings"

	"github.com/Arkosh744/auth-grpc/internal/model"
	desc "github.com/Arkosh744/auth-grpc/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToUser(user *desc.CreateRequest) *model.User {
	return &model.User{
		Username: user.GetUser().GetUsername(),
		Email:    strings.ToLower(strings.TrimSpace(user.GetUser().GetEmail())),
		Password: user.GetUser().GetUsername(),
		Role:     model.StringToRole(user.GetUser().GetRole()),
	}
}

func ToUpdateUser(req *desc.UpdateRequest) *model.User {
	user := &model.User{}

	if req.GetNewRole() != nil {
		user.Role = model.StringToRole(req.GetNewRole().GetValue())
	}

	if req.GetNewUsername() != nil {
		user.Username = req.GetNewUsername().GetValue()
	}

	if req.GetNewEmail() != nil {
		user.Email = req.GetNewEmail().GetValue()
	}

	if req.GetNewPassword() != nil {
		user.Password = req.GetNewPassword().GetValue()
	}

	return user
}

func ToGetResponse(user *model.User) *desc.GetResponse {
	return &desc.GetResponse{
		User: &desc.UserInfo{
			Username: user.Username,
			Email:    user.Email,
			Role:     user.Role.String()},
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}
