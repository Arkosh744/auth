package user_v1

import (
	"context"

	desc "github.com/Arkosh744/auth-grpc/pkg/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	if _, err := i.userService.Get(ctx, req.GetUsername()); err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found: %v", err)
	}

	if err := i.userService.Delete(ctx, req.GetUsername()); err != nil {
		return nil, status.Errorf(codes.Internal, "error delete user: %v", err)
	}

	return &emptypb.Empty{}, nil
}
