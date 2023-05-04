package users

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) CreatLastLogin(ctx context.Context, input *entities.LastLogin) (*entities.LastLogin, error) {
	return u.UsersRepo.CreateUserLogin(ctx , input)
}
