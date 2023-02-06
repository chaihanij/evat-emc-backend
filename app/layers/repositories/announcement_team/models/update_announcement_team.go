package models

import (
	"time"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateAnnouncementTeam(input *entities.AnnouncementTeam) *bson.D {
	updateFields := bson.D{
		bson.E{Key: "is_read", Value: input.IsRead},
		bson.E{Key: "updated_at", Value: time.Now()},
	}
	update := bson.D{{Key: "$set", Value: updateFields}}
	log.WithField("value", update).Debugln("models.UpdateAnnouncementTeam")
	return &update
}
