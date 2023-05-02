package user_v1

import (
	"context"
	"fmt"

	converter "github.com/Arkosh744/auth-grpc/internal/converter/user"
	desc "github.com/Arkosh744/auth-grpc/pkg/user_v1"
	"github.com/Arkosh744/auth-grpc/pkg/validator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (response *desc.UpdateResponse, err error) {
	err = validateUpdateRequest(req)
	if err != nil {
		return response, status.Errorf(codes.InvalidArgument, "Request validation failed: %v", err)
	}

	_, err = i.userService.Get(ctx, req.Username)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "User not found: %v", err)
	}

	user, err := converter.ToUpdateUser(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error converting request to user: %v", err)
	}

	err = i.userService.Update(ctx, req.GetUsername(), user)
	if err != nil {
		return nil, err
	}

	return &desc.UpdateResponse{}, nil
}

func validateUpdateRequest(req *desc.UpdateRequest) error {
	// for now, simple update password without confirmation
	if req.GetNewPassword() != nil && !validator.IsPasswordValid(req.GetNewPassword().GetValue()) {
		return fmt.Errorf(ErrNotValidPassword)
	}

	if req.GetNewEmail() != nil && !validator.IsValidEmail(req.GetNewEmail().GetValue()) {
		return fmt.Errorf(ErrNotValidEmail)
	}

	if req.GetNewUsername() != nil && !validator.IsUsernameValid(req.GetNewUsername().GetValue()) {
		return fmt.Errorf(ErrNotValidUsername)
	}

	return nil
}
