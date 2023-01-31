package models

import (
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func PushDocument(documentId string) *bson.M {
	update := &bson.M{
		"$push": bson.E{Key: "documents", Value: documentId},
	}
	log.WithField("value", update).Debugln("models.PushDocument")
	return update
}

func PullDocument(documentId string) *bson.M {
	update := &bson.M{
		"$pull": bson.E{Key: "documents", Value: documentId},
	}
	log.WithField("value", update).Debugln("models.PullDocument")
	return update
}
