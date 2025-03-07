package models

import (
	"time"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson"
)

func PartialUpdateFieldRace(input *entities.FieldRacePartialUpdate) *bson.D {
	updateFields := bson.D{
		bson.E{Key: "updated_at", Value: time.Now()},
		bson.E{Key: "updated_by", Value: input.UpdatedBy},
	}
	if input.No != nil {
		updateFields = append(updateFields, bson.E{Key: "no", Value: input.No})
	}
	if input.Title != nil {
		updateFields = append(updateFields, bson.E{Key: "title", Value: input.Title})
	}
	if input.Description != nil {
		updateFields = append(updateFields, bson.E{Key: "description", Value: input.Description})
	}
	if input.FullScore != nil {
		updateFields = append(updateFields, bson.E{Key: "full_score", Value: input.FullScore})
	}
	if input.Year != nil {
		updateFields = append(updateFields, bson.E{Key: "year", Value: input.Year})
	}

	if val, ok := input.Image.(*string); ok {
		updateFields = append(updateFields, bson.E{Key: "image", Value: val})
	}
	if val, ok := input.Document.(*string); ok {
		updateFields = append(updateFields, bson.E{Key: "document", Value: val})
	}
	update := bson.D{{Key: "$set", Value: updateFields}}
	log.WithField("value", update).Debugln("models.PartialUpdateFieldRace")
	return &update
}
