package visit

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
	FindOneVisited(ctx context.Context) (*entities.Visited, error)
	CreateVisit(ctx context.Context, input *entities.UpdateVisit) (*entities.UpdateVisit, error)
}

func InitRepo(mongoDBClient *mongo.Client) Repo {
	return &repo{
		MongoDBClient: mongoDBClient,
	}
}
