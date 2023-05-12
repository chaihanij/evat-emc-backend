package models

import (
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson"
)

func NewFilterConfig(input interface{}) *bson.M {
	var filter bson.M = bson.M{}
	if val, ok := input.(*entities.Config); ok {
		if val.UUID != "" {
			filter["uuid"] = val.UUID
		}
	}
	return &filter

}
