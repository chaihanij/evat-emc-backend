package logsetting

import (
	"context"

	"github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/logsetting/models"
)

func (r repo) CreateLogSetting(ctx context.Context, input *entities.LogSetting) (*entities.LogSetting, error) {
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	logSetting := models.NewCreateLogSetting(input)
	_, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionSetting).
		InsertOne(ctx, logSetting)
	if err != nil {
		logrus.Debugln("Can't keep Log Setting")
	}
	return logSetting.ToEntity()
}
