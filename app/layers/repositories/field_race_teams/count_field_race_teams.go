package field_race_teams


import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/field_race_teams/models"
)

func (r repo) CountFieldRaceTeamFilter(ctx context.Context, input *entities.FieldRaceTeamFilter) (*int64, error) {
	log.Debugln("DB FieldRaceTeams func Count")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()

	filter := models.NewFieldRaceTeamFilter(input)
	count, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionFieldRaceTeams).
		CountDocuments(ctx, filter)

	if err != nil {
		log.WithError(err).Errorln("DB FieldRaceTeams Error")
		return nil, err
	}
	log.WithField("value", count).Debugln("DB FieldRaceTeams")
	return &count, nil
}
