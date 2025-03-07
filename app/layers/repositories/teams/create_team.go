package teams

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/errors"
	"gitlab.com/chaihanij/evat/app/layers/repositories/teams/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r repo) CreateTeam(ctx context.Context, input *entities.Team) (*entities.Team, error) {
	log.Debugln("DB CreateTeam")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	filter := bson.M{}
	count, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionTeams).
		CountDocuments(ctx, filter)

		// eMC23S006
	paddedStr := fmt.Sprintf("%03d", count)
	year := time.Now().Year()
	str := fmt.Sprintf("%d", year)
	trimmedStr := str[2:]
	code := fmt.Sprintf("eMC%sS%s", trimmedStr, paddedStr)

	input.Code = code

	team := models.NewTeam(input)
	result, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionTeams).
		InsertOne(ctx, team)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"is_duplicate_key_error": mongo.IsDuplicateKeyError(err),
		}).Errorln("DB CreateTeam Error")
		if mongo.IsDuplicateKeyError(err) {
			return nil, errors.DuplicateKeyError{Message: err.Error()}
		}
		return nil, err
	}
	if result.InsertedID == 0 {
		return nil, err
	}
	log.WithField("value", team).Debugln("DB CreateTeam")
	return team.ToEntity()
}
