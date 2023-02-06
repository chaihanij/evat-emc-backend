package models

import (
	"time"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson"
)

func PartialUpdateAnnouncementTeam(input *entities.AnnouncementTeamPartialUpdate) *bson.D {
	updateFields := bson.D{
		bson.E{Key: "updated_at", Value: time.Now()},
	}
	if input.IsRead != nil {
		updateFields = append(updateFields, bson.E{Key: "is_read", Value: input.IsRead})
	}
	update := bson.D{{Key: "$set", Value: updateFields}}
	log.WithField("value", update).Debugln("models.PartialUpdateAnnouncementTeam")
	return &update
}
