package user

import (
	"strings"

	"github.com/Arkosh744/auth-service-api/internal/model"
	desc "github.com/Arkosh744/auth-service-api/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToRole(role desc.Role) string {
	switch role {
	case desc.Role_ADMIN:
		return model.RoleAdmin
	case desc.Role_USER:
		return model.RoleUser
	default:
		return model.RoleUnknown
	}
}

func ToRoleDesc(role string) desc.Role {
	switch role {
	case model.RoleAdmin:
		return desc.Role_ADMIN
	case model.RoleUser:
		return desc.Role_USER
	default:
		return desc.Role_NULL
	}
}

func ToUser(user *desc.CreateRequest) *model.User {
	return &model.User{
		Username: user.GetUser().GetUsername(),
		Email:    strings.ToLower(strings.TrimSpace(user.GetUser().GetEmail())),
		Password: user.GetUser().GetUsername(),
		Role:     ToRole(user.GetUser().GetRole()),
	}
}

func ToUserUpdate(req *desc.UpdateRequest) *model.UpdateUser {
	user := &model.UpdateUser{}

	if req.GetNewRole() != desc.Role_NULL {
		user.Role.String = ToRole(req.GetNewRole())
		user.Role.Valid = true
	}

	if req.GetNewUsername() != nil {
		user.Username.String = req.GetNewUsername().GetValue()
		user.Username.Valid = true
	}

	if req.GetNewEmail() != nil {
		user.Email.String = req.GetNewEmail().GetValue()
		user.Email.Valid = true
	}

	if req.GetNewPassword() != nil {
		user.Password.String = req.GetNewPassword().GetValue()
		user.Password.Valid = true
	}

	return user
}

func ToUserGetDesc(user *model.User) *desc.GetResponse {
	return &desc.GetResponse{
		User: &desc.UserInfo{
			Username: user.Username,
			Email:    user.Email,
			Role:     ToRoleDesc(user.Role)},
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}

func ToUserListDesc(users []*model.User) *desc.ListResponse {
	var list []*desc.UserInfo

	for _, user := range users {
		list = append(list, &desc.UserInfo{
			Username: user.Username,
			Email:    user.Email,
			Role:     ToRoleDesc(user.Role)})
	}

	return &desc.ListResponse{Users: list}
}
