package files

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/files/models"
)

func (r repo) CreateFiles(ctx context.Context, input []entities.File) ([]entities.File, error) {
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	files := models.NewFiles(input)
	var documents []interface{}
	for _, v := range files {
		documents = append(documents, v)
	}
	_, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionFiles).
		InsertMany(ctx, documents)
	if err != nil {
		log.WithError(err).Errorln("DB CreateFiles Error")
		return nil, err
	}
	return files.ToEntity(), nil
}
