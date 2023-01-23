package users

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/users/models"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r repo) PartialUpdateUser(ctx context.Context, userFilter *entities.UserFilter, userPartialUpdate *entities.UserPartialUpdate) (*entities.User, error) {
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	filter := models.FilterUserCriteria(userFilter)
	statement := models.ToUpdateUserFields(userPartialUpdate)
	var user models.User
	err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionUsers).
		FindOneAndUpdate(ctx, filter, statement, opts).
		Decode(&user)
	if err != nil {
		log.WithError(err).Errorln("DB PartialUpdateUser Error")
		return nil, err
	}
	return user.ToEntity()
}
