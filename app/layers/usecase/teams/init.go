package teams

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
	assignmentTeams "gitlab.com/chaihanij/evat/app/layers/repositories/assignment_teams"
	"gitlab.com/chaihanij/evat/app/layers/repositories/files"
	"gitlab.com/chaihanij/evat/app/layers/repositories/members"
	"gitlab.com/chaihanij/evat/app/layers/repositories/omise"
	"gitlab.com/chaihanij/evat/app/layers/repositories/teams"
	"gitlab.com/chaihanij/evat/app/layers/repositories/users"
)

type useCase struct {
	TeamsRepo           teams.Repo
	UsersRepo           users.Repo
	MembersRepo         members.Repo
	FilesRepo           files.Repo
	AssignmentTeamsRepo assignmentTeams.Repo
	OmiseRepo           omise.Repo
}

type UseCase interface {
	CreateTeam(ctx context.Context, input *entities.Team) (*entities.Team, error)
	DeleteTeam(ctx context.Context, input *entities.TeamFilter) error
	FindAllTeam(ctx context.Context, input *entities.TeamFilter) (*int64, []entities.Team, error)
	FindOneTeam(ctx context.Context, input *entities.TeamFilter) (*entities.Team, error)
	UpdateTeam(ctx context.Context, input *entities.TeamPartialUpdate) (*entities.Team, error)
	UpdateTeamSlip(ctx context.Context, teamUUID string, file *entities.File) (*entities.Team, error)
	FindOneAssignmentTeam(ctx context.Context, input *entities.AssignmentTeamFilter) (*entities.AssignmentTeam, error)
	SendAssignmentTeam(ctx context.Context, input *entities.AssignmentTeamPartialUpdate) (*entities.AssignmentTeam, error)
	SendAssignmentTeamPushDocument(ctx context.Context, input *entities.AssignmentTeamPartialUpdate, file *entities.File) (*entities.File, error)
	SendAssignmentTeamPullDocument(ctx context.Context, input *entities.AssignmentTeamPartialUpdate, documentUUID string) error

	FindAllSearchTeam(ctx context.Context, input *entities.TeamFilter) (*int64, []entities.TeamSearch, error)

	RegisterTeam(ctx context.Context, team *entities.Team, user *entities.User) (*entities.Team, *entities.User, *entities.OmiseCharge, error)
}

func InitUseCase(teamsRepo teams.Repo,
	usersRepo users.Repo, membersRepo members.Repo,
	filesRepo files.Repo,
	assignmentTeamsRepo assignmentTeams.Repo,
	omiseRepo omise.Repo) UseCase {
	return &useCase{
		TeamsRepo:           teamsRepo,
		UsersRepo:           usersRepo,
		MembersRepo:         membersRepo,
		FilesRepo:           filesRepo,
		AssignmentTeamsRepo: assignmentTeamsRepo,
		OmiseRepo:           omiseRepo,
	}
}
