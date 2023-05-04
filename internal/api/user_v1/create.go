package user_v1

import (
	"context"
	"fmt"

	converter "github.com/Arkosh744/auth-grpc/internal/converter/user"
	"github.com/Arkosh744/auth-grpc/internal/model"
	"github.com/Arkosh744/auth-grpc/internal/pkg/encrypt"
	"github.com/Arkosh744/auth-grpc/internal/pkg/validator"
	desc "github.com/Arkosh744/auth-grpc/pkg/user_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*emptypb.Empty, error) {
	err := validateCreateRequest(req)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "request validation failed: %v", err)
	}

	user := converter.ToUser(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting request to user: %v", err)
	}

	if user.Role == model.RoleUnknown {
		return nil, status.Errorf(codes.InvalidArgument, "invalid role provided: %v", user.Role)
	}

	user.Password, err = encrypt.HashPassword(user.Password)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error data hashing: %v", err)
	}

	err = i.userService.Create(ctx, user)
	if err != nil {
		switch status.Code(err) {
		case codes.Unknown:
			return nil, status.Errorf(codes.Internal, "error creating user: %v", err)
		default:
			return nil, err
		}
	}

	return &emptypb.Empty{}, nil
}

func validateCreateRequest(req *desc.CreateRequest) error {
	if !validator.IsPasswordValid(req.GetPassword()) {
		return fmt.Errorf(ErrNotValidPassword)
	}

	if !validator.IsPasswordConfirmed(req.GetPassword(), req.GetPasswordConfirm()) {
		return fmt.Errorf(ErrPasswordConfirmation)
	}

	if !validator.IsValidEmail(req.GetUser().GetEmail()) {
		return fmt.Errorf(ErrNotValidEmail)
	}

	if !validator.IsUsernameValid(req.GetUser().GetUsername()) {
		return fmt.Errorf(ErrNotValidUsername)
	}

	if model.StringToRole(req.GetUser().GetRole()) == model.RoleUnknown {
		return fmt.Errorf("invalid role provided: %v", req.GetUser().GetRole())
	}

	return nil
}
