package members

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) FindOneMember(ctx context.Context, input *entities.MemberFilter) (*entities.Member, error) {
	return u.MembersRepo.FindOneMember(ctx, input)
}
