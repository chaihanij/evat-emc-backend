package assignments

import (
	"context"
	"sync"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) countAssignment(wg *sync.WaitGroup, ctx context.Context, input *entities.AssignmentFilter, count chan *int64, errors chan error) {
	defer wg.Done()
	result, err := u.AssignmentsRepo.CountAssignments(ctx, input)
	if err != nil {
		errors <- err
		return
	}
	count <- result
}

func (u useCase) findAllAssignment(wg *sync.WaitGroup, ctx context.Context, input *entities.AssignmentFilter, teams chan []entities.Assignment, errors chan error) {
	defer wg.Done()
	result, err := u.AssignmentsRepo.FindAllAssignment(ctx, input)

	if err != nil {
		errors <- err
		return
	}
	teams <- result
}

func (u useCase) FindAllAssignment(ctx context.Context, input *entities.AssignmentFilter) (*int64, []entities.Assignment, error) {
	count := make(chan *int64, 1)
	assignments := make(chan []entities.Assignment, 1)
	errors := make(chan error, 2)
	var wg sync.WaitGroup
	wg.Add(1)
	go u.countAssignment(&wg, ctx, input, count, errors)
	wg.Add(1)
	go u.findAllAssignment(&wg, ctx, input, assignments, errors)

	go func() {
		wg.Wait()
		close(count)
		close(assignments)
		close(errors)
	}()

	for err := range errors {
		return nil, nil, err
	}

	totalRecords, result, err := <-count, <-assignments, <-errors
	return totalRecords, result, err
}
