package models

import (
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateAssignmentTeamPushDocument(documentUUID string, assignmentTopic string) *bson.M {
	doc := []string{}
	doc = append(doc, documentUUID)

	// file_uuid
	// assignment_topic

	DocumentsAssignment := new(DocumentAssignment)

	DocumentsAssignment.AssignmentTopic = assignmentTopic
	DocumentsAssignment.FileUUID = documentUUID

	update := bson.M{
		// "$addToSet": bson.M{
		// 	"documents": bson.M{"$each": doc},
		// },
		"$push": bson.M{
			"document": DocumentsAssignment,
		},
	}
	log.WithField("value", update).Debugln("models.UpdateAssignmentTeamPushDocument")
	return &update
}
