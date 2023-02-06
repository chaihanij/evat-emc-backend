package assignment_teams

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/assignment_teams/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r repo) getFindOptions(input *entities.AssignmentTeamFilter) *options.FindOptions {
	findOptions := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})
	if input.PageSize != nil && input.Page != nil {
		pageSize := *input.PageSize
		page := *input.Page
		findOptions.SetLimit(*input.PageSize)
		offset := (page - 1) * pageSize
		findOptions.SetSkip(offset)
	}
	return findOptions
}

func (r repo) FindAllAssignmentTeam(ctx context.Context, input *entities.AssignmentTeamFilter) ([]entities.AssignmentTeam, error) {
	log.Debugln("DB FindAllAssignmentTeam")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	findOptions := r.getFindOptions(input)
	filter := models.NewAssignmentTeamFilter(input)
	cursor, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionAssignmentTeams).
		Find(ctx, filter, findOptions)
	if err != nil {
		log.WithError(err).Errorln("DB FindAllAssignmentTeam Error")
		return nil, err
	}
	var assignmentTeams models.AssignmentTeams
	err = cursor.All(ctx, &assignmentTeams)
	if err != nil {
		log.WithError(err).Errorln("DB FindAllAssignmentTeam Error")
		return nil, err
	}
	log.WithField("value", assignmentTeams).Debugln("DB FindAllAssignmentTeam")
	return assignmentTeams.ToEntity(), nil
}
