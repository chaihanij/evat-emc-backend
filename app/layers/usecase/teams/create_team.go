package teams

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) CreateTeam(ctx context.Context, input *entities.Team) (*entities.Team, error) {
	return u.TeamsRepo.CreateTeam(ctx, input)
}
