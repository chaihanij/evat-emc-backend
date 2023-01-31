package members

import (
	"context"
	"sync"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) countMember(wg *sync.WaitGroup, ctx context.Context, input *entities.MemberFilter, count chan *int64, errors chan error) {
	defer wg.Done()
	result, err := u.MembersRepo.CountMember(ctx, input)
	if err != nil {
		errors <- err
		return
	}
	count <- result
}

func (u useCase) findAllMemeber(wg *sync.WaitGroup, ctx context.Context, input *entities.MemberFilter, teams chan []entities.Member, errors chan error) {
	defer wg.Done()
	result, err := u.MembersRepo.FindAllMember(ctx, input)

	if err != nil {
		errors <- err
		return
	}
	teams <- result
}

func (u useCase) FindAllMember(ctx context.Context, input *entities.MemberFilter) (*int64, []entities.Member, error) {
	count := make(chan *int64, 1)
	members := make(chan []entities.Member, 1)
	errors := make(chan error, 2)
	var wg sync.WaitGroup
	wg.Add(1)
	go u.countMember(&wg, ctx, input, count, errors)
	wg.Add(1)
	go u.findAllMemeber(&wg, ctx, input, members, errors)

	go func() {
		wg.Wait()
		close(count)
		close(members)
		close(errors)
	}()

	for err := range errors {
		return nil, nil, err
	}

	totalRecords, result, err := <-count, <-members, <-errors
	return totalRecords, result, err
}
