package members

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

func (u useCase) CreateMember(ctx context.Context, input *entities.Member) (*entities.Member, error) {
	count, _, err := u.FindAllMember(ctx, &entities.MemberFilter{TeamUUID: &input.TeamUUID})
	if err != nil {
		return nil, err
	}
	if *count > 6 {
		return nil, errors.ParameterError{Message: "Maximum Member in Team"}
	}
	return u.MembersRepo.CreateMember(ctx, input)
}
