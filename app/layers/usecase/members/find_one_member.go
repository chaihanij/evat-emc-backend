package members

import (
	"context"

	"github.com/AlekSi/pointer"
	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/mongo"
)

func (u useCase) FindOneMember(ctx context.Context, input *entities.MemberFilter) (*entities.Member, error) {
	member, err := u.MembersRepo.FindOneMember(ctx, input)
	if err != nil {
		return nil, err
	}

	if val, ok := member.Image.(string); ok {
		log.WithField("value", val).Debugln("FindOneMember Image")
		image, err := u.FilesRepo.FindOneFile(ctx, &entities.FileFilter{UUID: pointer.ToString(val)})
		if err != nil && mongo.ErrNoDocuments != err {
			return nil, err
		}
		if image != nil {
			member.Image = *image
		}
	}
	if val, ok := member.Documents.([]string); ok {
		log.WithField("value", val).Debugln("FindOneMember documents")
		documents, err := u.FilesRepo.FindAllFile(ctx, &entities.FileFilter{UUIDs: val})
		if err != nil && mongo.ErrNoDocuments != err {
			return nil, err
		}
		member.Documents = documents
	}
	log.WithField("member", member).Debugln("FindOneMember")
	return member, nil
}
