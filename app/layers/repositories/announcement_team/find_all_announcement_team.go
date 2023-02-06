package announcement_teams

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/layers/repositories/announcement_team/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r repo) getFindOptions(input *entities.AnnouncementTeamFilter) *options.FindOptions {
	findOptions := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})
	return findOptions
}

func (r repo) FindAllAnnouncementTeam(ctx context.Context, input *entities.AnnouncementTeamFilter) ([]entities.AnnouncementTeam, error) {
	log.Debugln("DB FindAllAnnouncementTeam")
	ctx, cancel := context.WithTimeout(ctx, env.MongoDBRequestTimeout)
	defer cancel()
	findOptions := r.getFindOptions(input)
	filter := models.NewAnnouncementTeamFilter(input)
	cursor, err := r.MongoDBClient.Database(env.MongoDBName).
		Collection(constants.CollectionAnnouncementTeams).
		Find(ctx, filter, findOptions)
	if err != nil {
		log.WithError(err).Errorln("DB FindAllAnnouncementTeam Error")
		return nil, err
	}
	var announcementTeams models.AnnouncementTeams
	err = cursor.All(ctx, &announcementTeams)
	if err != nil {
		log.WithError(err).Errorln("DB FindAllAnnouncementTeam Error")
		return nil, err
	}
	log.WithField("value", announcementTeams).Debugln("DB FindAllAnnouncementTeam")
	return announcementTeams.ToEntity(), nil
}
