package announcement_teams

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/announcement_team/models"
)

func (r repo) FindOneAnnouncementTeam(ctx context.Context, input *entities.AnnouncementTeamFilter) (*entities.AnnouncementTeam, error) {
	log.Debugln("DB FindOneAnnouncementTeam")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	filter := models.NewAnnouncementTeamFilter(input)
	var announcementTeam models.AnnouncementTeam
	err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionAnnouncementTeams).
		FindOne(ctx, filter, nil).
		Decode(&announcementTeam)
	if err != nil {
		log.WithError(err).Errorln("DB FindOneAnnouncementTeam Error")
		return nil, err
	}
	log.WithField("value", announcementTeam).Debugln("DB FindOneAnnouncementTeam")
	return announcementTeam.ToEntity()
}
