package users

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/users/models"
)

func (r repo) FindOneUser(ctx context.Context, input *entities.UserFilter) (*entities.User, error) {
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	filter := models.FilterUserCriteria(input)
	var user models.User
	err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionUsers).
		FindOne(ctx, filter, nil).
		Decode(&user)
	if err != nil {
		log.WithError(err).Errorln("DB FindOneUser Error")
		return nil, err
	}
	return user.ToEntity()
}
