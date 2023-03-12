package models

import (
	"time"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateTeam(input *entities.Team) *bson.D {

	updateFields := bson.D{
		bson.E{Key: "code", Value: input.Code},
		bson.E{Key: "name", Value: input.Name},
		bson.E{Key: "team_type", Value: input.TeamType},
		bson.E{Key: "academy", Value: input.Academy},
		bson.E{Key: "detail", Value: input.Detail},
		bson.E{Key: "year", Value: input.Year},
		bson.E{Key: "payment_method", Value: input.PaymentMethod},
		bson.E{Key: "is_paid", Value: input.IsPaid},
		bson.E{Key: "is_verify", Value: input.IsVerify},
		bson.E{Key: "updated_at", Value: time.Now()},
		bson.E{Key: "updated_by", Value: input.UpdatedBy},
	}
	if val, ok := input.Slip.(string); ok {
		updateFields = append(updateFields, bson.E{Key: "members", Value: val})
	}
	if val, ok := input.Members.([]string); ok {
		updateFields = append(updateFields, bson.E{Key: "members", Value: val})
	}
	update := bson.D{{Key: "$set", Value: updateFields}}
	log.WithField("value", update).Debugln("models.UpdateTeam")
	return &update
}
