package models

import (
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateAssignmentTeamPushDocument(documentUUID string) *bson.M {
	doc := []string{}
	doc = append(doc, documentUUID)
	update := bson.M{
		"$addToSet": bson.M{
			"documents": bson.M{"$each": doc},
		},
	}
	log.WithField("value", update).Debugln("models.UpdateAssignmentTeamPushDocument")
	return &update
}
