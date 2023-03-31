package emailcontact

import (
	"context"

	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/emailcontact/models"
)

func (r repo) CreateEmailContact(ctx context.Context, input *entities.CreateContactEmail) (*entities.CreateContactEmail, error) {
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	emailContact := models.NewEmailContact(input)
	// models.SendEmail(input)
	_, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionEmailContact).
		InsertOne(ctx, emailContact)

	if err != nil {
		return nil, err
	}

	return emailContact.ToEntity()
}
