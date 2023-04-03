package users

import (
	"context"

	"github.com/AlekSi/pointer"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

func (u useCase) ChangePassword(ctx context.Context, input *entities.ResetPassword) (*entities.User, error) {
	if input.ConfirmPassword != input.NewPassword {
		return nil, errors.ParameterError{Message: "password not match"}
	}
	user := &entities.UserPartialUpdate{
		UID:      input.UID,
		Password: pointer.ToString(input.ConfirmPassword),
	}
	return u.UsersRepo.PartialUpdateUser(ctx, user)
}
