package models

import (
	"time"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateMemberCheckIn(input *entities.MemberCheckIn) *bson.D {
	updateFields := bson.D{bson.E{Key: "checkin_date", Value: time.Now()}}

	if input.Check_national != nil {
		updateFields = append(updateFields, bson.E{Key: "is_national", Value: input.Check_national})
	}

	if input.Is_Check_image != nil {
		updateFields = append(updateFields, bson.E{Key: "is_image", Value: input.Is_Check_image})
	}
	if input.Is_check_data != nil {
		updateFields = append(updateFields, bson.E{Key: "is_data", Value: input.Is_check_data})
	}

	if input.Is_checkin != nil {
		updateFields = append(updateFields, bson.E{Key: "is_checkin", Value: input.Is_checkin})
	}
	update := bson.D{{Key: "$set", Value: updateFields}}
	log.WithField("value", update).Debugln("models.MemberCheckeIn")

	return &update

}
