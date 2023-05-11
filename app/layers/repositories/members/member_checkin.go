package members

import (
	"context"

	"github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/members/models"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r repo) MemberCheckIn(ctx context.Context, input *entities.MemberCheckIn) (*entities.Member, error) {
	logrus.Debugln("UpdateMember Check In")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	filter := models.NewMemberFilter(input)
	update := models.UpdateMemberCheckIn(input)
	var member models.Member
	err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionMembers).
		FindOneAndUpdate(ctx, filter, update, opts).
		Decode(&member)
	if err != nil {
		logrus.WithError(err).Errorln("DB UpdateMemberCheckIn Error")
		return nil, err
	}
	return member.ToEntity()
}
