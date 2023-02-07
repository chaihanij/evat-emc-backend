package models

import (
	"time"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateAssignmentTeam(input *entities.AssignmentTeam) *bson.D {
	updateFields := bson.D{
		bson.E{Key: "description", Value: input.Description},
		bson.E{Key: "score", Value: input.Score},
		bson.E{Key: "is_confirmed", Value: input.IsConfirmed},
		bson.E{Key: "updated_at", Value: time.Now()},
		bson.E{Key: "updated_by", Value: input.UpdatedBy},
	}
	if val, ok := input.Documents.([]string); ok {
		updateFields = append(updateFields, bson.E{Key: "documents", Value: val})
	}
	update := bson.D{{Key: "$set", Value: updateFields}}
	log.WithField("value", update).Debugln("models.UpdateAssignmentTeam")
	return &update
}
