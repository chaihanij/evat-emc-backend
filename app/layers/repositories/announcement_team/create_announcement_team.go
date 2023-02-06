package announcement_teams

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/announcement_team/models"
)

func (r repo) CreateAnnouncementTeam(ctx context.Context, input *entities.AnnouncementTeam) (*entities.AnnouncementTeam, error) {
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	announcementTeam := models.NewAnnouncementTeam(input)
	result, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionAnnouncementTeams).
		InsertOne(ctx, announcementTeam)
	if err != nil {
		log.WithError(err).Errorln("DB CreateAnnouncementTeam Error")
		return nil, err
	}
	if result.InsertedID == 0 {
		return nil, err
	}
	log.WithField("value", announcementTeam).Debugln("DB CreateAnnouncementTeam")
	return announcementTeam.ToEntity()
}
