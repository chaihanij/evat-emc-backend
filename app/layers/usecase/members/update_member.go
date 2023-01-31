package members

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) UpdateMember(ctx context.Context, input *entities.MemberPartialUpdate) (*entities.Member, error) {
	return u.MembersRepo.PartialUpdateMember(ctx, input)
}
