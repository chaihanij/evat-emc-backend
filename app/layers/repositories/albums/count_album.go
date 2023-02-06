package albums

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/albums/models"
)

func (r repo) CountAlbum(ctx context.Context, input *entities.AlbumFilter) (*int64, error) {
	log.Debugln("DB CountAlbum")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()

	filter := models.NewAlbumFilter(input)
	count, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionMembers).
		CountDocuments(ctx, filter)

	if err != nil {
		log.WithError(err).Errorln("DB CountAlbum Error")
		return nil, err
	}
	log.WithField("value", count).Debugln("DB CountAlbum")
	return &count, nil
}
