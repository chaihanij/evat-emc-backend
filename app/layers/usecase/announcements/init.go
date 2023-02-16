package announcements

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/layers/repositories/announcements"
)

type useCase struct {
	AnnouncementsRepo announcements.Repo
}

type UseCase interface {
	CreateAnnouncements(ctx context.Context, input *entities.Announcement) (*entities.Announcement, error)
	DeleteAnnouncements(ctx context.Context, input *entities.AnnouncementFilter) error
	FindAllAnnouncement(ctx context.Context, input *entities.AnnouncementFilter) (*int64, []entities.Announcement, error)
	FindOneAnnouncements(ctx context.Context, input *entities.AnnouncementFilter) (*entities.Announcement, error)
	UpdateAnnouncements(ctx context.Context, input *entities.AnnouncementPartialUpdate) (*entities.Announcement, error)
}

func InitUseCase(announcementsRepo announcements.Repo) UseCase {
	return &useCase{
		AnnouncementsRepo: announcementsRepo,
	}
}
