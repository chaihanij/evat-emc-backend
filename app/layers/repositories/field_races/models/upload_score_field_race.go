package models

import (
	"fmt"
	"time"

	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson"
)

func UploadScoreFieldRace(input *entities.FieldRace) *bson.M {
	fmt.Println("input", input)
	updateFields := bson.M{
		"$set": bson.M{
			"consideration": input.Consideration,
			"updated_at":    time.Now(),
		},
	}
	return &updateFields

}
