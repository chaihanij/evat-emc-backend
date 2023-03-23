package models

import (
	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson"
)

func NewAssignmentTeam(input interface{}) *bson.M {
	var filter bson.M = bson.M{}

	if val, ok := input.(*entities.FieldRaceTeamFilter); ok {
		if val.TeamUUID != nil {
			filter["team_uuid"] = val.TeamUUID
		}
	}
	log.WithField("value", filter).Debugln("models.NewAssignmentFilter")
	return &filter
}
