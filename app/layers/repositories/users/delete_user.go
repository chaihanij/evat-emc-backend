package users

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/errors"
	"gitlab.com/chaihanij/evat/app/layers/repositories/users/models"
)

func (r repo) DeleteUser(ctx context.Context, input *entities.UserFilter) error {
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	filter := models.NewUserFilter(input)
	res, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionUsers).
		DeleteOne(ctx, filter)

	if err != nil {
		log.WithError(err).Errorln("DB DeleteUser Error")
		return err
	}

	if res.DeletedCount < 1 {
		return errors.RecordNotFoundError{Message: constants.DataNotFound}
	}

	return nil
}
