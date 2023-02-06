package assignment_teams

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/assignment_teams/models"
)

func (r repo) CreateAssignmentTeam(ctx context.Context, input *entities.AssignmentTeam) (*entities.AssignmentTeam, error) {
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	assignmentTeam := models.NewAssignmentTeam(input)
	result, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionAssignmentTeams).
		InsertOne(ctx, assignmentTeam)
	if err != nil {
		log.WithError(err).Errorln("DB CreateAssignment Error")
		return nil, err
	}
	if result.InsertedID == 0 {
		return nil, err
	}
	log.WithField("value", assignmentTeam).Debugln("DB CreateAssignment")
	return assignmentTeam.ToEntity()
}
