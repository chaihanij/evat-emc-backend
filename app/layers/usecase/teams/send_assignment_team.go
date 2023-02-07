package teams

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) SendAssignmentTeam(ctx context.Context, input *entities.AssignmentTeamPartialUpdate) (*entities.AssignmentTeam, error) {
	return u.AssignmentTeamsRepo.PartialUpdateAssignmentTeam(ctx, input)
}
