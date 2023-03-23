package assignments

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/assignments/models"
)

func (r repo) FindTeamAssignment(ctx context.Context, input *entities.AssignmentFilter) ([]entities.TeamAssignment, error) {
	log.Debugln("DB FindAllTeamAssignment")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	findOptions := r.getFindOptions(input)
	filter := models.NewTeamAssignmentFilter(input)
	cursor, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionAssignments).
		Find(ctx, filter, findOptions)

	if err != nil {
		log.WithError(err).Errorln("DB FindAllTeamAssignment Error")
		return nil, err
	}
	var teamAssignments models.TeamAssignments
	err = cursor.All(ctx, &teamAssignments)
	if err != nil {
		return nil, err
	}
	log.WithField("value", teamAssignments).Debugln("DB FindAllTeamAssignment")

	return teamAssignments.ToEntity(), nil
}
