package users

import (
	"context"

	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/env"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r repo) Config() ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), env.MongoDBRequestTimeout)
	defer cancel()
	coll := r.MongoDBClient.Database(env.MongoDBName).Collection(constants.CollectionUsers)
	return coll.Indexes().CreateMany(
		ctx,
		[]mongo.IndexModel{
			{
				Keys:    bson.D{{Key: "uid", Value: 1}},
				Options: options.Index().SetUnique(true),
			},
			{
				Keys:    bson.D{{Key: "email", Value: 1}, {Key: "year", Value: 1}},
				Options: options.Index().SetUnique(true),
			},
		},
	)
}
