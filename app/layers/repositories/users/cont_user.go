package users

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/users/models"
)

func (r repo) CountUser(ctx context.Context, input *entities.UserFilter) (*int64, error) {
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()

	filter := models.FilterUserCriteria(input)
	count, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionUsers).
		CountDocuments(ctx, filter)

	if err != nil {
		log.WithError(err).Errorln("DB CountUser Error")
		return nil, err
	}
	return &count, nil
}
