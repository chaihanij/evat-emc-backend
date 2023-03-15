package fieldraceteams

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/layers/repositories/field_race_teams"
	"gitlab.com/chaihanij/evat/app/layers/repositories/files"
	"gitlab.com/chaihanij/evat/app/layers/repositories/members"
	"gitlab.com/chaihanij/evat/app/layers/repositories/users"
)

type useCase struct {
	Field_race_teamsRepo field_race_teams.Repo
	UsersRepo            users.Repo
	MembersRepo          members.Repo
	FilesRepo            files.Repo
	// AssignmentTeamsRepo assignmentTeams.Repo
}

type UseCase interface {
	FindAllFieldraceteam(ctx context.Context, input *entities.FieldRaceTeamFilter) (*int64, []entities.FieldRaceTeam, error)
}

func InitUseCase(field_race_teamsRepo field_race_teams.Repo, usersRepo users.Repo, membersRepo members.Repo, filesRepo files.Repo) UseCase {
	return &useCase{
		Field_race_teamsRepo: field_race_teamsRepo,
		UsersRepo:            usersRepo,
		MembersRepo:          membersRepo,
		FilesRepo:            filesRepo,
		// AssignmentTeamsRepo: assignmentTeamsRepo,
	}
}
