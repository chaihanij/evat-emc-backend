package assignments

import (
	"context"

	"github.com/AlekSi/pointer"
	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) UpdateAssignmentImage(ctx context.Context, assignmentUUID string, file *entities.File) (*entities.File, error) {
	log.WithField("assignment_uuid", assignmentUUID).WithField("file", file).Debugln("UpdateAssignmentImage")

	file, err := u.FilesRepo.CreateFile(ctx, file)
	if err != nil {
		return nil, err
	}

	_, err = u.AssignmentsRepo.PartialUpdateAssignment(ctx, &entities.AssignmentPartialUpdate{UUID: assignmentUUID, Image: pointer.ToString(file.UUID)})
	if err != nil {
		return nil, err
	}
	return file, nil
}
