package teams

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/teams/models"
)

func (r repo) CountTeam(ctx context.Context, input *entities.TeamFilter) (*int64, error) {
	log.Debugln("DB CountTeam")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()

	filter := models.NewTeamFilter(input)
	count, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionTeams).
		CountDocuments(ctx, filter)

	if err != nil {
		log.WithError(err).Errorln("DB CountMember Error")
		return nil, err
	}
	log.WithField("value", count).Debugln("DB CountTeam")
	return &count, nil
}
