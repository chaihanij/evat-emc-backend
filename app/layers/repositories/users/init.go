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
	CreateUser(ctx context.Context, input *entities.User) (*entities.User, error)
	DeleteUser(ctx context.Context, input *entities.UserFilter) error
	FindOneUser(ctx context.Context, input *entities.UserFilter) (*entities.User, error)
	FindAllUser(ctx context.Context, input *entities.UserFilter) ([]entities.User, error)
	PartialUpdateUser(ctx context.Context, input *entities.UserPartialUpdate) (*entities.User, error)
	UpdateUser(ctx context.Context, userFilter, input *entities.User) (*entities.User, error)
}

func InitRepo(mongoDBClient *mongo.Client) Repo {
	return &repo{
		MongoDBClient: mongoDBClient,
	}
}
