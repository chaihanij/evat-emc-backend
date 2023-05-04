package members

import (
	"context"
	"fmt"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) UpdateMemberPushDocument(ctx context.Context, memberUUID string, file *entities.File) (*entities.File, error) {
	file, err := u.FilesRepo.CreateFile(ctx, file)
	if err != nil {
		return nil, err
	}
	updatePushMember, err := u.MembersRepo.PushDocument(ctx, memberUUID, file.UUID)
	if err != nil {
		return nil, err
	}

	logsetting := entities.LogSetting{
		NewData:     updatePushMember,
		UUID_User:   memberUUID,
		Discription: "UpdateMemberPushDocument",
	}
	_, err = u.LogsettingRepo.CreateLogSetting(ctx, &logsetting)
	if err != nil {
		fmt.Println("err :", err)
	}

	return file, nil
}
