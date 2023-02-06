package announcements

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/errors"
	"gitlab.com/chaihanij/evat/app/layers/repositories/announcements/models"
)

func (r repo) DeleteAnnouncement(ctx context.Context, input *entities.AnnouncementFilter) error {
	log.Debugln("DB DeleteOnAssignment")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	filter := models.NewAnnouncementFilter(input)
	result, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionAnnouncements).
		DeleteOne(ctx, filter, nil)

	if err != nil {
		log.WithError(err).Errorln("DB DeleteAnnouncement Error")
		return err
	}

	if result.DeletedCount < 1 {
		return errors.RecordNotFoundError{Message: constants.DataNotFound}
	}

	return nil
}
