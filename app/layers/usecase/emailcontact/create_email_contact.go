package emailcontact

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) CraeteEmailContact(ctx context.Context, data *entities.CreateContactEmail) (*entities.CreateContactEmail, error) {
	return u.EmailContactRepo.CreateEmailContact(ctx, data)
}
