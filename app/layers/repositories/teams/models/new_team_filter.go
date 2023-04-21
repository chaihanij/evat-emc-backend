package models

import (
	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewTeamFilter(input interface{}) *bson.M {
	filter := bson.M{}
	if val, ok := input.(*entities.Team); ok {
		if val.UUID != "" {
			filter["uuid"] = val.UUID
		}
	}
	if val, ok := input.(*entities.TeamPartialUpdate); ok {
		if val.UUID != "" {
			filter["uuid"] = val.UUID
		}
	}
	if val, ok := input.(*entities.TeamFilter); ok {
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
		if val.Name != nil {
			filter["name"] = bson.M{
				"$regex":   val.Name,
				"$options": "i",
			}
		}
		if val.TeamType != nil {
			filter["team_type"] = val.TeamType
		}
	}
	log.WithField("value", filter).Debugln("models.NewTeamFilter")
	return &filter
}
