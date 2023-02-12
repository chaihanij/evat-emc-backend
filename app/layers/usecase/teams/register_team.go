package teams

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) RegisterTeam(ctx context.Context, team *entities.Team, user *entities.User) (*entities.Team, *entities.User, error) {
	team, err := u.TeamsRepo.CreateTeam(ctx, team)
	if err != nil {
		return nil, nil, err
	}

	user.TeamUUID = team.UUID
	user, err = u.UsersRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, nil, err
	}
	return team, user, nil
}
