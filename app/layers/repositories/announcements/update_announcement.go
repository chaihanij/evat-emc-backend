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

func (r repo) UpdateAnnouncement(ctx context.Context, input *entities.Announcement) (*entities.Announcement, error) {
	log.Debugln("DB UpdateAnnouncement")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	filter := models.NewAnnouncement(input)
	update := models.UpdateAnnouncement(input)
	var announcement models.Announcement
	err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionAnnouncements).
		FindOneAndUpdate(ctx, filter, update, opts).
		Decode(&announcement)
	if err != nil {
		log.WithError(err).Errorln("DB UpdateAnnouncement Error")
		return nil, err
	}
	log.WithField("value", announcement).Debugln("DB UpdateAnnouncement")
	return announcement.ToEntity()
}
