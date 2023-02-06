package models

import (
	"time"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson"
)

func PartialUpdateAlbum(input *entities.AlbumPartialUpdate) *bson.D {
	updateFields := bson.D{
		bson.E{Key: "updated_at", Value: time.Now()},
		bson.E{Key: "updated_by", Value: input.UpdatedBy},
	}
	if input.Title != nil {
		updateFields = append(updateFields, bson.E{Key: "title", Value: input.Title})
	}
	if input.Year != nil {
		updateFields = append(updateFields, bson.E{Key: "year", Value: input.Year})
	}
	if val, ok := input.Images.([]string); ok {
		updateFields = append(updateFields, bson.E{Key: "images", Value: val})
	}
	update := bson.D{{Key: "$set", Value: updateFields}}
	log.WithField("value", update).Debugln("models.UpdateAlbum")
	return &update
}
