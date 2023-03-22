package models

import (
	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

func NewFieldRaceFilter(input interface{}) *bson.M {
	var filter bson.M = bson.M{}

	if val, ok := input.(*entities.FieldRaceFilter); ok {
		if val.UUID != nil {
			filter["uuid"] = val.UUID
		}

	}
	log.WithField("value", filter).Debugln("models.NewFieldRaceFilter")
	return &filter
}
