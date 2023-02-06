package models

import (
	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson"
)

func NewAnnouncementTeamFilter(input interface{}) *bson.M {
	var filter bson.M = bson.M{}
	if val, ok := input.(*entities.AnnouncementTeam); ok {
		if val.UUID != "" {
			filter["uuid"] = val.UUID
		}
	}
	if val, ok := input.(*entities.AnnouncementTeamFilter); ok {
		if val.UUID != nil {
			filter["uuid"] = val.UUID
		}
		if val.TeamUUID != nil {
			filter["team_uuid"] = val.TeamUUID
		}
		if val.AnnouncementUUID != nil {
			filter["announcement_uuid"] = val.AnnouncementUUID
		}
		if val.IsRead != nil {
			filter["is_read"] = val.IsRead
		}
	}
	log.WithField("value", filter).Debugln("models.NewAnnouncementTeamFilter")
	return &filter
}
