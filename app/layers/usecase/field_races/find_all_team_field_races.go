package fieldraces

import (
	"context"
	"sync"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) countTeamFieldRacs(wg *sync.WaitGroup, ctx context.Context, input *entities.FieldRaceFilter, count chan *int64, errors chan error) {
	defer wg.Done()
	result, err := u.TeamFieldRacesRepo.CountFieldRaces(ctx, input)
	if err != nil {
		errors <- err
		return
	}
	count <- result
}

func (u useCase) findAllTeamFieldRacs(wg *sync.WaitGroup, ctx context.Context, input *entities.FieldRaceFilter, allteamfieldraces chan []entities.FieldRace, errors chan error) {
	defer wg.Done()
	result, err := u.TeamFieldRacesRepo.FindTeamAllFieldRace(ctx, input)

	if err != nil {
		errors <- err
		return
	}
	allteamfieldraces <- result
}

func (u useCase) FindAllTeamFieldRace(ctx context.Context, input *entities.FieldRaceFilter) (*int64, []entities.FieldRace, error) {
	count := make(chan *int64, 1)
	teamfieldRace := make(chan []entities.FieldRace, 1)
	errors := make(chan error, 2)
	var wg sync.WaitGroup
	wg.Add(1)
	go u.countTeamFieldRacs(&wg, ctx, input, count, errors)
	wg.Add(1)
	go u.findAllTeamFieldRacs(&wg, ctx, input, teamfieldRace, errors)

	go func() {
		wg.Wait()
		close(count)
		close(teamfieldRace)
		close(errors)
	}()

	for err := range errors {
		return nil, nil, err
	}

	totalRecords, result, err := <-count, <-teamfieldRace, <-errors
	return totalRecords, result, err
}
