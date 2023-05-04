package assignments

import (
	"context"

	"github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) CreateAssignment(ctx context.Context, input *entities.Assignment) (*entities.Assignment, error) {

	log := entities.LogSetting{
		NewData:   input,
		UUID_User: input.CreatedBy,
		Discription: "Create assignment",
	}
	_, err := u.LogsettingRepo.CreateLogSetting(ctx, &log)
	if err != nil {
		logrus.Debugln("error keep create log setting")
	}

	return u.AssignmentsRepo.CreateAssignment(ctx, input)
}
