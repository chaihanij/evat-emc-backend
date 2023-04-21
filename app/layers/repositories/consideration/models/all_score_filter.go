package models

import (
	log "github.com/sirupsen/logrus"

	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson"
)

func AllScoreFilter(input interface{}) bson.M {
	filter := bson.M{
		"$match": bson.M{},
	}

	if val, ok := input.(entities.AllScoreFilter); ok {
		if val.Name != "" {
			filter = bson.M{
				"$match": bson.M{
					"team": bson.M{"$regex": val.Name,
						"$options": "i",
					},
				},
			}
		}
		if val.Teamtype != "" {
			filter = bson.M{
				"$match": bson.M{
					"teamtype": bson.M{
						"$regex":   val.Teamtype,
						"$options": "i",
					},
				},
			}
		}
	}
	log.WithField("value", filter).Debugln("models.ScoreFilter")

	return filter
}
