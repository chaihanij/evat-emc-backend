package userscore

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	medles "gitlab.com/chaihanij/evat/app/layers/repositories/userscore/models"
	// models "gitlab.com/chaihanij/evat/app/layers/repositories/userscore/models"
	// "gitlab.com/chaihanij/evat/app/layers/repositories/userscore/medles"
)

func (r repo) CountScore(ctx context.Context, input *entities.ScoreFilter) (*int64, error) {
	log.Debugln("DB Score")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	filter := medles.NewScoreFilter(input)

	count, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.Collectionracesteam).
		CountDocuments(ctx, filter)

	if err != nil {
		log.WithError(err).Errorln("DB Countscore Error")
		return nil, err
	}
	log.WithField("value", count).Debugln("DB CountScore")
	return &count, nil
}
