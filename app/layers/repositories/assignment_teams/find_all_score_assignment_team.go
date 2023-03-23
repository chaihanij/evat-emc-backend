package assignment_teams

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/assignment_teams/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (r repo) FindAllscoreAssignmentTeam(ctx context.Context, input *entities.AssignmentTeamFilter) ([]entities.AssignmentTeamScore, error) {
	log.Debugln("DB FindAllAssignmentTeam")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	// findOptions := r.getFindOptions(input)
	log.Debugln("input :: 1234", *input.TeamUUID )
	filter := []bson.M{
		{
			"$match": bson.M{
				"team_uuid": *input.TeamUUID,
			},
		},
		{
			"$lookup": bson.M{
				"from":         "assignments",
				"localField":   "assignment_uuid",
				"foreignField": "uuid",
				"as":           "assignments",
			},
		},
		{
			"$unwind": "$assignments",
		},
		{
			"$project": bson.M{
				"_id":         1,
				"title":       "$assignments.title",
				"assignments": "$assignments.full_score",
			},
		},
	}
	cursor, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionAssignmentTeams).
		Aggregate(ctx, filter)
		// Find(ctx, filter, findOptions)
	if err != nil {
		log.WithError(err).Errorln("DB FindAllAssignmentTeam Error")
		return nil, err
	}
	var assignmentTeamsScore models.AssignmentTeamScores
	err = cursor.All(ctx, &assignmentTeamsScore)
	if err != nil {
		log.WithError(err).Errorln("DB FindAllAssignmentTeam Error")
		return nil, err
	}
	log.WithField("value", assignmentTeamsScore).Debugln("DB FindAllAssignmentTeam")
	return assignmentTeamsScore.ToEntity(), nil
}
