package assignment_teams

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/assignments/models"
)

func (r repo) CountAssignmentTeam(ctx context.Context, input *entities.AssignmentTeamFilter) (*int64, error) {
	log.Debugln("DB CountAssignmentTeam")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()

	filter := models.NewAssignmentFilter(input)
	count, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionAssignmentTeams).
		CountDocuments(ctx, filter)

	if err != nil {
		log.WithError(err).Errorln("DB CountAssignmentTeam Error")
		return nil, err
	}
	log.WithField("value", count).Debugln("DB CountAssignmentTeam")
	return &count, nil
}
