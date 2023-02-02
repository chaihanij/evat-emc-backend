package files

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/files/models"
)

func (r repo) FindAllFile(ctx context.Context, input interface{}) ([]entities.File, error) {
	log.Debugln("FindOneFile")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	filter := models.NewFileFilter(input)
	cursor, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionFiles).
		Find(ctx, filter, nil)

	if err != nil {
		log.WithError(err).Errorln("DB FindAllFile Error")
		return nil, err
	}
	var files models.Files
	err = cursor.All(ctx, &files)
	if err != nil {
		log.WithError(err).Errorln("DB FindAllFile Error")
		return nil, err
	}
	log.WithField("value", files).Debugln("DB FindAllFile")
	return files.ToEntity(), nil
}
