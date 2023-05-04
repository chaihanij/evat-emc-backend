package assignments

import (
	"context"
	"fmt"
	"time"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) UploadFileScore(ctx context.Context, assignmentUUID string, file *entities.File) (*entities.Assignment, error) {
	_, err := u.FilesRepo.CreateFile(ctx, file)
	if err != nil {
		return nil, err
	}
	fileinput := entities.Fileassignment{
		FileUrl:    file.UUID,
		FileName:   file.FileName,
		CreateDate: time.Now(),
	}

	assignmentOld, _ := u.AssignmentsRepo.FindOneAssignment(ctx, &entities.AssignmentFilter{UUID: &assignmentUUID})

	assignment, err := u.AssignmentsRepo.UploadFileScore(ctx, &entities.AssignmentPartialUpdateScore{AssignmentUUID: assignmentUUID, UploadFile: fileinput})
	if err != nil {
		return nil, err
	}

	log := entities.LogSetting{
		NewData:     assignment,
		UUID_User:   file.UpdateBy,
		OldData:     assignmentOld,
		Discription: "update score assignment",
	}

	_, err = u.LogsettingRepo.CreateLogSetting(ctx, &log)
	if err != nil {
		fmt.Println("err :", err)
	}

	return assignment, nil

}
