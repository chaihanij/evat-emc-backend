package members

import (
	"context"
	"fmt"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) DeleteMember(ctx context.Context, input *entities.MemberFilter) error {

	member, err := u.MembersRepo.FindOneMember(ctx, input)
	if err != nil {
		// return nil, err
	}

	logsetting := entities.LogSetting{
		NewData:     member,
		UUID_User:   *input.User_UUID,
		Discription: "Delete Member",
	}

	_, err = u.LogsettingRepo.CreateLogSetting(ctx, &logsetting)
	if err != nil {
		fmt.Println("err :", err)
	}

	return u.MembersRepo.DeleteOneMemeber(ctx, input)
}
