package announcement_teams

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/assignment_teams/models"
)

func (r repo) CountAnnouncementTeam(ctx context.Context, input *entities.AssignmentTeamFilter) (*int64, error) {
	log.Debugln("DB CountAnnouncementTeam")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()

	filter := models.NewAssignmentTeamFilter(input)
	count, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionAnnouncementTeams).
		CountDocuments(ctx, filter)

	if err != nil {
		log.WithError(err).Errorln("DB CountAnnouncementTeam Error")
		return nil, err
	}
	log.WithField("value", count).Debugln("DB CountAnnouncementTeam")
	return &count, nil
}
