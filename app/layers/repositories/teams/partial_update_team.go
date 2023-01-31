package teams

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/teams/models"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r repo) PartialUpdateTeam(ctx context.Context, input *entities.TeamPartialUpdate) (*entities.Team, error) {
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	filter := models.NewTeamFilter(input)
	update := models.PartialUpdateTeam(input)
	var team models.Team
	err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionTeams).
		FindOneAndUpdate(ctx, filter, update, opts).
		Decode(&team)
	if err != nil {
		log.WithError(err).Errorln("DB PartialUpdateTeam Error")
		return nil, err
	}
	return team.ToEntity()
}
