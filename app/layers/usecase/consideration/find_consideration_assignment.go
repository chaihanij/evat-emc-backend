package consideration

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) FindOneConsideration(ctx context.Context, input *entities.ConsiderationFilter) ([]entities.AssignmentScore, error) {
	return u.ConsiderationRepo.FindOneConsideration(ctx, input)
}
