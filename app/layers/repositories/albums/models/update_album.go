package models

import (
	"time"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateAlbum(input *entities.Album) *bson.D {
	updateFields := bson.D{
		bson.E{Key: "title", Value: input.Title},
		bson.E{Key: "year", Value: input.Year},
		bson.E{Key: "updated_at", Value: time.Now()},
		bson.E{Key: "updated_by", Value: input.UpdatedBy},
	}
	if val, ok := input.Images.([]string); ok {
		updateFields = append(updateFields, bson.E{Key: "images", Value: val})
	}
	update := bson.D{{Key: "$set", Value: updateFields}}
	log.WithField("value", update).Debugln("models.UpdateAlbum")
	return &update
}
