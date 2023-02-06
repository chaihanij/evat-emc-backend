package assignment_teams

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/errors"
	"gitlab.com/chaihanij/evat/app/layers/repositories/assignments/models"
)

func (r repo) DeleteAssignmentTeam(ctx context.Context, input *entities.AssignmentTeamFilter) error {
	log.Debugln("DB AssignmentTeamFilter")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	filter := models.NewAssignmentFilter(input)
	result, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionAssignmentTeams).
		DeleteOne(ctx, filter, nil)

	if err != nil {
		log.WithError(err).Errorln("DB AssignmentTeamFilter Error")
		return err
	}

	if result.DeletedCount < 1 {
		return errors.RecordNotFoundError{Message: constants.DataNotFound}
	}

	return nil
}
