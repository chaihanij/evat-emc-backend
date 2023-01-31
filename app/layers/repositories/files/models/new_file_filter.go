package models

import (
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewFileFilter(input interface{}) *bson.M {
	filter := bson.M{}
	if val, ok := input.(*entities.File); ok {
		if val.ID != nil {
			id, _ := primitive.ObjectIDFromHex(*val.ID)
			filter["_id"] = id
		}
	} else if val, ok := input.(*entities.FileFilter); ok {
		if val.ID != nil {
			id, _ := primitive.ObjectIDFromHex(*val.ID)
			filter["_id"] = id
		}
		if val.UUID != nil {
			filter["uuid"] = val.UUID
		}
		if val.OriginalFileName != nil {
			filter["original_filename"] = val.OriginalFileName
		}
		if val.FileName != nil {
			filter["filename"] = val.FileName
		}
		if val.FilePath != nil {
			filter["file_path"] = val.FilePath
		}
	}
	return &filter
}
