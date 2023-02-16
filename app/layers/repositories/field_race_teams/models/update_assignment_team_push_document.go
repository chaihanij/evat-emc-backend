package models

import (
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateAssignmentTeamPushDocument(documentUUID string) *bson.M {
	update := bson.M{
		"$addToSet": bson.M{
			"documents": bson.M{"$each": "documents"},
		},
	}
	log.WithField("value", update).Debugln("models.UpdateAssignmentTeamPushDocument")
	return &update
}
