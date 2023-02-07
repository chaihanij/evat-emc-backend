package models

import (
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateAssignmentTeamPullDocument(documentUUID string) *bson.M {
	update := bson.M{
		"$pull": bson.M{
			"documents": documentUUID,
		},
	}
	log.WithField("value", update).Debugln("models.UpdateAssignmentTeamPullDocument")
	return &update
}
