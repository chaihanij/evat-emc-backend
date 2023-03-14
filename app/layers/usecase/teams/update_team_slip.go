package teams

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) UpdateTeamSlip(ctx context.Context, teamUUID string, file *entities.File) (*entities.Team, error) {
	file, err := u.FilesRepo.CreateFile(ctx, file)
	if err != nil {
		return nil, err
	}
	team, err := u.TeamsRepo.PartialUpdateTeam(ctx, &entities.TeamPartialUpdate{UUID: teamUUID, Slip: &file.UUID})
	if err != nil {
		return nil, err
	}
	return u.FindOneTeam(ctx, &entities.TeamFilter{UUID: &team.UUID})
}
