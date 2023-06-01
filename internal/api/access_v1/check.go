package access_v1

import (
	"context"

	desc "github.com/Arkosh744/auth-service-api/pkg/access_v1"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CheckAccess(ctx context.Context, req *desc.CheckAccessRequest) (*empty.Empty, error) {
	err := i.accessService.CheckAccess(ctx, req.GetEndpoint())
	if err != nil {
		return &empty.Empty{}, status.Error(codes.PermissionDenied, "access denied")
	}

	return &empty.Empty{}, nil
}
