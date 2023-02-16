package announcements

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) CreateAnnouncements(ctx context.Context, input *entities.Announcement) (*entities.Announcement, error) {
	return u.AnnouncementsRepo.CreateAnnouncement(ctx, input)
}
