package teams

import (
	"context"

	// "fmt"

	// "github.com/sirupsen/logrus"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) RegisterTeam(ctx context.Context, team *entities.Team, user *entities.User, member *entities.Member) (*entities.Team, *entities.User, *entities.OmiseCharge, error) {
	team, err := u.TeamsRepo.CreateTeam(ctx, team)
	if err != nil {
		return nil, nil, nil, err
	}
	user.TeamUUID = team.UUID
	user, err = u.UsersRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, nil, nil, err
	}

	member.TeamUUID = team.UUID
	member.IsTeamLeader = true
	member, err = u.MembersRepo.CreateMember(ctx, member)
	if err != nil {
		return nil, nil, nil, err
	}

	var amount int64 = int64(200000)
	sourceID, err := u.OmiseRepo.CreateSource(amount)
	if err != nil {
		return nil, nil, nil, err
	}

	var metadata map[string]interface{} = map[string]interface{}{
		"team_uuid": team.UUID,
		"email":     user.Email,
		"tel":       user.Tel,
	}
	charge, err := u.OmiseRepo.CreateCharge(amount, *sourceID, metadata)
	if err != nil {
		return nil, nil, nil, err
	}

	// email := user.Email

	// err = u.TeamsRepo.SendEmailRegister(email)
	// if err != nil {
	// 	return nil, nil, nil, err
	// }

	return team, user, charge, nil
}
