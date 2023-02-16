package announcements

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) UpdateAnnouncements(ctx context.Context, input *entities.AnnouncementPartialUpdate) (*entities.Announcement, error) {
	return u.AnnouncementsRepo.PartialUpdateAnnouncement(ctx, input)
}
