package database

import (
	"context"
	"log"
	"time"

	"gitlab.com/chaihanij/evat/app/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectMongoDB connect to mongo database
func ConnectMongoDB() *mongo.Client {
	var (
		ctx, cancel = context.WithTimeout(context.Background(), time.Second*10)
	)
	defer cancel()

	credential := options.Credential{
		Username:   env.MongoDBUser,
		Password:   env.MongoDBPass,
		AuthSource: env.MongoDBName,
	}
	clientOpts := options.Client().ApplyURI(env.MongoDBUrl).SetAuth(credential)

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatalf("connect to :%v err: %s\n", env.MongoDBUrl, err.Error())
	}

	return client
}
