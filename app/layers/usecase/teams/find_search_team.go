package teams

import (
	"context"
	"sync"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) findAllSearchTeam(wg *sync.WaitGroup, ctx context.Context, input *entities.TeamFilter, teams chan []entities.TeamSearch, errors chan error) {
	defer wg.Done()
	result, err := u.TeamsRepo.FindAllSearchTeam(ctx, input)

	if err != nil {
		errors <- err
		return
	}
	teams <- result
}

func (u useCase) FindAllSearchTeam(ctx context.Context, input *entities.TeamFilter) (*int64, []entities.TeamSearch, error) {
	count := make(chan *int64, 1)
	teams := make(chan []entities.TeamSearch, 1)
	errors := make(chan error, 2)
	var wg sync.WaitGroup
	wg.Add(1)
	go u.countTeam(&wg, ctx, input, count, errors)
	wg.Add(1)
	go u.findAllSearchTeam(&wg, ctx, input, teams, errors)

	go func() {
		wg.Wait()
		close(count)
		close(teams)
		close(errors)
	}()

	for err := range errors {
		return nil, nil, err
	}

	totalRecords, result, err := <-count, <-teams, <-errors
	return totalRecords, result, err
}
