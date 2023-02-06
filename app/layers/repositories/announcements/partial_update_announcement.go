package announcements

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/announcements/models"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r repo) PartialUpdateAnnouncement(ctx context.Context, input *entities.AnnouncementPartialUpdate) (*entities.Announcement, error) {
	log.Debugln("DB PartialUpdateMember")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	filter := models.NewAnnouncementFilter(input)
	update := models.PartialUpdateAnnouncement(input)
	var announcement models.Announcement
	err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionAnnouncements).
		FindOneAndUpdate(ctx, filter, update, opts).
		Decode(&announcement)
	if err != nil {
		log.WithError(err).Errorln("DB PartialUpdateAnnouncement Error")
		return nil, err
	}
	log.WithField("value", announcement).Debugln("DB PartialUpdateAnnouncement")
	return announcement.ToEntity()
}
