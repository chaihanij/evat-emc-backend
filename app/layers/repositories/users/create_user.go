package users

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/errors"
	"gitlab.com/chaihanij/evat/app/layers/repositories/users/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r repo) CreateUser(ctx context.Context, input *entities.User) (*entities.User, error) {
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	user := models.NewUser(input)
	result, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionUsers).
		InsertOne(ctx, user)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"is_duplicate_key_error": mongo.IsDuplicateKeyError(err),
		}).Errorln("DB CreateUser Error")
		if mongo.IsDuplicateKeyError(err) {
			return nil, errors.DuplicateKeyError{Message: err.Error()}
		}
		return nil, err
	}
	if result.InsertedID == 0 {
		return nil, err
	}
	return user.ToEntity()
}
