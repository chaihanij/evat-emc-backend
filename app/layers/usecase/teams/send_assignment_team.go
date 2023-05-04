package teams

import (
	"context"
	"fmt"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) SendAssignmentTeam(ctx context.Context, input *entities.AssignmentTeamPartialUpdate) (*entities.AssignmentTeam, error) {

	sendassignmentteam, err := u.AssignmentTeamsRepo.PartialUpdateAssignmentTeam(ctx, input)
	if err != nil {
		return nil, err
	}

	logsetting := entities.LogSetting{
		NewData:     sendassignmentteam,
		UUID_User:   input.UpdatedBy,
		Discription: "Send assignment Team",
	}
	_, err = u.LogsettingRepo.CreateLogSetting(ctx, &logsetting)
	if err != nil {
		fmt.Println("err :", err)
	}

	return sendassignmentteam, nil
}
