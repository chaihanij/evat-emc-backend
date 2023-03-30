package email

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	MongoDBClient *mongo.Client
}

type Repo interface {
	Config() ([]string, error)

	CreateEmail(ctx context.Context, input *entities.Email) (*entities.Email, error)
}

func InitRepo(mongoDBClient *mongo.Client) Repo {
	return &repo{
		MongoDBClient: mongoDBClient,
	}
}
