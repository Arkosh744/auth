package user_v1

import (
	"context"

	desc "github.com/Arkosh744/auth-service-api/pkg/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	if err := i.userService.Delete(ctx, req.GetUsername()); err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, status.Errorf(codes.Internal, "error update user: %v", err)
		}

		return nil, err
	}

	return &emptypb.Empty{}, nil
}
