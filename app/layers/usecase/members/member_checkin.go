package members

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) UseCaseMemberCheckeIn(ctx context.Context, input *entities.MemberCheckIn) (*entities.Member, error) {
	memberCheckIn, err := u.MembersRepo.MemberCheckIn(ctx, input)
	if err != nil {

		log.Debugln("error", err)

		return nil, err
	}
	// fmt.Println("input :", input)
	return memberCheckIn, nil
}
