package assignment_teams

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	MongoDBClient *mongo.Client
}

type Repo interface {
}

func InitRepo(mongoDBClient *mongo.Client) Repo {
	return &repo{
		MongoDBClient: mongoDBClient,
	}
}
