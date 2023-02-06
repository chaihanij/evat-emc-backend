package models

import (
	"time"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateFieldRace(input *entities.FieldRace) *bson.D {
	updateFields := bson.D{
		bson.E{Key: "no", Value: input.No},
		bson.E{Key: "title", Value: input.Title},
		bson.E{Key: "description", Value: input.Description},
		bson.E{Key: "full_score", Value: input.FullScore},
		bson.E{Key: "year", Value: input.Year},
		bson.E{Key: "is_active", Value: input.IsActive},
		bson.E{Key: "updated_at", Value: time.Now()},
		bson.E{Key: "updated_by", Value: input.UpdatedBy},
	}

	if val, ok := input.Image.(string); ok {
		updateFields = append(updateFields, bson.E{Key: "image", Value: val})
	}
	if val, ok := input.Document.(string); ok {
		updateFields = append(updateFields, bson.E{Key: "document", Value: val})
	}
	update := bson.D{{Key: "$set", Value: updateFields}}
	log.WithField("value", update).Debugln("models.UpdateFieldRace")
	return &update
}
