package assignment_teams

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/assignment_teams/models"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r repo) UpdateAssignmentTeam(ctx context.Context, input *entities.AssignmentTeam) (*entities.AssignmentTeam, error) {
	log.Debugln("DB UpdateAssignmentTeam")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	filter := models.NewAssignmentTeamFilter(input)
	update := models.UpdateAssignmentTeam(input)
	var assignmentTeam models.AssignmentTeam
	err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionAssignmentTeams).
		FindOneAndUpdate(ctx, filter, update, opts).
		Decode(&assignmentTeam)
	if err != nil {
		log.WithError(err).Errorln("DB UpdateAssignmentTeam Error")
		return nil, err
	}
	log.WithField("value", assignmentTeam).Debugln("DB UpdateAssignmentTeam")
	return assignmentTeam.ToEntity()
}
