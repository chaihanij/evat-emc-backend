package field_races

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/field_races/models"
)

func (r repo) CreateFieldRace(ctx context.Context, input *entities.FieldRace) (*entities.FieldRace, error) {
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	fieldRace := models.NewFieldRace(input)
	result, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionFieldRaces).
		InsertOne(ctx, fieldRace)
	if err != nil {
		log.WithError(err).Errorln("DB CreateFieldRace Error")
		return nil, err
	}
	if result.InsertedID == 0 {
		return nil, err
	}
	log.WithField("value", fieldRace).Debugln("DB CreateFieldRace")
	return fieldRace.ToEntity()
}
