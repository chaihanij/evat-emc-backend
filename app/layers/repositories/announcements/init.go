package announcements

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
	CountAnnouncement(ctx context.Context, input *entities.AnnouncementFilter) (*int64, error)
	CreateAnnouncement(ctx context.Context, input *entities.Announcement) (*entities.Announcement, error)
	DeleteAnnouncement(ctx context.Context, input *entities.AnnouncementFilter) error
	FindAllAnnouncement(ctx context.Context, input *entities.AnnouncementFilter) ([]entities.Announcement, error)
	FindOneAnnouncement(ctx context.Context, input *entities.AnnouncementFilter) (*entities.Announcement, error)
	PartialUpdateAnnouncement(ctx context.Context, input *entities.AnnouncementPartialUpdate) (*entities.Announcement, error)
	UpdateAnnouncement(ctx context.Context, input *entities.Announcement) (*entities.Announcement, error)
}

func InitRepo(mongoDBClient *mongo.Client) Repo {
	return &repo{
		MongoDBClient: mongoDBClient,
	}
}
