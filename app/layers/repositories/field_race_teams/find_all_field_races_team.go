package field_race_teams

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/field_race_teams/models"
)

func (r repo) FindAllAssignment(ctx context.Context, input *entities.FieldRaceTeamFilter) ([]entities.FieldRaceTeam, error) {
	log.Debugln("DB FindAllAssignment")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	findOptions := r.getFindOptions(input)
	filter := models.NewAssignmentTeam(input)
	cursor, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionFieldRaceTeams).
		Find(ctx, filter, findOptions)
	if err != nil {
		log.WithError(err).Errorln("DB FindAllAssignment Error")
		return nil, err
	}
	var fieldRaceTeams models.FieldRaceTeams
	err = cursor.All(ctx, &fieldRaceTeams)
	if err != nil {
		log.WithError(err).Errorln("DB FindAllAssignment Error")
		return nil, err
	}
	log.WithField("value", fieldRaceTeams).Debugln("DB FindAllAssignment")
	return fieldRaceTeams.ToEntity(), nil
}
