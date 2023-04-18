package assignments

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) FindTopicAssignment(ctx context.Context, input *entities.AssignmentFilter) (*entities.ExportAssignmentTopic , error){
	topic, err := u.AssignmentsRepo.FindTopicAssignment(ctx, input)
	if err != nil {
		return nil, err
	}
	return topic, err
}