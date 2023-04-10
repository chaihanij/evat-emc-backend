package field_races

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/field_races/models"
)

func (r repo) UploadScoreFieldRace(ctx context.Context, input *entities.FieldRace) (*entities.FieldRace, error) {
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()

	filter := models.NewFieldRaceFilter(input)
	update := models.UploadScoreFieldRace(input)

	var FieldRace models.FieldRace

	err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionFieldRaces).
		FindOneAndUpdate(ctx, filter, update, nil).
		Decode(&FieldRace)
	if err != nil {
		log.WithError(err).Errorln("DB UpdateAssignmentError Error")
		return nil, err
	}
	log.WithField("value", FieldRace).Debugln("DB UpdateAssignment")
	return FieldRace.ToEntity()
}
