package fieldraceteams

import (
	"context"
	"sync"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) FindFieldracesteam(ctx context.Context, input *entities.FieldRaceTeamFilter) (*int64, []entities.FieldRaceTeam, error) {
	count := make(chan *int64, 1)
	fieldraceteams := make(chan []entities.FieldRaceTeam, 1)
	errors := make(chan error, 2)
	var wg sync.WaitGroup
	wg.Add(1)
	go u.countFieldraceteam(&wg, ctx, input, count, errors)
	wg.Add(1)
	go u.findAllFieldraceteam(&wg, ctx, input, fieldraceteams, errors)

	go func() {
		wg.Wait()
		close(count)
		close(fieldraceteams)
		close(errors)
	}()

	for err := range errors {
		return nil, nil, err
	}

	totalRecords, result, err := <-count, <-fieldraceteams, <-errors
	return totalRecords, result, err
}
