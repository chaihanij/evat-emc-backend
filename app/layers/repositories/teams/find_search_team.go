package teams

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/teams/models"
)

func (r repo) FindAllSearchTeam(ctx context.Context, input *entities.TeamFilter) ([]entities.TeamSearch, error) {
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	findOptions := r.getFindOptions(input)
	filter := models.NewTeamFilter(input)
	cursor, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionTeams).
		Find(ctx, filter, findOptions)
	if err != nil {
		log.WithError(err).Errorln("DB FindAllSearchTeam Error")
		return nil, err
	}
	var teams models.TeamsSearch
	err = cursor.All(ctx, &teams)
	if err != nil {
		log.WithError(err).Errorln("DB FindAllSearchTeam Error")
		return nil, err
	}
	return teams.ToEntityTeamSearch()   , nil
}
