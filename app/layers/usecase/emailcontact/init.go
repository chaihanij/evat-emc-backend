package emailcontact

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
	emailcontact "gitlab.com/chaihanij/evat/app/layers/repositories/emailcontact"
)

type useCase struct {
	EmailContactRepo emailcontact.Repo
}

type UseCase interface {
	CraeteEmailContact(ctx context.Context, data *entities.CreateContactEmail) (*entities.CreateContactEmail, error)
}

func InitUseCase(email_contact emailcontact.Repo) UseCase {
	return &useCase{
		EmailContactRepo: email_contact,
	}
}
