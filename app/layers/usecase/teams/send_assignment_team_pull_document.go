package teams

import (
	"context"
	"fmt"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) SendAssignmentTeamPullDocument(ctx context.Context, input *entities.AssignmentTeamPartialUpdate, documentUUID string) error {
	deleteDocTeam, err := u.AssignmentTeamsRepo.UpdateAssignmentTeamPullDocument(ctx, input, documentUUID)

	logsetting := entities.LogSetting{
		NewData:     deleteDocTeam,
		UUID_User:   input.UpdatedBy,
		Discription: "Delete Assignment Team",
	}
	_, err = u.LogsettingRepo.CreateLogSetting(ctx, &logsetting)
	if err != nil {
		fmt.Println("err :", err)
	}

	return err
}
