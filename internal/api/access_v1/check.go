package access_v1

import (
	"context"
	"github.com/Arkosh744/auth-service-api/internal/service/access"
	"github.com/golang/protobuf/ptypes/empty"

	desc "github.com/Arkosh744/auth-service-api/pkg/access_v1"
)

var ErrAccessDenied = access.ErrAccessDenied

func (i *Implementation) CheckAccess(ctx context.Context, req *desc.CheckAccessRequest) (*empty.Empty, error) {
	err := i.accessService.CheckAccess(ctx, req.GetEndpoint())
	if err != nil {
		if err == ErrAccessDenied {
			return &empty.Empty{}, ErrAccessDenied
		}

		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}
