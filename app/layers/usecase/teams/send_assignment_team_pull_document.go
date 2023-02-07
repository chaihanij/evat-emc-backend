package teams

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) SendAssignmentTeamPullDocument(ctx context.Context, input *entities.AssignmentTeamPartialUpdate, documentUUID string) error {
	_, err := u.AssignmentTeamsRepo.UpdateAssignmentTeamPullDocument(ctx, input, documentUUID)
	return err
}
