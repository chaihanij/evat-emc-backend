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

func (r repo) UpdateAnnouncementTeam(ctx context.Context, input *entities.AnnouncementTeam) (*entities.AnnouncementTeam, error) {
	log.Debugln("DB UpdateAnnouncementTeam")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	filter := models.NewAnnouncementTeam(input)
	update := models.UpdateAnnouncementTeam(input)
	var announcementTeam models.AnnouncementTeam
	err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionAnnouncements).
		FindOneAndUpdate(ctx, filter, update, opts).
		Decode(&announcementTeam)
	if err != nil {
		log.WithError(err).Errorln("DB UpdateAnnouncementTeam Error")
		return nil, err
	}
	log.WithField("value", announcementTeam).Debugln("DB UpdateAnnouncementTeam")
	return announcementTeam.ToEntity()
}
