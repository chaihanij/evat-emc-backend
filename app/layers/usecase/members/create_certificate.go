package members

import (
	"context"
	"encoding/json"
	"time"

	"gitlab.com/chaihanij/evat/app/entities"
)

type ConfigProject struct {
	Start_date time.Time
	End_date   time.Time
}

func (u useCase) CreateCertificate(ctx context.Context, member_uuid string) (*entities.Member, error) {

	member, err := u.MembersRepo.FindOneMember(ctx, &entities.MemberFilter{UUID: &member_uuid})
	if err != nil {
		return nil, err
	}

	config, err := u.ConfigRepo.FindOneConfig(ctx, &entities.Config{UUID: "b8900428-c97d-48f0-aec0-2d47a1bf44b6"})
	if err != nil {
		return nil, err
	}

	fltB, _ := json.Marshal(config.StartProject)
	res := ConfigProject{}
	json.Unmarshal([]byte(string(fltB)), &res)

	if time.Now().Local().Unix() <= res.End_date.Local().Unix() {
		return nil, nil
	}

	// fmt.Println("configEnd :", configEnd    )

	// if val, ok := config.StartProject.(entities.DateStartProject); ok {
	// 	fmt.Println("val :", val.End_date)
	// } else {
	// 	fmt.Println("Conversion failed")
	// }

	// if str, ok := config.StartProject.(string); ok {
	// 	var val entities.DateStartProject

	// 	// Parse the string into a DateStartProject struct
	// 	_, err := fmt.Sscanf(str, "{%s %s}", &val.Start_date, &val.End_date)
	// 	if err == nil {
	// 		fmt.Println("val:", val.End_date)
	// 	}
	// }

	// if *member.Is_Check_image == true && *member.Is_check_data == true && *member.Is_checkin == true {
	// 	return member, nil
	// }

	return member, nil

}
