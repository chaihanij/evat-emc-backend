package field_races

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/field_races/models"
)

func (r repo) FindOneFieldRace(ctx context.Context, input *entities.FieldRaceFilter) (*entities.FieldRace, error) {
	log.Debugln("DB FindOneFieldRace")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	filter := models.NewFieldRaceFilter(input)
	var fieldRace models.FieldRace
	err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionFieldRaces).
		FindOne(ctx, filter, nil).
		Decode(&fieldRace)
	if err != nil {
		log.WithError(err).Errorln("DB FindOneFieldRace Error")
		return nil, err
	}
	log.WithField("value", fieldRace).Debugln("DB FindOneFieldRace")
	return fieldRace.ToEntity()
}
