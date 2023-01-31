package members

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) CreateMember(ctx context.Context, input *entities.Member) (*entities.Member, error) {
	return u.MembersRepo.CreateMember(ctx, input)
}
