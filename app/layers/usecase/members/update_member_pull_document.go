package members

import (
	"context"
	"fmt"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) UpdateMemberPullDocument(ctx context.Context, memberUUID string, fileUUID string) error {

	updatememberDoc, err := u.MembersRepo.PullDocument(ctx, memberUUID, fileUUID)
	if err != nil {
		return err
	}

	logsetting := entities.LogSetting{
		NewData:     updatememberDoc,
		UUID_User:   memberUUID,
		Discription: "UpdateMemberPullDocument",
	}
	_, err = u.LogsettingRepo.CreateLogSetting(ctx, &logsetting)
	if err != nil {
		fmt.Println("err :", err)
	}

	return nil
}
