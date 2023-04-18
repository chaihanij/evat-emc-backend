package assignments

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
	CountAssignment(ctx context.Context, input *entities.AssignmentFilter) (*int64, error)
	CreateAssignment(ctx context.Context, input *entities.Assignment) (*entities.Assignment, error)
	DeleteOneAssignment(ctx context.Context, input *entities.AssignmentFilter) error
	FindAllAssignment(ctx context.Context, input *entities.AssignmentFilter) ([]entities.Assignment, error)
	FindOneAssignment(ctx context.Context, input *entities.AssignmentFilter) (*entities.Assignment, error)
	PartialUpdateAssignment(ctx context.Context, input *entities.AssignmentPartialUpdate) (*entities.Assignment, error)
	UpdateAssignment(ctx context.Context, input *entities.Assignment) (*entities.Assignment, error)
	CountTeamAssignment(ctx context.Context, input *entities.AssignmentFilter) (*int64, error)
	FindTeamAssignment(ctx context.Context, input *entities.AssignmentFilter) ([]entities.TeamAssignment, error)

	FindTopicAssignment(ctx context.Context, input *entities.AssignmentFilter) (*entities.ExportAssignmentTopic, error)
	UploadScoreAssignment(ctx context.Context, input *entities.Assignment) (*entities.Assignment, error)
}

func InitRepo(mongoDBClient *mongo.Client) Repo {
	return &repo{
		MongoDBClient: mongoDBClient,
	}
}
