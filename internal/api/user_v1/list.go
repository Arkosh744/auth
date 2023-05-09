package user_v1

import (
	"context"

	converter "github.com/Arkosh744/auth-grpc/internal/converter/user"
	desc "github.com/Arkosh744/auth-grpc/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) List(ctx context.Context, _ *emptypb.Empty) (*desc.ListResponse, error) {
	users, records, err := i.userService.List(ctx)
	if err != nil {
		i.log.Error("error list users", "error", err)

		return nil, err
	}

	return converter.ToListResponse(users, records), nil
}
