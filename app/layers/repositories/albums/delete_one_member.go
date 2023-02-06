package albums

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/errors"
	"gitlab.com/chaihanij/evat/app/layers/repositories/albums/models"
)

func (r repo) DeleteAlbum(ctx context.Context, input *entities.AlbumFilter) error {
	log.Debugln("DB DeleteAlbum")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	filter := models.NewAlbumFilter(input)
	result, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionAlbums).
		DeleteOne(ctx, filter, nil)

	if err != nil {
		log.WithError(err).Errorln("DB DeleteAlbum Error")
		return err
	}

	if result.DeletedCount < 1 {
		return errors.RecordNotFoundError{Message: constants.DataNotFound}
	}

	return nil
}
