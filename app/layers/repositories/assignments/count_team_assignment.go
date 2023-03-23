package assignments

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/assignments/models"
)

func (r repo) CountTeamAssignment(ctx context.Context, input *entities.AssignmentFilter) (*int64, error) {
	log.Debugln("DB CountAssignment")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()

	filter := models.NewAssignmentFilter(input)
	count, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionAssignments).
		CountDocuments(ctx, filter)

	if err != nil {
		log.WithError(err).Errorln("DB CountAssignment Error")
		return nil, err
	}
	log.WithField("value", count).Debugln("DB CountAssignment")
	return &count, nil
}
