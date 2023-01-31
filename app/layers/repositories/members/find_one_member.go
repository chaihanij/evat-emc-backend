package members

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/members/models"
)

func (r repo) FindOneMember(ctx context.Context, input *entities.MemberFilter) (*entities.Member, error) {
	log.Debugln("DB FindOneMemeber")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	filter := models.NewMemberFilter(input)
	var member models.Member
	err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionUsers).
		FindOne(ctx, filter, nil).
		Decode(&member)
	if err != nil {
		log.WithError(err).Errorln("DB FineOneMemeber Error")
		return nil, err
	}
	log.WithField("value", member).Debugln("DB FindOneMemeber")
	return member.ToEntity()
}
