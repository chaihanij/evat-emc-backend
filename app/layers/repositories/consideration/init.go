package consideration

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
	CountConsideration(ctx context.Context, input *entities.Consideration) (*int64, error)
	FindOneConsideration(ctx context.Context, input *entities.ConsiderationFilter) ([]entities.AssignmentScore, error)
	FindConsiderationFieldRaceTeam(ctx context.Context, input *entities.ConsiderationFilter) ([]entities.FieldRaceTeamScore, error)

	AllScore(ctx context.Context, input entities.AllScoreFilter) ([]entities.AllScore, error)
	AllScoreConsiderationAssignment(ctx context.Context, input entities.AllScoreFilter) ([]entities.AllScore, error)
}

func InitRepo(mongoDBClient *mongo.Client) Repo {
	return &repo{
		MongoDBClient: mongoDBClient,
	}
}
