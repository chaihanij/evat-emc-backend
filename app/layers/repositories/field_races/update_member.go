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

func (r repo) UpdateFieldRace(ctx context.Context, input *entities.FieldRace) (*entities.FieldRace, error) {
	log.Debugln("DB UpdateFieldRace")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	filter := models.NewFieldRaceFilter(input)
	update := models.UpdateFieldRace(input)
	var assignment models.FieldRace
	err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionFieldRaces).
		FindOneAndUpdate(ctx, filter, update, opts).
		Decode(&assignment)
	if err != nil {
		log.WithError(err).Errorln("DB UpdateFieldRace Error")
		return nil, err
	}
	log.WithField("value", assignment).Debugln("DB UpdateFieldRace")
	return assignment.ToEntity()
}
