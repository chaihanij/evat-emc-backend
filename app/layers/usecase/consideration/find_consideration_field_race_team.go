package consideration

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) FindConsiderationFieldRaceTeam(ctx context.Context, input *entities.ConsiderationFilter) ([]entities.FieldRaceTeamScore, error) {
	return u.ConsiderationRepo.FindConsiderationFieldRaceTeam(ctx, input)
}
