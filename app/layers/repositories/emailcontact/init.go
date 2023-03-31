package emailcontact

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
	CreateEmailContact(ctx context.Context, input *entities.CreateContactEmail) (*entities.CreateContactEmail, error)
}

func InitRepo(mongoDBClient *mongo.Client) Repo {
	return &repo{
		MongoDBClient: mongoDBClient,
	}
}
