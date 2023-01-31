package members

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) DeleteMember(ctx context.Context, input *entities.MemberFilter) error {
	return u.MembersRepo.DeleteOneMemeber(ctx, input)
}
