package config

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"

	"gitlab.com/chaihanij/evat/app/layers/repositories/config/models"
)

func (r repo) FindOneConfig(ctx context.Context, input *entities.Config) (*entities.Config, error) {
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	filter := models.NewFilterConfig(input)

	var config models.Config

	err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionConfig).
		FindOne(ctx, filter, nil).
		Decode(&config)
	if err != nil {
		log.WithError(err).Errorln("DB FineOneConfig Error")
		return nil, err
	}

	return config.ToEntity()

}
