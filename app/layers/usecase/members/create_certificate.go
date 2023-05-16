package members

import (
	"context"

	"gitlab.com/chaihanij/evat/app/entities"
)

func (u useCase) CreateCertificate(ctx context.Context, member_uuid string) (*entities.Member, error) {

	member, err := u.MembersRepo.FindOneMember(ctx, &entities.MemberFilter{UUID: &member_uuid})
	if err != nil {
		return nil, err
	}
	// config, err := u.ConfigRepo.FindOneConfig(ctx, &entities.Config{UUID: "b8900428-c97d-48f0-aec0-2d47a1bf44b6"})
	// if err != nil {
	// 	return nil, err
	// }

	// if *member.Is_Check_image == true && *member.Is_check_data == true && *member.Is_checkin == true {
	// 	return member, nil
	// }

	return member, nil

}
