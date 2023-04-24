package consideration

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) AllScoreConsiderationAssignmentMember(ctx context.Context, input *entities.AllScoreFilter) ([]entities.AllScore, error) {
	return u.ConsiderationRepo.AllScoreConsiderationAssignmentMember(ctx, *input)

}
