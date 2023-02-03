package assignments

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) DeleteAssignment(ctx context.Context, input *entities.AssignmentFilter) error {
	return u.AssignmentsRepo.DeleteOneAssignment(ctx, input)
}
