package teams

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) UpdateTeam(ctx context.Context, input *entities.TeamPartialUpdate) (*entities.Team, error) {
	return u.TeamsRepo.PartialUpdateTeam(ctx, input)
}
