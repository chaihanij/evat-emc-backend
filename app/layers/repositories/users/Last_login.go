package users

import (
	"context"

	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/users/models"
)

func (r repo) CreateUserLogin(ctx context.Context, input *entities.LastLogin) (*entities.LastLogin, error) {

	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	lastLogin := models.NewLastLogin(input)
	_, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionLastLogin).
		InsertOne(ctx, lastLogin)

	if err != nil {
		return nil, err
	}
	return lastLogin.ToEntity()

}
