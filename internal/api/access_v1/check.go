package access_v1

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	desc "github.com/Arkosh744/auth-service-api/pkg/access_v1"
)

func (i *Implementation) CheckAccess(ctx context.Context, req *desc.CheckAccessRequest) (*empty.Empty, error) {
	err := i.accessService.CheckAccess(ctx, req.GetEndpoint())
	if err != nil {
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}
