package email

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/layers/repositories/email"
)

type useCase struct {
	EmailRepo email.Repo
}

type UseCase interface {
	CreateEmail(ctx context.Context, data *entities.Email) (*entities.Email, error)
}

func InitUseCase(emailRepo email.Repo) UseCase {
	return &useCase{
		EmailRepo: emailRepo,
	}

}
