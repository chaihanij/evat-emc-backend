package members

import (
	"context"
	"fmt"

	"github.com/AlekSi/pointer"
	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/mongo"
)

func (u useCase) UpdateMember(ctx context.Context, input *entities.MemberPartialUpdate) (*entities.Member, error) {
	memberOld, err := u.MembersRepo.FindOneMember(ctx, &entities.MemberFilter{UUID: &input.UUID})
	if err != nil {
		fmt.Println("err 123 :", err)
		return nil, err
	}
	member, err := u.MembersRepo.PartialUpdateMember(ctx, input)
	if err != nil {
		return nil, err
	}

	if val, ok := member.Image.(string); ok {
		log.WithField("value", val).Debugln("UpdateMember Image")
		image, err := u.FilesRepo.FindOneFile(ctx, &entities.FileFilter{UUID: pointer.ToString(val)})
		if err != nil && mongo.ErrNoDocuments != err {
			return nil, err
		}
		if image != nil {
			member.Image = *image
		}
	}
	if val, ok := member.Documents.([]string); ok {
		log.WithField("value", val).Debugln("UpdateMember documents")
		documents, err := u.FilesRepo.FindAllFile(ctx, &entities.FileFilter{UUIDs: val})
		if err != nil && mongo.ErrNoDocuments != err {
			return nil, err
		}
		member.Documents = documents
	}

	logsetting := entities.LogSetting{
		OldData:     memberOld,
		NewData:     member,
		UUID_User:   *input.UpdatedBy,
		Discription: "Update Member",
	}

	_, err = u.LogsettingRepo.CreateLogSetting(ctx, &logsetting)
	if err != nil {
		fmt.Println("err :", err)
	}

	log.WithField("member", member).Debugln("UpdateMember")
	return member, nil
}
