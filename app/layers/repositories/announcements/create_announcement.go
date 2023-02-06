package announcements

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/announcements/models"
)

func (r repo) CreateAnnouncement(ctx context.Context, input *entities.Announcement) (*entities.Announcement, error) {
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	announcement := models.NewAnnouncement(input)
	result, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionAnnouncements).
		InsertOne(ctx, announcement)
	if err != nil {
		log.WithError(err).Errorln("DB CreateAnnouncement Error")
		return nil, err
	}
	if result.InsertedID == 0 {
		return nil, err
	}
	log.WithField("value", announcement).Debugln("DB CreateAnnouncement")
	return announcement.ToEntity()
}
