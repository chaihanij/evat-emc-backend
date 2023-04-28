package models

import (
	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson"
)

type UpdateFileScore struct {
	Assingment_uuid string     `bson:"uuid"`
	UploadFile      UploadFile `bson:"files"`
}

func (am *UpdateFileScore) ToEntity() (*entities.AssignmentPartialUpdateScore, error) {
	var assignment entities.AssignmentPartialUpdateScore
	// assignment.UploadFile.FileUrl = am.UploadFile.Fileurl
	// assignment.UploadFile.CreateBy = am.UploadFile.Createby

	err := copier.Copy(&assignment, am)
	return &assignment, err
}

func UploadFileScore(input *entities.AssignmentPartialUpdateScore) *bson.M {
	// updateFields := bson.D{
	// 	bson.E{Key: "updated_at", Value: time.Now()},
	// 	bson.E{Key: "create_by", Value: input.CreateBy},
	// 	bson.E{Key: "filename", Value: input.FileName},
	// }
	// update := bson.D{{Key: "$set", Value: updateFields}}

	updateFields := bson.M{
		"$push": bson.M{
			"files": input.UploadFile,
		},
	}

	// log.WithField("value", updateFields).Debugln("models.PartialUpdateAssignment")
	return &updateFields
}
