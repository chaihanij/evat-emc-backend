package email

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) CreateEmail(ctx context.Context, data *entities.Email) (*entities.Email, error) {
	return u.EmailRepo.CreateEmail(ctx,data)
}