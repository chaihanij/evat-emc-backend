package members

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/members/models"
)

func (r repo) CreateMember(ctx context.Context, input *entities.Member) (*entities.Member, error) {
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	member := models.NewMember(input)
	result, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionMembers).
		InsertOne(ctx, member)
	if err != nil {
		log.WithError(err).Errorln("DB CreateMember Error")
		return nil, err
	}
	if result.InsertedID == 0 {
		return nil, err
	}
	log.WithField("value", member).Debugln("DB CreateMember")
	return member.ToEntity()
}
