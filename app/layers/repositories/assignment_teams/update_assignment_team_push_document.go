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

func (r repo) UpdateAssignmentTeamPushDocument(ctx context.Context, input *entities.AssignmentTeam, documentUUIDs []string) (*entities.AssignmentTeam, error) {
	log.Debugln("DB UpdateAssignmentTeamPushDocument")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After).SetUpsert(true)
	filter := models.NewAssignmentTeamFilter(input)
	update := models.UpdateAssignmentTeamPushDocument(documentUUIDs)
	var assignmentTeam models.AssignmentTeam
	err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionAssignmentTeams).
		FindOneAndUpdate(ctx, filter, update, opts).
		Decode(&assignmentTeam)
	if err != nil {
		log.WithError(err).Errorln("DB UpdateAssignmentTeamPushDocument Error")
		return nil, err
	}
	log.WithField("value", assignmentTeam).Debugln("DB UpdateAssignmentTeamPushDocument")
	return assignmentTeam.ToEntity()
}
