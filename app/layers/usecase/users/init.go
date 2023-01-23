package users

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/layers/repositories/users"
)

type useCase struct {
	UsersRepo users.Repo
}

type UseCase interface {
	CreateUser(ctx context.Context, input *entities.UserCreate) (*entities.User, error)
	DeleteUser(ctx context.Context, input *entities.UserFilter) error
	FindAllHubCredentials(ctx context.Context, input *entities.UserFilter) (*int64, []entities.User, error)
	FindOneUser(ctx context.Context, input *entities.UserFilter) (*entities.User, error)
	UpdateUser(ctx context.Context, filter *entities.UserFilter, input *entities.UserPartialUpdate) (*entities.User, error)
}

func InitUseCase(usersRepo users.Repo) UseCase {
	return &useCase{
		UsersRepo: usersRepo,
	}
}
