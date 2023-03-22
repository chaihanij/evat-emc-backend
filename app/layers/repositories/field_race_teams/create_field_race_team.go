package field_race_teams

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/errors"
	"gitlab.com/chaihanij/evat/app/layers/repositories/field_race_teams/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r repo) CreateFieldRaceTeam(ctx context.Context, input *entities.FieldRaceTeam) (*entities.FieldRaceTeam, error) {
	log.Debugln("DB FieldRaceTeams")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()

	fieldRaceTeam := models.NewFieldRaceTeam(input)
	result, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionFieldRaceTeams).
		InsertOne(ctx, fieldRaceTeam)
	if err != nil {
		log.WithError(err).Errorln("DB FieldRaceTeams Error")
		if mongo.IsDuplicateKeyError(err) {
			return nil, errors.DuplicateKeyError{Message: err.Error()}
		}
		return nil, err
	}
	if result.InsertedID == 0 {
		return nil, err
	}

	// fieldRace := models.FieldRaceTeam(input)

	log.WithField("value", fieldRaceTeam).Debugln("DB FieldRaceTeams")
	return fieldRaceTeam.ToEntity()
}
