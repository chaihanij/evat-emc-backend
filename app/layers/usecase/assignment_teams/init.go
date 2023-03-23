package assignmentteams

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/layers/repositories/assignment_teams"
)

type useCase struct {
	AssignmentsTeamRepo assignment_teams.Repo
}

// assignment_teams
type UseCase interface {
	FindAllAssignmentTeamscore(ctx context.Context, input *entities.AssignmentTeamFilter) (*int64, []entities.AssignmentTeamScore, error)
}

func InitUseCase(assignmentsTeamRepo assignment_teams.Repo) UseCase {
	return &useCase{
		AssignmentsTeamRepo: assignmentsTeamRepo,
	}
}
