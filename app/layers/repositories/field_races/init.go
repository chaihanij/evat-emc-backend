package field_races

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
	CountFieldRaces(ctx context.Context, input *entities.FieldRaceFilter) (*int64, error)
	FindOneFieldRace(ctx context.Context, input *entities.FieldRaceFilter) (*entities.FieldRace, error)
	FindTeamAllFieldRace(ctx context.Context, input *entities.FieldRaceFilter) ([]entities.FieldRace, error)

	UploadScoreFieldRace(ctx context.Context, input *entities.FieldRace) (*entities.FieldRace, error)
}

func InitRepo(mongoDBClient *mongo.Client) Repo {
	return &repo{
		MongoDBClient: mongoDBClient,
	}
}
