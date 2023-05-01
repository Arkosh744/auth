package user_v1

import (
	"context"
	"fmt"

	converter "github.com/Arkosh744/auth-grpc/internal/converter/user"
	desc "github.com/Arkosh744/auth-grpc/pkg/user_v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	if !isPasswordConfirmed(req.Password, req.PasswordConfirm) {
		return nil, fmt.Errorf(ErrPasswordConfirmation)
	}

	user, err := converter.ToUser(req)
	if err != nil {
		return nil, err
	}

	err = i.noteService.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{}, nil
}

func isPasswordConfirmed(password string, passwordConfirmation string) bool {
	return password == passwordConfirmation
}
