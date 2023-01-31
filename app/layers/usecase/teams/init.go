package teams

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/layers/repositories/files"
	"gitlab.com/chaihanij/evat/app/layers/repositories/members"
	"gitlab.com/chaihanij/evat/app/layers/repositories/teams"
)

type useCase struct {
	TeamsRepo   teams.Repo
	MembersRepo members.Repo
	FilesRepo   files.Repo
}

type UseCase interface {
	CreateTeam(ctx context.Context, input *entities.Team) (*entities.Team, error)
	DeleteTeam(ctx context.Context, input *entities.TeamFilter) error
	FindAllTeam(ctx context.Context, input *entities.TeamFilter) (*int64, []entities.Team, error)
	FindOneTeam(ctx context.Context, input *entities.TeamFilter) (*entities.Team, error)
	UpdateTeam(ctx context.Context, input *entities.TeamPartialUpdate) (*entities.Team, error)

	//
	CreateMember(ctx context.Context, input *entities.Member) (*entities.Member, error)
}

func InitUseCase(teamsRepo teams.Repo, membersRepo members.Repo, filesRepo files.Repo) UseCase {
	return &useCase{
		TeamsRepo:   teamsRepo,
		MembersRepo: membersRepo,
		FilesRepo:   filesRepo,
	}
}
