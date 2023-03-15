package userscore

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
	FindAllScore(ctx context.Context, input *entities.ScoreFilter) ([]entities.Score, error)
	CountScore(ctx context.Context, input *entities.ScoreFilter) (*int64, error)
	// CreateScore(ctx context.Context, input *entities.Score) (*entities.Score, error)
	// Create
	//	CreateAssignmentTeam(ctx context.Context, input *entities.AssignmentTeam) (*entities.AssignmentTeam, error)

}

func InitRepo(mongoDBClient *mongo.Client) Repo {
	return &repo{
		MongoDBClient: mongoDBClient,
	}
}
