package announcements

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) DeleteAnnouncements(ctx context.Context, input *entities.AnnouncementFilter) error {
	return u.AnnouncementsRepo.DeleteAnnouncement(ctx, input)
}
