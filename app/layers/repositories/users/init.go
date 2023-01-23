package users

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
	CountUser(ctx context.Context, input *entities.UserFilter) (*int64, error)
	CreateUserMinimal(ctx context.Context, input *entities.UserMinimalCreate) (*entities.User, error)
	CreateUser(ctx context.Context, input *entities.UserCreate) (*entities.User, error)
	DeleteUser(ctx context.Context, input *entities.UserFilter) error
	FindOneUser(ctx context.Context, input *entities.UserFilter) (*entities.User, error)
	FindAllUser(ctx context.Context, input *entities.UserFilter) ([]entities.User, error)
	PartialUpdateUser(ctx context.Context, userFilter *entities.UserFilter, userPartialUpdate *entities.UserPartialUpdate) (*entities.User, error)
}

func InitRepo(mongoDBClient *mongo.Client) Repo {
	return &repo{
		MongoDBClient: mongoDBClient,
	}
}
