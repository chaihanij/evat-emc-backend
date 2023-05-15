package teams

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) ExportTeamMember(ctx context.Context, inputteam *entities.TeamFilter, inputmember *entities.MemberFilter) ([]entities.Team, []entities.Member, error ) {
	teams, _ := u.TeamsRepo.FindAllTeam(ctx, inputteam)

	members, _ := u.MembersRepo.FindAllMember(ctx, inputmember)

	return teams, members, nil

}
