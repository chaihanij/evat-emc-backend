package members

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/members/models"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r repo) UpdateMember(ctx context.Context, input *entities.Member) (*entities.Member, error) {
	log.Debugln("DB UpdateMember")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	filter := models.NewMemberFilter(input)
	update := models.UpdateMember(input)
	var member models.Member
	err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionMembers).
		FindOneAndUpdate(ctx, filter, update, opts).
		Decode(&member)
	if err != nil {
		log.WithError(err).Errorln("DB UpdateMember Error")
		return nil, err
	}
	log.WithField("value", member).Debugln("DB UpdateMember")
	return member.ToEntity()
}
