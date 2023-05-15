package teams

import (
	"context"

	"github.com/AlekSi/pointer"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/mongo"
)

func (u useCase) FindOneTeam(ctx context.Context, input *entities.TeamFilter) (*entities.Team, error) {
	members, err := u.MembersRepo.FindAllMember(ctx, &entities.MemberFilter{TeamUUID: input.UUID})
	if err != nil && mongo.ErrNoDocuments != err {
		return nil, err
	}
	for index, member := range members {
		if val, ok := member.Image.(string); ok {
			// log.WithField("value", val).Debugln("FindOneTeam Member Image")
			image, err := u.FilesRepo.FindOneFile(ctx, &entities.FileFilter{UUID: pointer.ToString(val)})
			if err != nil && mongo.ErrNoDocuments != err {
				return nil, err
			}
			if image != nil {
				member.Image = *image
			}
		}
		if val, ok := member.Documents.([]string); ok {
			// log.WithField("value", val).Debugln("FindOneTeam Member Documents")
			documents, err := u.FilesRepo.FindAllFile(ctx, &entities.FileFilter{UUIDs: val})
			if err != nil && mongo.ErrNoDocuments != err {
				return nil, err
			}
			if len(documents) > 0 {
				member.Documents = documents
			} else {
				member.Documents = []entities.File{}
			}
		}
		// log.WithField("member", member).Debugln("FindOneTeam Member")

		members[index] = member
	}

	team, err := u.TeamsRepo.FindOneTeam(ctx, input)
	if err != nil {
		return nil, err
	}
	team.Members = members
	if val, ok := team.Slip.(string); ok {
		silp, err := u.FilesRepo.FindOneFile(ctx, &entities.FileFilter{UUID: pointer.ToString(val)})
		if err != nil && mongo.ErrNoDocuments != err {
			return nil, err
		}
		if silp != nil {
			team.Slip = *silp
		}
	}
	// log.WithField("team", team).Debugln("FindOneTeam")

	return team, nil
}
