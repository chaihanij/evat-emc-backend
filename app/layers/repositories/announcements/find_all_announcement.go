package announcements

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/announcements/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r repo) getFindOptions(input *entities.AnnouncementFilter) *options.FindOptions {
	findOptions := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})
	if input.PageSize != nil && input.Page != nil {
		pageSize := *input.PageSize
		page := *input.Page
		findOptions.SetLimit(*input.PageSize)
		offset := (page - 1) * pageSize
		findOptions.SetSkip(offset)
	}
	return findOptions
}

func (r repo) FindAllAnnouncement(ctx context.Context, input *entities.AnnouncementFilter) ([]entities.Announcement, error) {
	log.Debugln("DB FindAllAnnouncement")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	findOptions := r.getFindOptions(input)
	filter := models.NewAnnouncementFilter(input)
	cursor, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionAnnouncements).
		Find(ctx, filter, findOptions)
	if err != nil {
		log.WithError(err).Errorln("DB FindAllAnnouncement Error")
		return nil, err
	}
	var announcements models.Announcements
	err = cursor.All(ctx, &announcements)
	if err != nil {
		log.WithError(err).Errorln("DB FindAllAnnouncement Error")
		return nil, err
	}
	log.WithField("value", announcements).Debugln("DB FindAllAnnouncement")
	return announcements.ToEntity(), nil
}
