package announcement_teams

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/errors"
	"gitlab.com/chaihanij/evat/app/layers/repositories/announcement_team/models"
)

func (r repo) DeleteAnnouncementTeam(ctx context.Context, input *entities.AnnouncementTeamFilter) error {
	log.Debugln("DB DeleteAnnouncementTeam")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	filter := models.NewAnnouncementTeamFilter(input)
	result, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionAnnouncementTeams).
		DeleteOne(ctx, filter, nil)

	if err != nil {
		log.WithError(err).Errorln("DB DeleteAnnouncementTeam Error")
		return err
	}

	if result.DeletedCount < 1 {
		return errors.RecordNotFoundError{Message: constants.DataNotFound}
	}

	return nil
}
