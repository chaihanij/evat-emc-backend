package assignments

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/layers/repositories/assignments"
	"gitlab.com/chaihanij/evat/app/layers/repositories/files"
)

type useCase struct {
	AssignmentsRepo assignments.Repo
	FilesRepo       files.Repo
}

type UseCase interface {
	CreateAssignment(ctx context.Context, input *entities.Assignment) (*entities.Assignment, error)
	DeleteAssignment(ctx context.Context, input *entities.AssignmentFilter) error
	FindAllAssignment(ctx context.Context, input *entities.AssignmentFilter) (*int64, []entities.Assignment, error)
	FindOneAssignment(ctx context.Context, input *entities.AssignmentFilter) (*entities.Assignment, error)
	UpdateAssignment(ctx context.Context, input *entities.AssignmentPartialUpdate) (*entities.Assignment, error)

	UpdateAssignmentImage(ctx context.Context, assignmentUUID string, file *entities.File) (*entities.File, error)
	UpdateAssignmentDocument(ctx context.Context, assignmentUUID string, file *entities.File) (*entities.File, error)

	FindTopicAssignment(ctx context.Context, input *entities.AssignmentFilter) (*entities.ExportAssignmentTopic, error)
	FindAllTeamAssignment(ctx context.Context, input *entities.AssignmentFilter) (*int64, []entities.TeamAssignment, error)
	UploadScoreAssignment(ctx context.Context, input *entities.Assignment) (*entities.Assignment, error)
}

func InitUseCase(assignmentsRepo assignments.Repo, filesRepo files.Repo) UseCase {
	return &useCase{
		AssignmentsRepo: assignmentsRepo,
		FilesRepo:       filesRepo,
	}
}
