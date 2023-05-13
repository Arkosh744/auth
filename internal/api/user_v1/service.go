package user_v1

import (
	"github.com/Arkosh744/auth-service-api/internal/service/user"
	desc "github.com/Arkosh744/auth-service-api/pkg/user_v1"
)

type Implementation struct {
	desc.UnimplementedUserServer

	userService user.Service
}

func NewImplementation(noteService user.Service) *Implementation {
	return &Implementation{
		userService: noteService,
	}
}
