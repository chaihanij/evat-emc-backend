package models

import (
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func PushImages(imageUUIDs []string) *bson.M {
	update := bson.M{
		"$push": bson.M{
			"images": imageUUIDs,
		},
	}
	log.WithField("value", update).Debugln("models.PushImage")
	return &update
}

func PullImage(imageUUID string) *bson.M {
	update := bson.M{
		"$pull": bson.M{
			"images": imageUUID,
		},
	}
	log.WithField("value", update).Debugln("models.PullImage")
	return &update
}
