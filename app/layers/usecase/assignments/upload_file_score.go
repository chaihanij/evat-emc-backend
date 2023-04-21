package assignments

import (
	"context"
	"time"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) UploadFileScore(ctx context.Context, assignmentUUID string, file *entities.File) (*entities.Assignment, error) {

	file, err := u.FilesRepo.CreateFile(ctx, file)
	if err != nil {
		return nil, err
	}
	fileinput := entities.UploadFile{
		FileName:   file.UUID,
		CreateDate: time.Now(),
	}

	assignment, err := u.AssignmentsRepo.UploadFileScore(ctx, &entities.Assignment{UUID: assignmentUUID, UploadFile: fileinput})
	if err != nil {
		return nil, err
	}

	return assignment, nil

}
