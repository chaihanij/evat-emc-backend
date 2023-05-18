package members

import (
	"context"
	"fmt"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) UploadPDFMember(ctx context.Context, memberUUID string, file *entities.File) (*entities.File, error) {
	file, err := u.FilesRepo.CreateFile(ctx, file)
	if err != nil {
		return nil, err
	}

	uploadPDFMember, err := u.MembersRepo.UploadPDFMember(ctx, &entities.MemberUpdatePDF{MemberUUID: memberUUID, Files: file.UUID})

	logsetting := entities.LogSetting{
		NewData:     uploadPDFMember,
		UUID_User:   memberUUID,
		Discription: "UpdatePDFMember",
	}
	_, err = u.LogsettingRepo.CreateLogSetting(ctx, &logsetting)
	if err != nil {
		fmt.Println("err :", err)
	}
	return file, nil

}
