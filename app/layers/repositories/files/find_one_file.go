package files

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/files/models"
)

func (r repo) FindOneFile(ctx context.Context, input interface{}) (*entities.File, error) {
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	filter := models.NewFileFilter(input)
	var file models.File
	err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionFiles).
		FindOne(ctx, filter, nil).
		Decode(&file)
	if err != nil {
		log.WithError(err).Errorln("DB FindOneFile Error")
		return nil, err
	}
	return file.ToEntity(), nil
}
