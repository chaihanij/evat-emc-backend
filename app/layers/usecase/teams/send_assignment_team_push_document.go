package teams

import (
	"context"
	"fmt"

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

	senddocTeam, err := u.AssignmentTeamsRepo.UpdateAssignmentTeamPushDocument(ctx, input, file.UUID)
	if err != nil {
		return nil, err
	}

	logsetting := entities.LogSetting{
		NewData:     senddocTeam,
		UUID_User:   input.UpdatedBy,
		Discription: "send document Team",
	}
	_, err = u.LogsettingRepo.CreateLogSetting(ctx, &logsetting)
	if err != nil {
		fmt.Println("err :", err)
	}

	return file, nil
}
