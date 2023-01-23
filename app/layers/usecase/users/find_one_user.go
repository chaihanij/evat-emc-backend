package users

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) FindOneUser(ctx context.Context, input *entities.UserFilter) (*entities.User, error) {
	return u.UsersRepo.FindOneUser(ctx, input)
}
