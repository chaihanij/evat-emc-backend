package fieldraces

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/layers/repositories/field_races"
)

type useCase struct {
	TeamFieldRacesRepo field_races.Repo
}

type UseCase interface {
	FindAllTeamFieldRace(ctx context.Context, input *entities.FieldRaceFilter) (*int64, []entities.FieldRace, error)
}

func InitUseCase(teamfieldracesrepoRepo field_races.Repo) UseCase {
	return &useCase{
		TeamFieldRacesRepo: teamfieldracesrepoRepo,
	}
}
