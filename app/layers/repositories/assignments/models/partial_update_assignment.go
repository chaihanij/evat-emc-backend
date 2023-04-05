package models

import (
	"time"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson"
)

func PartialUpdateAssignment(input *entities.AssignmentPartialUpdate) *bson.D {
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
	if input.DueDate != nil {
		updateFields = append(updateFields, bson.E{Key: "due_date", Value: input.DueDate})
	}
	if input.IsActive != nil {
		updateFields = append(updateFields, bson.E{Key: "is_active", Value: input.IsActive})
	}
	if val, ok := input.Image.(*string); ok {
		updateFields = append(updateFields, bson.E{Key: "image", Value: val})
	}
	if val, ok := input.Document.(*string); ok {
		updateFields = append(updateFields, bson.E{Key: "document", Value: val})
	}
	if input.SendDoc != nil {
		updateFields = append(updateFields, bson.E{Key: "senddoc", Value: input.SendDoc})
	}
	if input.DeliveryTime != nil {
		updateFields = append(updateFields, bson.E{Key: "delivery_time", Value: input.DeliveryTime})
	}
	update := bson.D{{Key: "$set", Value: updateFields}}
	log.WithField("value", update).Debugln("models.PartialUpdateAssignment")
	return &update
}
