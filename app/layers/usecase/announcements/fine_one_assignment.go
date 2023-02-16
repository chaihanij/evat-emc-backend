package announcements

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) FindOneAnnouncements(ctx context.Context, input *entities.AnnouncementFilter) (*entities.Announcement, error) {
	return u.AnnouncementsRepo.FindOneAnnouncement(ctx, input)
}
