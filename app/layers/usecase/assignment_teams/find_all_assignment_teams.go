package assignmentteams

import (
	"context"
	"sync"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) countAssignmentTeamscore(wg *sync.WaitGroup, ctx context.Context, input *entities.AssignmentTeamFilter, count chan *int64, errors chan error) {
	defer wg.Done()
	result, err := u.AssignmentsTeamRepo.CountAssignmentTeam(ctx, input)
	if err != nil {
		errors <- err
		return
	}
	count <- result
}

func (u useCase) findAllAssignmentTeamscore(wg *sync.WaitGroup, ctx context.Context, input *entities.AssignmentTeamFilter, Assignmentteams chan []entities.AssignmentTeamScore, errors chan error) {
	defer wg.Done()
	result, err := u.AssignmentsTeamRepo.FindAllscoreAssignmentTeam(ctx, input)

	if err != nil {
		errors <- err
		return
	}
	Assignmentteams <- result
}

func (u useCase) FindAllAssignmentTeamscore(ctx context.Context, input *entities.AssignmentTeamFilter) (*int64, []entities.AssignmentTeamScore, error) {
	count := make(chan *int64, 1)
	assignmentteamscores := make(chan []entities.AssignmentTeamScore, 1)
	errors := make(chan error, 2)
	var wg sync.WaitGroup
	wg.Add(1)
	go u.countAssignmentTeamscore(&wg, ctx, input, count, errors)
	wg.Add(1)
	go u.findAllAssignmentTeamscore(&wg, ctx, input, assignmentteamscores, errors)

	go func() {
		wg.Wait()
		close(count)
		close(assignmentteamscores)
		close(errors)
	}()

	for err := range errors {
		return nil, nil, err
	}

	totalRecords, result, err := <-count, <-assignmentteamscores, <-errors
	return totalRecords, result, err
}
