package albums

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/albums/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r repo) PushImages(ctx context.Context, uuid string, input []string) (*entities.Album, error) {
	log.Debugln("DB PushImages")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	filter := bson.M{"uuid": uuid}
	log.WithField("filter", filter).Debugln("PushImages filter")
	update := models.PushImages(input)
	var member models.Album
	err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionAlbums).
		FindOneAndUpdate(ctx, filter, update, opts).
		Decode(&member)
	if err != nil {
		log.WithError(err).Errorln("DB PushImages Error")
		return nil, err
	}
	log.WithField("value", member).Debugln("DB PushImages")
	return member.ToEntity()
}

func (r repo) PullImage(ctx context.Context, uuid string, input string) (*entities.Album, error) {
	log.Debugln("DB PullImage")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	filter := bson.M{"uuid": uuid}
	update := models.PullImage(input)
	var member models.Album
	err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionAlbums).
		FindOneAndUpdate(ctx, filter, update, opts).
		Decode(&member)
	if err != nil {
		log.WithError(err).Errorln("DB PullImage Error")
		return nil, err
	}
	log.Debugln("DB PullImage")
	return member.ToEntity()
}
