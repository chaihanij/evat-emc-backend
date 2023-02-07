package models

import (
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateAssignmentTeamPushDocument(documentUUIDs []string) *bson.M {
	update := bson.M{
		"$push": bson.M{
			"documents": documentUUIDs,
		},
	}
	log.WithField("value", update).Debugln("models.UpdateAssignmentTeamPushDocument")
	return &update
}
