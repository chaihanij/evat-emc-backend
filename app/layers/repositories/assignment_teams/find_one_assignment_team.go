package assignment_teams

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/assignment_teams/models"
)

func (r repo) FindOneAssignmentTeam(ctx context.Context, input *entities.AssignmentTeamFilter) (*entities.AssignmentTeam, error) {
	log.Debugln("DB FindOneAssignment")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	filter := models.NewAssignmentTeamFilter(input)
	var assignmentTeam models.AssignmentTeam
	err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionAssignmentTeams).
		FindOne(ctx, filter, nil).
		Decode(&assignmentTeam)
	if err != nil {
		log.WithError(err).Errorln("DB FindOneAssignmentTeam Error")
		return nil, err
	}

	log.WithField("value", assignmentTeam).Debugln("DB FindOneAssignmentTeam")
	return assignmentTeam.ToEntity()
}
