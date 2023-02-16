package announcements

import (
	"context"
	"sync"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) countAnnouncement(wg *sync.WaitGroup, ctx context.Context, input *entities.AnnouncementFilter, count chan *int64, errors chan error) {
	defer wg.Done()
	result, err := u.AnnouncementsRepo.CountAnnouncement(ctx, input)
	if err != nil {
		errors <- err
		return
	}
	count <- result
}

func (u useCase) findAllAnnouncement(wg *sync.WaitGroup, ctx context.Context, input *entities.AnnouncementFilter, teams chan []entities.Announcement, errors chan error) {
	defer wg.Done()
	result, err := u.AnnouncementsRepo.FindAllAnnouncement(ctx, input)

	if err != nil {
		errors <- err
		return
	}
	teams <- result
}

func (u useCase) FindAllAnnouncement(ctx context.Context, input *entities.AnnouncementFilter) (*int64, []entities.Announcement, error) {
	count := make(chan *int64, 1)
	announcements := make(chan []entities.Announcement, 1)
	errors := make(chan error, 2)
	var wg sync.WaitGroup
	wg.Add(1)
	go u.countAnnouncement(&wg, ctx, input, count, errors)
	wg.Add(1)
	go u.findAllAnnouncement(&wg, ctx, input, announcements, errors)

	go func() {
		wg.Wait()
		close(count)
		close(announcements)
		close(errors)
	}()

	for err := range errors {
		return nil, nil, err
	}

	totalRecords, result, err := <-count, <-announcements, <-errors
	return totalRecords, result, err
}
