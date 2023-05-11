package user_v1

import (
	"context"

	converter "github.com/Arkosh744/auth-service-api/internal/converter/user"
	desc "github.com/Arkosh744/auth-service-api/pkg/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	user, err := i.userService.Get(ctx, req.GetUsername())
	if err != nil {
		if status.Code(err) == codes.Unknown {
			i.log.Error("error get user", "error", err)

			return nil, status.Errorf(codes.Internal, "error update user: %v", err)
		}

		return nil, err
	}

	return converter.ToGetResponse(user), nil
}
