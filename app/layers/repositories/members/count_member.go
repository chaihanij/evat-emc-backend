package members

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/members/models"
)

func (r repo) CountMember(ctx context.Context, input *entities.MemberFilter) (*int64, error) {
	log.Debugln("DB CountMember")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()

	filter := models.NewMemberFilter(input)
	count, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionMembers).
		CountDocuments(ctx, filter)

	if err != nil {
		log.WithError(err).Errorln("DB CountMember Error")
		return nil, err
	}
	log.WithField("value", count).Debugln("DB CountMember")
	return &count, nil
}
