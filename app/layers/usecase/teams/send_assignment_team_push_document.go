package teams

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) SendAssignmentTeamPushDocument(ctx context.Context, input *entities.AssignmentTeamPartialUpdate, file *entities.File) (*entities.File, error) {
	file, err := u.FilesRepo.CreateFile(ctx, file)
	if err != nil {
		return nil, err
	}

	// assignmentTeam, err := u.AssignmentTeamsRepo.FindOneAssignmentTeam(ctx, &entities.AssignmentTeamFilter{TeamUUID: input.TeamUUID, AssignmentUUID: input.AssignmentUUID})
	// if err != nil && err != mongo.ErrNoDocuments {
	// 	return nil, err
	// }
	// logrus.WithField("value", assignmentTeam).Debugln("SendAssignmentTeamPushDocument")

	// if value, ok := assignmentTeam.Documents.([]string); ok {
	// 	if len(value) <= 0 {
	// 		assignmentTeam.Documents = []string{}
	// 		u.AssignmentTeamsRepo.UpdateAssignmentTeam(ctx, assignmentTeam)
	// 	}
	// } else {
	// 	assignmentTeam.Documents = []string{}
	// 	u.AssignmentTeamsRepo.UpdateAssignmentTeam(ctx, assignmentTeam)
	// }

	_, err = u.AssignmentTeamsRepo.UpdateAssignmentTeamPushDocument(ctx, input, file.UUID)
	if err != nil {
		return nil, err
	}
	return file, nil
}
