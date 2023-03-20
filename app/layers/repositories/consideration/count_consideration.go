package consideration

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/assignments/models"
)

func (r repo) CountConsideration(ctx context.Context, input *entities.Consideration) (*int64, error) {
	log.Debugln("CountAnnouncement")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()

	filter := models.NewAssignmentFilter(input)
	count, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionFieldRaceTeams).
		CountDocuments(ctx, filter)

	if err != nil {
		log.WithError(err).Errorln("CountAnnouncement Error")
		return nil, err
	}
	log.WithField("value", count).Debugln("CountConsideration")
	return &count, nil
}
