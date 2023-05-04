package teams

import (
	"context"
	"fmt"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) DeleteTeam(ctx context.Context, input *entities.TeamFilter) error {

	teamDelete, _ := u.TeamsRepo.FindOneTeam(ctx, input)

	logsetting := entities.LogSetting{
		NewData:     teamDelete,
		UUID_User:   *input.User_UUID,
		Discription: "Delete Team",
	}
	_, err := u.LogsettingRepo.CreateLogSetting(ctx, &logsetting)
	if err != nil {
		fmt.Println("err :", err)
	}

	return u.TeamsRepo.DeleteOneTeam(ctx, input)
}
