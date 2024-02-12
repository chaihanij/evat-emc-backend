package members

import (
	"context"
	"fmt"

	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

func (u useCase) CreateMember(ctx context.Context, input *entities.Member) (*entities.Member, error) {
	count, _, err := u.FindAllMember(ctx, &entities.MemberFilter{TeamUUID: &input.TeamUUID})
	if err != nil {
		return nil, err
	}
	//OAK Test 7 member in team
	if *count > 7 {
		return nil, errors.ParameterError{Message: "Maximum Member in Team"}
	}

	logsetting := entities.LogSetting{
		NewData:     input,
		UUID_User:   input.CreatedBy,
		Discription: "Create Member",
	}

	_, err = u.LogsettingRepo.CreateLogSetting(ctx, &logsetting)
	if err != nil {
		fmt.Println("err :", err)
	}

	return u.MembersRepo.CreateMember(ctx, input)
}
