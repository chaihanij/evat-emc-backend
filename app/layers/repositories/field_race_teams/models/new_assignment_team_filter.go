package models

import (
	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson"
)

func NewFieldRaceTeamFilter(input interface{}) *bson.M {
	var filter bson.M = bson.M{}
	if val, ok := input.(*entities.FieldRaceTeam); ok {
		if val.FieldRaceUUID != "" {
			filter["field_race_uuid"] = val.FieldRaceUUID
		}
		if val.TeamUUID != "" {
			filter["team_uuid"] = val.TeamUUID
		}
	}
	if val, ok := input.(*entities.AssignmentTeamFilter); ok {
		if val.AssignmentUUID != nil {
			filter["assignment_uuid"] = val.AssignmentUUID
		}
		if val.TeamUUID != nil {
			filter["team_uuid"] = val.TeamUUID
		}
	}
	if val, ok := input.(*entities.AssignmentTeamPartialUpdate); ok {
		if val.AssignmentUUID != nil {
			filter["assignment_uuid"] = val.AssignmentUUID
		}
		if val.TeamUUID != nil {
			filter["team_uuid"] = val.TeamUUID
		}
	}
	log.WithField("value", filter).Debugln("models.NewAssignmentFilter")
	return &filter
}
