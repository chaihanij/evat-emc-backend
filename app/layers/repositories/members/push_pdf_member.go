package members

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/members/models"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r repo) UploadPDFMember(ctx context.Context, input *entities.MemberUpdatePDF) (*entities.Member, error) {

	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	filter := models.UpdateFileMemberFilter(*&input.MemberUUID)
	update := models.UploadFileMember(input)

	var member models.Member

	err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionMembers).
		FindOneAndUpdate(ctx, filter, update, opts).
		Decode(&member)
	if err != nil {
		log.WithError(err).Errorln("DB UpdatePDFMember Error")
		return nil, err
	}

	return member.ToEntity()

}
