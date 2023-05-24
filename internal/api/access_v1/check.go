package access_v1

import (
	"context"

	desc "github.com/Arkosh744/auth-service-api/pkg/access_v1"
)

func (i *Implementation) CheckAccess(ctx context.Context, req *desc.CheckAccessRequest) (*desc.CheckAccessResponse, error) {
	isAllowed, err := i.accessService.CheckAccess(ctx, req.GetEndpoint())
	if err != nil {
		return nil, err
	}

	return &desc.CheckAccessResponse{IsAllowed: isAllowed}, nil
}
