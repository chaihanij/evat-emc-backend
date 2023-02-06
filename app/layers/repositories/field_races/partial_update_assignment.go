package field_races

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/field_races/models"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r repo) PartialUpdateFieldRace(ctx context.Context, input *entities.FieldRacePartialUpdate) (*entities.FieldRace, error) {
	log.Debugln("DB PartialUpdateMember")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	filter := models.NewFieldRaceFilter(input)
	update := models.PartialUpdateFieldRace(input)
	var fieldRace models.FieldRace
	err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionFieldRaces).
		FindOneAndUpdate(ctx, filter, update, opts).
		Decode(&fieldRace)
	if err != nil {
		log.WithError(err).Errorln("DB PartialUpdateFieldRace Error")
		return nil, err
	}
	log.WithField("value", fieldRace).Debugln("DB PartialUpdateFieldRace")
	return fieldRace.ToEntity()
}
