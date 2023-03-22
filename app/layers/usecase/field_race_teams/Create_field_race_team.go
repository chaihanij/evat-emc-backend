package fieldraceteams

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) CreateFieldRaceTeam(ctx context.Context , input *entities.FieldRaceTeam) (*entities.FieldRaceTeam , error) {
	return u.Field_race_teamsRepo.CreateFieldRaceTeam(ctx, input)
}