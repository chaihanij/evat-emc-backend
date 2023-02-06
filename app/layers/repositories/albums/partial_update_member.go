package albums

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/albums/models"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r repo) PartialUpdateAlbum(ctx context.Context, input *entities.AlbumPartialUpdate) (*entities.Album, error) {
	log.Debugln("DB PartialUpdateAlbum")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	filter := models.NewAlbumFilter(input)
	update := models.PartialUpdateAlbum(input)
	var album models.Album
	err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionAlbums).
		FindOneAndUpdate(ctx, filter, update, opts).
		Decode(&album)
	if err != nil {
		log.WithError(err).Errorln("DB PartialUpdateAlbum Error")
		return nil, err
	}
	log.WithField("value", album).Debugln("DB PartialUpdateAlbum")
	return album.ToEntity()
}
