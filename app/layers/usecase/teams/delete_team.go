package teams

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) DeleteTeam(ctx context.Context, input *entities.TeamFilter) error {
	return u.TeamsRepo.DeleteOneTeam(ctx, input)
}
