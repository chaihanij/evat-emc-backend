package models

import (
	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson"
)

func NewAssignmentTeamFilter(input interface{}) *bson.M {
	var filter bson.M = bson.M{}
	if val, ok := input.(*entities.AssignmentTeam); ok {
		if val.UUID != "" {
			filter["uuid"] = val.UUID
		}
	}
	if val, ok := input.(*entities.AssignmentTeamFilter); ok {
		if val.UUID != nil {
			filter["uuid"] = val.UUID
		}
	}
	if val, ok := input.(*entities.AssignmentTeamPartialUpdate); ok {
		if val.UUID != nil {
			filter["uuid"] = val.UUID
		}
	}
	log.WithField("value", filter).Debugln("models.NewAssignmentFilter")
	return &filter
}
