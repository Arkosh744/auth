package access_v1

import (
	"github.com/Arkosh744/auth-service-api/internal/service/access"
	desc "github.com/Arkosh744/auth-service-api/pkg/access_v1"
)

type Implementation struct {
	desc.UnimplementedAccessV1Server

	accessService access.Service
}

func NewImplementation(accessService access.Service) *Implementation {
	return &Implementation{
		accessService: accessService,
	}
}
