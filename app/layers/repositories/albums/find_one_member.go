package albums

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/albums/models"
)

func (r repo) FindOneAlbum(ctx context.Context, input *entities.AlbumFilter) (*entities.Album, error) {
	log.Debugln("DB FindOneMemeber")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	filter := models.NewAlbumFilter(input)
	var album models.Album
	err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionAlbums).
		FindOne(ctx, filter, nil).
		Decode(&album)
	if err != nil {
		log.WithError(err).Errorln("DB FindOneAlbum Error")
		return nil, err
	}
	log.WithField("value", album).Debugln("DB FindOneAlbum")
	return album.ToEntity()
}
