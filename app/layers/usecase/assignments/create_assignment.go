package assignments

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) CreateAssignment(ctx context.Context, input *entities.Assignment) (*entities.Assignment, error) {
	return u.AssignmentsRepo.CreateAssignment(ctx, input)
}
