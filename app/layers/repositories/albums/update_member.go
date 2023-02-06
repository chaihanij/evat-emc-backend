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

func (r repo) UpdateAlbum(ctx context.Context, input *entities.Album) (*entities.Album, error) {
	log.Debugln("DB UpdateAlbum")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	filter := models.NewAlbumFilter(input)
	update := models.UpdateAlbum(input)
	var album models.Album
	err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionAlbums).
		FindOneAndUpdate(ctx, filter, update, opts).
		Decode(&album)
	if err != nil {
		log.WithError(err).Errorln("DB UpdateAlbum Error")
		return nil, err
	}
	log.WithField("value", album).Debugln("DB UpdateAlbum")
	return album.ToEntity()
}
