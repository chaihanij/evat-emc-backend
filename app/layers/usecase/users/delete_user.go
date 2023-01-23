package users

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) DeleteUser(ctx context.Context, input *entities.UserFilter) error {
	return u.UsersRepo.DeleteUser(ctx, input)
}
