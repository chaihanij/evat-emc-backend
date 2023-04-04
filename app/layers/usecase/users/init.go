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
	Login(ctx context.Context, input *entities.Login) (*entities.User, error)
	ChangePassword(ctx context.Context, input *entities.ResetPassword) (*entities.User, error)
	ResetPassword(ctx context.Context, input *entities.ResetPassword) (*entities.User, error)
	//
	CreateUser(ctx context.Context, input *entities.User) (*entities.User, error)
	DeleteUser(ctx context.Context, input *entities.UserFilter) error
	FindAllUser(ctx context.Context, input *entities.UserFilter) (*int64, []entities.User, error)
	FindOneUser(ctx context.Context, input *entities.UserFilter) (*entities.User, error)
	UpdateUser(ctx context.Context, input *entities.UserPartialUpdate) (*entities.User, error)
}

func InitUseCase(usersRepo users.Repo) UseCase {
	return &useCase{
		UsersRepo: usersRepo,
	}
}
