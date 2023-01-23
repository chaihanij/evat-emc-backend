package users

import (
	"context"
	"sync"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) countUser(wg *sync.WaitGroup, ctx context.Context, input *entities.UserFilter, count chan *int64, errors chan error) {
	defer wg.Done()
	result, err := u.UsersRepo.CountUser(ctx, input)
	if err != nil {
		errors <- err
		return
	}
	count <- result
}

func (u useCase) findAllUser(wg *sync.WaitGroup, ctx context.Context, input *entities.UserFilter, users chan []entities.User, errors chan error) {
	defer wg.Done()
	result, err := u.UsersRepo.FindAllUser(ctx, input)

	if err != nil {
		errors <- err
		return
	}
	users <- result
}

func (u useCase) FindAllHubCredentials(ctx context.Context, input *entities.UserFilter) (*int64, []entities.User, error) {
	count := make(chan *int64, 1)
	users := make(chan []entities.User, 1)
	errors := make(chan error, 2)
	var wg sync.WaitGroup
	wg.Add(1)
	go u.countUser(&wg, ctx, input, count, errors)
	wg.Add(1)
	go u.findAllUser(&wg, ctx, input, users, errors)

	go func() {
		wg.Wait()
		close(count)
		close(users)
		close(errors)
	}()

	for err := range errors {
		return nil, nil, err
	}

	totalRecords, result, err := <-count, <-users, <-errors
	return totalRecords, result, err
}
