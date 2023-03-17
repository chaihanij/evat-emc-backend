package field_race_teams

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
	CountFieldRaceTeamFilter(ctx context.Context, input *entities.FieldRaceTeamFilter) (*int64, error)
	FindAllFieldRaceTeams(ctx context.Context, input *entities.FieldRaceTeamFilter) ([]entities.FieldRaceTeam, error)
	CreateFieldRaceTeam(ctx context.Context, input *entities.FieldRaceTeam) (*entities.FieldRaceTeam, error)
}

func InitRepo(mongoDBClient *mongo.Client) Repo {
	return &repo{
		MongoDBClient: mongoDBClient,
	}
}
