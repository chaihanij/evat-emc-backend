package email

import (
	"context"

	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/email/models"
)

func (r repo) CreateEmail(ctx context.Context, input *entities.Email) (*entities.Email, error) {
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()

	email := models.NewEmail(input)

	_, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionEmail).
		InsertOne(ctx, email)

	if err != nil {
		return nil, err
	}
	return email.ToEntity()

}
