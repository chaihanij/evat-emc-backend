package albums

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/albums/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r repo) getFindOptions(input *entities.AlbumFilter) *options.FindOptions {
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

func (r repo) FindAllAlbum(ctx context.Context, input *entities.AlbumFilter) ([]entities.Album, error) {
	log.Debugln("DB FindAllAlbum")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	findOptions := r.getFindOptions(input)
	filter := models.NewAlbumFilter(input)
	cursor, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionAlbums).
		Find(ctx, filter, findOptions)
	if err != nil {
		log.WithError(err).Errorln("DB FineAllAlbum Error")
		return nil, err
	}
	var albums models.Albums
	err = cursor.All(ctx, &albums)
	if err != nil {
		log.WithError(err).Errorln("DB FineAllAlbum Error")
		return nil, err
	}
	log.WithField("value", albums).Debugln("DB FindAllAlbum")
	return albums.ToEntity(), nil
}
