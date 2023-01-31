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

func (r repo) PushDocument(ctx context.Context, uuid string, input string) (*entities.Member, error) {
	log.Debugln("DB PushDocument")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	filter := bson.M{"uuid": uuid}
	update := models.PushDocument(input)
	var member models.Member
	err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionUsers).
		FindOneAndUpdate(ctx, filter, update, opts).
		Decode(&member)
	if err != nil {
		log.WithError(err).Errorln("DB PushDocument Error")
		return nil, err
	}
	log.WithField("value", member).Debugln("DB PushDocument")
	return member.ToEntity()
}

func (r repo) PullDocument(ctx context.Context, uuid string, input string) (*entities.Member, error) {
	log.Debugln("DB PullDocument")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	filter := bson.M{"uuid": uuid}
	update := models.PullDocument(input)
	var member models.Member
	err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionUsers).
		FindOneAndUpdate(ctx, filter, update, opts).
		Decode(&member)
	if err != nil {
		log.WithError(err).Errorln("DB PullDocument Error")
		return nil, err
	}
	log.Debugln("DB PullDocument")
	return member.ToEntity()
}
