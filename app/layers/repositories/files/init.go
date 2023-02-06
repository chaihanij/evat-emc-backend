package files

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
	CreateFile(ctx context.Context, input *entities.File) (*entities.File, error)
	CreateFiles(ctx context.Context, input []entities.File) ([]entities.File, error)
	FindOneFile(ctx context.Context, input interface{}) (*entities.File, error)
	FindAllFile(ctx context.Context, input interface{}) ([]entities.File, error)
}

func InitRepo(mongoDBClient *mongo.Client) Repo {
	return &repo{
		MongoDBClient: mongoDBClient,
	}
}
