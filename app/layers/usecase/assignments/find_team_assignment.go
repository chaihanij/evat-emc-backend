package assignments

import (
	"context"
	"sync"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) countTeamAssignment(wg *sync.WaitGroup, ctx context.Context, input *entities.AssignmentFilter, count chan *int64, errors chan error) {
	defer wg.Done()

	result, err := u.AssignmentsRepo.CountTeamAssignment(ctx, input)
	if err != nil {
		errors <- err
		return
	}
	count <- result
}

func (u useCase) findAllTeamAssignment(wg *sync.WaitGroup, ctx context.Context, input *entities.AssignmentFilter, teams chan []entities.TeamAssignment, errors chan error) {
	defer wg.Done()
	result, err := u.AssignmentsRepo.FindTeamAssignment(ctx, input)

	if err != nil {
		errors <- err
		return
	}
	teams <- result
}

func (u useCase) FindAllTeamAssignment(ctx context.Context, input *entities.AssignmentFilter) (*int64, []entities.TeamAssignment, error) {
	count := make(chan *int64, 1)

	teamassignments := make(chan []entities.TeamAssignment, 1)
	errors := make(chan error, 2)
	var wg sync.WaitGroup
	wg.Add(1)
	go u.countTeamAssignment(&wg, ctx, input, count, errors)
	wg.Add(1)
	go u.findAllTeamAssignment(&wg, ctx, input, teamassignments, errors)

	go func() {
		wg.Wait()
		close(count)
		close(teamassignments)
		close(errors)
	}()

	for err := range errors {
		return nil, nil, err
	}

	totalRecords, result, err := <-count, <-teamassignments, <-errors
	return totalRecords, result, err
}
