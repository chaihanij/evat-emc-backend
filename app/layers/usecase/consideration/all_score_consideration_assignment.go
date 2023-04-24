package consideration

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
)

// AllScoreConsiderationAssignment

func (u useCase) AllScoreConsiderationAssignment(ctx context.Context, input *entities.AllScoreFilter) ([]entities.AllScore, error) {
	return u.ConsiderationRepo.AllScoreConsiderationAssignment(ctx, *input)
}
