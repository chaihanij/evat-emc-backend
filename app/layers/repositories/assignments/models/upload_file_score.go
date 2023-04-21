package models

import (
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson"
)

func UploadFileScore(input *entities.Assignment) *bson.M {
	// updateFields := bson.D{
	// 	bson.E{Key: "updated_at", Value: time.Now()},
	// 	bson.E{Key: "create_by", Value: input.CreateBy},
	// 	bson.E{Key: "filename", Value: input.FileName},
	// }
	// update := bson.D{{Key: "$set", Value: updateFields}}

	updateFields := bson.M{
		"$set": bson.M{
			"files": input.UploadFile,
		},
	}

	// log.WithField("value", updateFields).Debugln("models.PartialUpdateAssignment")
	return &updateFields
}
