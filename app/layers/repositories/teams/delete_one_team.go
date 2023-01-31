package teams

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/errors"
	"gitlab.com/chaihanij/evat/app/layers/repositories/teams/models"
)

func (r repo) DeleteOneTeam(ctx context.Context, input *entities.TeamFilter) error {
	log.Debugln("DB DeleteOneTeam")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	filter := models.NewTeamFilter(input)
	result, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionTeams).
		DeleteOne(ctx, filter, nil)
	if err != nil {
		log.WithError(err).Errorln("DB DeleteOneTeam Error")
		return err
	}
	if result.DeletedCount < 1 {
		return errors.RecordNotFoundError{Message: constants.DataNotFound}
	}
	return nil
}
