package teams

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) FindOneTeam(ctx context.Context, input *entities.TeamFilter) (*entities.Team, error) {
	return u.TeamsRepo.FindOneTeam(ctx, input)
}
