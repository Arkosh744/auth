package user_v1

import (
	"github.com/Arkosh744/auth-service-api/internal/service/user"
	desc "github.com/Arkosh744/auth-service-api/pkg/user_v1"
	"go.uber.org/zap"
)

type Implementation struct {
	desc.UnimplementedUserServer

	log         *zap.SugaredLogger
	userService user.Service
}

func NewImplementation(noteService user.Service, log *zap.SugaredLogger) *Implementation {
	return &Implementation{
		userService: noteService,
		log:         log,
	}
}
