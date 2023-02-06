package announcements

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/announcements/models"
)

func (r repo) FindOneAnnouncement(ctx context.Context, input *entities.AnnouncementFilter) (*entities.Announcement, error) {
	log.Debugln("DB FindOneAnnouncement")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	filter := models.NewAnnouncementFilter(input)
	var announcement models.Announcement
	err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionAnnouncements).
		FindOne(ctx, filter, nil).
		Decode(&announcement)
	if err != nil {
		log.WithError(err).Errorln("DB FindOneAnnouncement Error")
		return nil, err
	}
	log.WithField("value", announcement).Debugln("DB FindOneAnnouncement")
	return announcement.ToEntity()
}
