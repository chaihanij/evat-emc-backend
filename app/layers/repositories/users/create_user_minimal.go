package users

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/users/models"
)

func (r repo) CreateUserMinimal(ctx context.Context, input *entities.UserMinimalCreate) (*entities.User, error) {
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()

	user := models.User{}
	user.ParseToMinimalModel(input)
	result, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionUsers).
		InsertOne(ctx, user)
	if err != nil {
		log.WithError(err).Errorln("DB CreateUserMinimal Error")
		return nil, err
	}
	if result.InsertedID == 0 {
		return nil, err
	}
	return user.ToEntity()
}
