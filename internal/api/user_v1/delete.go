package user_v1

import (
	"context"

	desc "github.com/Arkosh744/auth-grpc/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := i.userService.Delete(ctx, req.GetUsername())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
