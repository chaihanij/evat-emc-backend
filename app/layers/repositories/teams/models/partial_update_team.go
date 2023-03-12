package models

import (
	"time"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson"
)

func PartialUpdateTeam(input *entities.TeamPartialUpdate) *bson.D {
	updateFields := bson.D{
		bson.E{Key: "updated_at", Value: time.Now()},
	}
	if input.Code != nil {
		updateFields = append(updateFields, bson.E{Key: "code", Value: input.Code})
	}
	if input.Name != nil {
		updateFields = append(updateFields, bson.E{Key: "name", Value: input.Name})
	}
	if input.TeamType != nil {
		updateFields = append(updateFields, bson.E{Key: "team_type", Value: input.TeamType})
	}
	if input.Academy != nil {
		updateFields = append(updateFields, bson.E{Key: "academy", Value: input.Academy})
	}
	if input.Detail != nil {
		updateFields = append(updateFields, bson.E{Key: "detail", Value: input.Detail})
	}
	if input.Slip != nil {
		updateFields = append(updateFields, bson.E{Key: "slip", Value: input.Slip})
	}
	if input.IsPaid != nil {
		updateFields = append(updateFields, bson.E{Key: "is_paid", Value: input.IsPaid})
	}
	if input.PaymentMethod != nil {
		updateFields = append(updateFields, bson.E{Key: "payment_method", Value: input.PaymentMethod})
	}
	if input.IsVerify != nil {
		updateFields = append(updateFields, bson.E{Key: "is_verify", Value: input.IsVerify})
	}
	if val, ok := input.Members.(*[]string); ok {
		updateFields = append(updateFields, bson.E{Key: "members", Value: val})
	}
	if input.UpdatedBy != nil {
		updateFields = append(updateFields, bson.E{Key: "updated_by", Value: input.UpdatedBy})
	}
	update := bson.D{{Key: "$set", Value: updateFields}}
	log.WithField("value", update).Debugln("models.PartialUpdateTeam")
	return &update
}
