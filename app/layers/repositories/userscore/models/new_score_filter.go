package medles

import (
	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewScoreFilter(input interface{}) *bson.M {
	filter := bson.M{}
	if val, ok := input.(*entities.Score); ok {
		if val.ID != "" {
			filter["uuid"] = val.ID
		}
	}
	if val, ok := input.(*entities.ScoreFilter); ok {
		if val.ID != nil {
			id, _ := primitive.ObjectIDFromHex(*val.ID)
			filter["_id"] = id
		}
		// if val.UUID != nil {
		// 	filter["uuid"] = val.UUID
		// }
	}
	log.WithField("value", filter).Debugln("models.NewScoreFilter")
	return &filter
}
