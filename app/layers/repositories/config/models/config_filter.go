package models

import (
	"go.mongodb.org/mongo-driver/bson"
)

func ConfigFilter(input interface{}) *bson.M {
	filter := bson.M{}

	return &filter

}
