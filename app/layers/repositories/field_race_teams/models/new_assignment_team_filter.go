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

	log.WithField("value", filter).Debugln("models.NewFieldRaceTeamFilter")
	return &filter
}
