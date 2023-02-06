package field_races

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/errors"
	"gitlab.com/chaihanij/evat/app/layers/repositories/field_races/models"
)

func (r repo) DeleteFieldRace(ctx context.Context, input *entities.FieldRaceFilter) error {
	log.Debugln("DB DeleteOnFieldRace")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	filter := models.NewFieldRaceFilter(input)
	result, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionFieldRaces).
		DeleteOne(ctx, filter, nil)

	if err != nil {
		log.WithError(err).Errorln("DB DeleteFieldRace Error")
		return err
	}

	if result.DeletedCount < 1 {
		return errors.RecordNotFoundError{Message: constants.DataNotFound}
	}

	return nil
}
