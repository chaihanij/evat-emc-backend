package users

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) UpdateUser(ctx context.Context, filter *entities.UserFilter, input *entities.UserPartialUpdate) (*entities.User, error) {
	return u.UsersRepo.PartialUpdateUser(ctx, filter, input)
}
