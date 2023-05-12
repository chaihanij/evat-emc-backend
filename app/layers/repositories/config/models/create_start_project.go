package models

import (
	"time"

	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateConfig(input *entities.Config) *bson.D {
	updateFields := bson.D{bson.E{Key: "update_at", Value: time.Now()}}

	if input.RegisterConfig != nil {
		updateFields = append(updateFields, bson.E{Key: "registerconfig", Value: input.RegisterConfig})
	}
	if input.StartProject != nil {
		updateFields = append(updateFields, bson.E{Key: "startproject", Value: input.StartProject})
	}

	// fmt.Println("ee :", *&input.StartProject)
	// if val, ok := input.StartProject ; ok {

	// updateFields = append(updateFields, bson.E{Key: "start_project", Value: val.Start_date})
	// updateFields = append(updateFields, bson.E{Key: "end_project", Value: val.End_date})

	// }

	// if  *&input.StartProject != nil {
	// 	updateFields = append(updateFields, bson.E{Key: "start_project", Value: input.Start_date})
	// 	updateFields = append(updateFields, bson.E{Key: "end_project", Value: val.End_date})
	// }

	// data := entities.Config{
	// 	RegisterConfig: input.RegisterConfig,
	// 	StartProject:   input.StartProject,
	// }

	update := bson.D{{Key: "$set", Value: input}}

	return &update

}
