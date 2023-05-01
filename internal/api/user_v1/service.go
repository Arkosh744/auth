package user_v1

import (
	"github.com/Arkosh744/auth-grpc/internal/service/user"
	desc "github.com/Arkosh744/auth-grpc/pkg/user_v1"
)

type Implementation struct {
	desc.UnimplementedUserServer

	noteService user.Service
}

func NewImplementation(noteService user.Service) *Implementation {
	return &Implementation{
		noteService: noteService,
	}
}
