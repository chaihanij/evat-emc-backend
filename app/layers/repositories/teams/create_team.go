package teams

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/errors"
	"gitlab.com/chaihanij/evat/app/layers/repositories/teams/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r repo) CreateTeam(ctx context.Context, input *entities.Team) (*entities.Team, error) {
	log.Debugln("DB CreateTeam")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	team := models.NewTeam(input)
	result, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionTeams).
		InsertOne(ctx, team)
	if err != nil {
		log.WithError(err).Errorln("DB CreateTeam Error")
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
