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
		return desc.Role_UNKNOWN
	}
}

func ToUserSpec(user *desc.CreateRequest) *model.UserSpec {
	u := model.UserSpec{
		User: model.User{
			Username: user.GetUser().GetUsername(),
			Email:    strings.ToLower(strings.TrimSpace(user.GetUser().GetEmail())),
			Password: user.GetUser().GetUsername(),
			Role:     ToRole(user.GetUser().GetRole()),
		},
	}

	ToSpec(user, u)

	return &u
}

func ToSpec(user *desc.CreateRequest, u model.UserSpec) {
	switch user.GetUser().GetSpecialization().(type) {
	case *desc.UserInfo_Manager:
		u.Type = model.SpecializationManager
		u.Manager = &model.Manager{
			Level:      user.GetUser().GetManager().GetLevel(),
			Company:    user.GetUser().GetManager().GetCompany(),
			Department: user.GetUser().GetManager().GetDepartment(),
		}
	case *desc.UserInfo_Engineer:
		u.Type = model.SpecializationEngineer
		u.Engineer = &model.Engineer{
			Level:    user.GetUser().GetEngineer().GetLevel(),
			Company:  user.GetUser().GetEngineer().GetCompany(),
			Language: user.GetUser().GetEngineer().GetLanguage(),
		}
	}
}

func ToUserUpdate(req *desc.UpdateRequest) *model.UpdateUser {
	user := &model.UpdateUser{}

	if req.GetNewRole() != desc.Role_UNKNOWN {
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

func ToUserGetDesc(user *model.UserSpec) *desc.GetResponse {
	res := &desc.GetResponse{
		User: &desc.UserInfo{
			Username: user.Username,
			Email:    user.Email,
			Role:     ToRoleDesc(user.Role)},
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}

	ToSpecDesc(user, res.User)

	return res
}

func ToSpecDesc(user *model.UserSpec, resUser *desc.UserInfo) {
	switch user.Specialization.Type {
	case model.SpecializationEngineer:
		resUser.Specialization = &desc.UserInfo_Engineer{Engineer: &desc.Engineer{
			Level:    user.Specialization.Engineer.Level,
			Company:  user.Specialization.Engineer.Company,
			Language: user.Specialization.Engineer.Language,
		}}
	case model.SpecializationManager:
		resUser.Specialization = &desc.UserInfo_Manager{Manager: &desc.Manager{
			Level:      user.Specialization.Manager.Level,
			Company:    user.Specialization.Manager.Company,
			Department: user.Specialization.Manager.Department,
		}}
	}
}

func ToUserListDesc(users []*model.UserSpec) *desc.ListResponse {
	list := make([]*desc.UserInfo, 0, len(users))

	for _, user := range users {
		userDesc := &desc.UserInfo{
			Username: user.Username,
			Email:    user.Email,
			Role:     ToRoleDesc(user.Role),
		}
		ToSpecDesc(user, userDesc)

		list = append(list, userDesc)
	}

	return &desc.ListResponse{Users: list}
}
