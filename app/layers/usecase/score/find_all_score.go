package score

import (
	"context"

	// "fmt"
	"sync"

	// log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) countScore(wg *sync.WaitGroup, ctx context.Context, input *entities.ScoreFilter, count chan *int64, errors chan error) {
	defer wg.Done()
	result, err := u.ScoreRepo.CountScore(ctx, input)
	if err != nil {
		errors <- err
		return
	}
	count <- result
}

func (u useCase) findAllScore(wg *sync.WaitGroup, ctx context.Context, input *entities.ScoreFilter, scores chan []entities.Score, errors chan error) {
	defer wg.Done()
	result, err := u.ScoreRepo.FindAllScore(ctx, input)

	if err != nil {
		errors <- err
		return
	}
	scores <- result
}

func (u useCase) FindAllScore(ctx context.Context, input *entities.ScoreFilter) (*int64, []entities.Score, error) {
	count := make(chan *int64, 1)
	scores := make(chan []entities.Score, 1)
	errors := make(chan error, 2)
	var wg sync.WaitGroup
	wg.Add(1)
	go u.countScore(&wg, ctx, input, count, errors)
	wg.Add(1)
	go u.findAllScore(&wg, ctx, input, scores, errors)

	go func() {
		wg.Wait()
		close(count)
		close(scores)
		close(errors)
	}()

	for err := range errors {
		return nil, nil, err
	}

	totalRecords, result, err := <-count, <-scores, <-errors
	return totalRecords, result, err
}
