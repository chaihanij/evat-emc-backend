package members

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/members/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r repo) getFindOptions(input *entities.MemberFilter) *options.FindOptions {
	findOptions := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})
	if input.PageSize != nil && input.Page != nil {
		pageSize := *input.PageSize
		page := *input.Page
		findOptions.SetLimit(*input.PageSize)
		offset := (page - 1) * pageSize
		findOptions.SetSkip(offset)
	}
	return findOptions
}

func (r repo) FindAllMember(ctx context.Context, input *entities.MemberFilter) ([]entities.Member, error) {
	log.Debugln("DB FindAllMember")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	findOptions := r.getFindOptions(input)
	filter := models.NewMemberFilter(input)
	cursor, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionMembers).
		Find(ctx, filter, findOptions)
	if err != nil {
		log.WithError(err).Errorln("DB FineAllMember Error")
		return nil, err
	}
	var members models.Members
	err = cursor.All(ctx, &members)
	if err != nil {
		log.WithError(err).Errorln("DB FineAllMember Error")
		return nil, err
	}
	// log.WithField("value", members).Debugln("DB FindAllMember")
	return members.ToEntity(), nil
}
