package teams

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
	CountTeam(ctx context.Context, input *entities.TeamFilter) (*int64, error)
	CreateTeam(ctx context.Context, input *entities.Team) (*entities.Team, error)
	DeleteOneTeam(ctx context.Context, input *entities.TeamFilter) error
	FindAllTeam(ctx context.Context, input *entities.TeamFilter) ([]entities.Team, error)
	FindOneTeam(ctx context.Context, input *entities.TeamFilter) (*entities.Team, error)
	PartialUpdateTeam(ctx context.Context, input *entities.TeamPartialUpdate) (*entities.Team, error)
	UpdateTeam(ctx context.Context, input *entities.Team) (*entities.Team, error)
}

func InitRepo(mongoDBClient *mongo.Client) Repo {
	return &repo{
		MongoDBClient: mongoDBClient,
	}
}
