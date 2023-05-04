package assignments

import (
	"context"
	"fmt"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) DeleteAssignment(ctx context.Context, input *entities.AssignmentFilter) error {

	assignment, _ := u.AssignmentsRepo.FindOneAssignment(ctx, input)

	log := entities.LogSetting{
		NewData:     assignment,
		UUID_User:   *input.User_UUID,
		Discription: "delete assignment",
	}

	_, err := u.LogsettingRepo.CreateLogSetting(ctx, &log)
	if err != nil {
		fmt.Println("err :", err)
	}

	return u.AssignmentsRepo.DeleteOneAssignment(ctx, input)
}
