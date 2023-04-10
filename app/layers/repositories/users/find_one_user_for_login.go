package users

import (
	"context"
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/users/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (r repo) FindOneUserLogin(ctx context.Context, input *entities.UserFilter) (*entities.User, error) {
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	Email := strings.ToLower(*input.Email)
	EmailUser := fmt.Sprintf("^%v$", Email)
	filter := bson.M{
		"email": bson.M{
			"$regex":   EmailUser,
			"$options": "i",
		},
		"is_active": true,
	}
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
