package models

import (
	log "github.com/sirupsen/logrus"

	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson"
)

func AllScoreFilter(input interface{}) bson.M {
	var filter bson.M

	if val, ok := input.(*entities.AllScoreFilter); ok {
		if val.Name != "" {
			filter = bson.M{
					"$match": bson.M{
						"name": val.Name,
					},
			}
		}
	}
	log.WithField("value", filter).Debugln("models.ScoreFilter")

	return filter
}
