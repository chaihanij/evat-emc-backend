package assignment_teams

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
	CountAssignmentTeam(ctx context.Context, input *entities.AssignmentTeamFilter) (*int64, error)
	CreateAssignmentTeam(ctx context.Context, input *entities.AssignmentTeam) (*entities.AssignmentTeam, error)
	DeleteAssignmentTeam(ctx context.Context, input *entities.AssignmentTeamFilter) error
	FindAllAssignmentTeam(ctx context.Context, input *entities.AssignmentTeamFilter) ([]entities.AssignmentTeam, error)
	FindOneAssignmentTeam(ctx context.Context, input *entities.AssignmentTeamFilter) (*entities.AssignmentTeam, error)
	PartialUpdateAssignmentTeam(ctx context.Context, input *entities.AssignmentTeamPartialUpdate) (*entities.AssignmentTeam, error)
	UpdateAssignmentTeam(ctx context.Context, input *entities.AssignmentTeam) (*entities.AssignmentTeam, error)

	UpdateAssignmentTeamPushDocument(ctx context.Context, input *entities.AssignmentTeamPartialUpdate, documentUUID string) (*entities.AssignmentTeam, error)
	UpdateAssignmentTeamPullDocument(ctx context.Context, input *entities.AssignmentTeamPartialUpdate, documentUUID string) (*entities.AssignmentTeam, error)
}

func InitRepo(mongoDBClient *mongo.Client) Repo {
	return &repo{
		MongoDBClient: mongoDBClient,
	}
}
