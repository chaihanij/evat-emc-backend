package models

import (
	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewAnnouncementFilter(input interface{}) *bson.M {
	var filter bson.M = bson.M{}
	if val, ok := input.(*entities.Announcement); ok {
		if val.UUID != "" {
			filter["uuid"] = val.UUID
		}
	}
	if val, ok := input.(*entities.AnnouncementPartialUpdate); ok {
		if val.UUID != "" {
			filter["uuid"] = val.UUID
		}
	}
	if val, ok := input.(*entities.AnnouncementFilter); ok {
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
	}
	log.WithField("value", filter).Debugln("models.NewAnnouncementFilter")
	return &filter
}
