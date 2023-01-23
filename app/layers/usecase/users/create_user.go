package users

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) CreateUser(ctx context.Context, input *entities.UserCreate) (*entities.User, error) {
	return u.UsersRepo.CreateUser(ctx, input)
}
