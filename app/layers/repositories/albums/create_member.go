package albums

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/albums/models"
)

func (r repo) CreateAlbum(ctx context.Context, input *entities.Album) (*entities.Album, error) {
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	album := models.NewAlbum(input)
	result, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionAlbums).
		InsertOne(ctx, album)
	if err != nil {
		log.WithError(err).Errorln("DB CreateAlbum Error")
		return nil, err
	}
	if result.InsertedID == 0 {
		return nil, err
	}
	log.WithField("value", album).Debugln("DB CreateAlbum")
	return album.ToEntity()
}
