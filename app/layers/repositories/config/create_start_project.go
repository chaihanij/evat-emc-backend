package config

import (
	"context"

	"github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/config/models"
)

func (r repo) UpdateConfig(ctx context.Context, input *entities.Config) (*entities.Config, error) {
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()

	filter := models.NewFilterConfig(input)
	update := models.UpdateConfig(input)

	var config models.Config
	err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionConfig).
		FindOneAndUpdate(ctx, filter, update, nil).
		Decode(&config)
	if err != nil {
		logrus.WithError(err).Errorln("DB UpdateConfig Error")
		return nil, err
	}
	return config.ToEntity()
}
