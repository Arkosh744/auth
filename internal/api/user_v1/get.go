package user_v1

import (
	"context"

	converter "github.com/Arkosh744/auth-grpc/internal/converter/user"
	desc "github.com/Arkosh744/auth-grpc/pkg/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	user, err := i.userService.Get(ctx, req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "User not found: %v", err)
	}

	response := converter.ToGetResponse(user)
	return response, nil
}
