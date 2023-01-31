package teams

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/teams/models"
)

func (r repo) FindOneTeam(ctx context.Context, input *entities.TeamFilter) (*entities.Team, error) {
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	filter := models.NewTeamFilter(input)
	var team models.Team
	err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionTeams).
		FindOne(ctx, filter, nil).
		Decode(&team)
	if err != nil {
		log.WithError(err).Errorln("DB FineOneMemeber Error")
		return nil, err
	}
	return team.ToEntity()
}
