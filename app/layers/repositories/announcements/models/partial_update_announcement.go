package models

import (
	"time"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson"
)

func PartialUpdateAnnouncement(input *entities.AnnouncementPartialUpdate) *bson.D {
	updateFields := bson.D{
		bson.E{Key: "updated_at", Value: time.Now()},
		bson.E{Key: "updated_by", Value: input.UpdatedBy},
	}
	if input.Title != nil {
		updateFields = append(updateFields, bson.E{Key: "title", Value: input.Title})
	}
	if input.Description != nil {
		updateFields = append(updateFields, bson.E{Key: "description", Value: input.Description})
	}
	if input.Year != nil {
		updateFields = append(updateFields, bson.E{Key: "year", Value: input.Year})
	}
	update := bson.D{{Key: "$set", Value: updateFields}}
	log.WithField("value", update).Debugln("models.PartialUpdateAnnouncement")
	return &update
}
