package user_v1

import (
	"context"
	"fmt"

	converter "github.com/Arkosh744/auth-grpc/internal/converter/user"
	"github.com/Arkosh744/auth-grpc/internal/model"
	"github.com/Arkosh744/auth-grpc/internal/pkg/validator"
	desc "github.com/Arkosh744/auth-grpc/pkg/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	err := validateUpdateRequest(req)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "request validation failed: %v", err)
	}

	if err = i.userService.Update(ctx, req.GetUsername(), converter.ToUpdateUser(req)); err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, status.Errorf(codes.Internal, "error update user: %v", err)
		}

		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func validateUpdateRequest(req *desc.UpdateRequest) error {
	// for now, simple update password without any other checks
	if req.GetNewPassword() != nil && !validator.IsPasswordValid(req.GetNewPassword().GetValue()) {
		return fmt.Errorf(ErrNotValidPassword)
	}

	if req.GetNewEmail() != nil && !validator.IsValidEmail(req.GetNewEmail().GetValue()) {
		return fmt.Errorf(ErrNotValidEmail)
	}

	if req.GetNewUsername() != nil && !validator.IsUsernameValid(req.GetNewUsername().GetValue()) {
		return fmt.Errorf(ErrNotValidUsername)
	}

	if req.GetNewRole() != nil && model.StringToRole(req.GetNewRole().GetValue()) == model.RoleUnknown {
		return fmt.Errorf("invalid role provided: %v", req.GetNewRole().GetValue())
	}

	return nil
}
