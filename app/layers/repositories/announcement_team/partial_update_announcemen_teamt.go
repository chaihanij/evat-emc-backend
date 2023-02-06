package announcement_teams

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/announcement_team/models"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r repo) PartialUpdateAnnouncementTeam(ctx context.Context, input *entities.AnnouncementTeamPartialUpdate) (*entities.AnnouncementTeam, error) {
	log.Debugln("DB PartialUpdateMember")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	filter := models.NewAnnouncementTeamFilter(input)
	update := models.PartialUpdateAnnouncementTeam(input)
	var announcementTeam models.AnnouncementTeam
	err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionAnnouncementTeams).
		FindOneAndUpdate(ctx, filter, update, opts).
		Decode(&announcementTeam)
	if err != nil {
		log.WithError(err).Errorln("DB PartialUpdateAnnouncement Error")
		return nil, err
	}
	log.WithField("value", announcementTeam).Debugln("DB PartialUpdateAnnouncement")
	return announcementTeam.ToEntity()
}
