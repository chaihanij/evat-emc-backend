package models

import (
	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewAssignmentFilter(input interface{}) *bson.M {
	var filter bson.M = bson.M{}
	if val, ok := input.(*entities.Assignment); ok {
		if val.UUID != "" {
			filter["uuid"] = val.UUID
		}
	}
	if val, ok := input.(*entities.AssignmentPartialUpdate); ok {
		if val.UUID != "" {
			filter["uuid"] = val.UUID
		}
	}
	if val, ok := input.(*entities.AssignmentFilter); ok {
		if val.ID != nil {
			id, _ := primitive.ObjectIDFromHex(*val.ID)
			filter["_id"] = id
		}
		if val.UUID != nil {
			filter["uuid"] = val.UUID
		}
		if val.Year != nil {
			filter["year"] = val.Year
		}
		if val.UUID != nil {
			filter["uuid"] = val.UUID
		}
	}

	// if val, ok := input.(*entities.UploadFile); ok {
	// 	if val.UUID != nil {
	// 		filter["uuid"] = val.UUID
	// 	}
	// }

	log.WithField("value", filter).Debugln("models.NewAssignmentFilter")
	return &filter
}

func NewTeamFilter(input interface{}) *bson.M {
	var filter bson.M = bson.M{}

	log.WithField("value", filter).Debugln("models.NewTeamFilter")
	return &filter
}



func UpdateFileAssignmentFilter(input string) *bson.M {
	var filter bson.M = bson.M{}

	if input != "" {
		filter["uuid"] = input
	}

	// if val, ok := input.(*entities.Assignment); ok {
	// 	if val.UUID != "" {
	// 		filter["uuid"] = val.UUID
	// 	}
	// }
	// if val, ok := input.(*entities.AssignmentPartialUpdate); ok {
	// 	if val.UUID != "" {
	// 		filter["uuid"] = val.UUID
	// 	}
	// }
	// if val, ok := input.(*entities.AssignmentFilter); ok {
	// 	if val.ID != nil {
	// 		id, _ := primitive.ObjectIDFromHex(*val.ID)
	// 		filter["_id"] = id
	// 	}
	// 	if val.UUID != nil {
	// 		filter["uuid"] = val.UUID
	// 	}
	// 	if val.Year != nil {
	// 		filter["year"] = val.Year
	// 	}
	// 	if val.UUID != nil {
	// 		filter["uuid"] = val.UUID
	// 	}
	// }

	// if val, ok := input.(*entities.UploadFile); ok {
	// 	if val.UUID != nil {
	// 		filter["uuid"] = val.UUID
	// 	}
	// }

	log.WithField("value", filter).Debugln("models.NewAssignmentFilter")
	return &filter
}
