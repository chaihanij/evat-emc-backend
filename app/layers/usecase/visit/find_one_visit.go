package visit

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) FindVisit(ctx context.Context) (*entities.Visited, error) {
	visit, err := u.VisitRepo.FindOneVisited(ctx)
	if err != nil {
		return nil, err
	}
	return visit, nil
}
