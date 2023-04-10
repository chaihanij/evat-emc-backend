package assignments

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) UploadScoreAssignment(ctx context.Context, input *entities.Assignment) (*entities.Assignment, error) {
	if input.UUID == "" {
		return nil, fmt.Errorf("fail")
	}
	assignment, err := u.AssignmentsRepo.UploadScoreAssignment(ctx, input)
	if err != nil {
		return nil, err
	}
	log.WithField("value", assignment).Debugln("UseCase FindOneAssignments")

	return assignment, nil

}
