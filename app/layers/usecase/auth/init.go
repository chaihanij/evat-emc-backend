package auth

import "gitlab.com/chaihanij/evat/app/layers/repositories/users"

type useCase struct {
	UsersRepo users.Repo
}

type UseCase interface {
}

func InitUseCase(usersRepo users.Repo) UseCase {
	return &useCase{
		UsersRepo: usersRepo,
	}
}
