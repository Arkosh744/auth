package user_v1

import (
	"context"

	desc "github.com/Arkosh744/auth-grpc/pkg/user_v1"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*desc.DeleteResponse, error) {
	err := i.userService.Delete(ctx, req.GetUsername())
	if err != nil {
		return nil, err
	}

	return &desc.DeleteResponse{}, nil
}
