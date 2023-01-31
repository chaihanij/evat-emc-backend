package files

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/errors"
	"gitlab.com/chaihanij/evat/app/layers/repositories/files/models"
)

func (r repo) CreateFile(ctx context.Context, input *entities.File) (*entities.File, error) {
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	file := models.NewFile(input.OriginalFileName, input.FileName, input.FileExtension, input.FileName)
	result, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionFiles).
		InsertOne(ctx, file)
	if err != nil {
		log.WithError(err).Errorln("DB FindOneFile Error")
		return nil, err
	}
	if result.InsertedID == 0 {
		return nil, errors.InternalError{Message: "create file fail"}
	}
	return file.ToEntity(), nil
}
