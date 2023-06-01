package auth_v1

import (
	"github.com/Arkosh744/auth-service-api/internal/service/auth"
	desc "github.com/Arkosh744/auth-service-api/pkg/auth_v1"
)

type Implementation struct {
	desc.UnimplementedAuthV1Server

	authService auth.Service
}

func NewImplementation(authService auth.Service) *Implementation {
	return &Implementation{
		authService: authService,
	}
}
