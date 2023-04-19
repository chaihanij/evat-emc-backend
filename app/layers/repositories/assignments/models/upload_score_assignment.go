package models

import (
	"time"

	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson"
)

func UploadScoreAssignment(input *entities.Assignment) *bson.M {
	updateFields := bson.M{
		"$set": bson.M{
			"consideration": input.Consideration,
			"updated_at":    time.Now(),
		},
	}
	return &updateFields

}
