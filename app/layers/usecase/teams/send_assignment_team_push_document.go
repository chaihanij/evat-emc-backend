package teams

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) SendAssignmentTeamPushDocument(ctx context.Context, input *entities.AssignmentTeam, file *entities.File) (*entities.File, error) {
	file, err := u.FilesRepo.CreateFile(ctx, file)
	if err != nil {
		return nil, err
	}

	_, err = u.AssignmentTeamsRepo.UpdateAssignmentTeamPushDocument(ctx, input, []string{file.UUID})
	if err != nil {
		return nil, err
	}
	return file, nil
}
