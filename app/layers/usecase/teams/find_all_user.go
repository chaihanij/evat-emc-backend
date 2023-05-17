package teams

import (
	"context"

	// "fmt"
	"sync"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) countTeam(wg *sync.WaitGroup, ctx context.Context, input *entities.TeamFilter, count chan *int64, errors chan error) {
	defer wg.Done()
	result, err := u.TeamsRepo.CountTeam(ctx, input)
	if err != nil {
		errors <- err
		return
	}
	count <- result
}

func (u useCase) findAllTeam(wg *sync.WaitGroup, ctx context.Context, input *entities.TeamFilter, teams chan []entities.Team, errors chan error) {
	defer wg.Done()

	// var result []entities.TeamFilter

	result, err := u.TeamsRepo.FindAllTeam(ctx, input)

	// if input.Name != nil {
	if *input.Name != "" {
		resultmember, _ := u.UsersRepo.FindAllUser(ctx, &entities.UserFilter{Tel: input.Name})
		if resultmember != nil {
			for _, v := range resultmember {
				resultm, _ := u.TeamsRepo.FindAllTeam(ctx, &entities.TeamFilter{UUID: &v.TeamUUID})
				result = append(result, resultm...)
			}

		}
	}

	// resultmember, _ := u.UsersRepo.FindAllUser(ctx, &entities.UserFilter{Tel: input.Name})
	// if resultmember != nil {
	// 	for _, v := range resultmember {
	// 		resultm, _ := u.TeamsRepo.FindAllTeam(ctx, &entities.TeamFilter{UUID: &v.TeamUUID})
	// 		result = append(result, resultm...)
	// 	}

	// }

	// }

	// for _, v := range result {
	// 	fmt.Println("v :", v.Name)
	// }

	if err != nil {
		errors <- err
		return
	}
	teams <- result
}

func (u useCase) FindAllTeam(ctx context.Context, input *entities.TeamFilter) (*int64, []entities.Team, error) {
	count := make(chan *int64, 1)
	teams := make(chan []entities.Team, 1)
	errors := make(chan error, 2)
	var wg sync.WaitGroup
	wg.Add(1)
	go u.countTeam(&wg, ctx, input, count, errors)
	wg.Add(1)
	go u.findAllTeam(&wg, ctx, input, teams, errors)

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
