package announcement_teams

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
	CountAnnouncementTeam(ctx context.Context, input *entities.AssignmentTeamFilter) (*int64, error)
	CreateAnnouncementTeam(ctx context.Context, input *entities.AnnouncementTeam) (*entities.AnnouncementTeam, error)
	DeleteAnnouncementTeam(ctx context.Context, input *entities.AnnouncementTeamFilter) error
	FindAllAnnouncementTeam(ctx context.Context, input *entities.AnnouncementTeamFilter) ([]entities.AnnouncementTeam, error)
	FindOneAnnouncementTeam(ctx context.Context, input *entities.AnnouncementTeamFilter) (*entities.AnnouncementTeam, error)
	PartialUpdateAnnouncementTeam(ctx context.Context, input *entities.AnnouncementTeamPartialUpdate) (*entities.AnnouncementTeam, error)
	UpdateAnnouncementTeam(ctx context.Context, input *entities.AnnouncementTeam) (*entities.AnnouncementTeam, error)
}

func InitRepo(mongoDBClient *mongo.Client) Repo {
	return &repo{
		MongoDBClient: mongoDBClient,
	}
}
