package userscore

import (
	"context"

	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/env"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r repo) Config() ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), env.MongoDBRequestTimeout)
	defer cancel()
	coll := r.MongoDBClient.Database(env.MongoDBName).Collection(constants.CollectionScore)
	return coll.Indexes().CreateMany(
		ctx,
		[]mongo.IndexModel{
			// {
			// 	Keys:    bson.D{{Key: "uuid", Value: 1}},
			// 	Options: options.Index().SetUnique(true),
			// },
		},
	)
}
