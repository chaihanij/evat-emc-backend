package models

import (
	"time"

	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Album struct {
	ID        primitive.ObjectID `bson:"_id"`
	UUID      string             `bson:"uuid"`
	Title     string             `bson:"title"`
	Images    []string           `bson:"images"`
	Year      string             `bson:"year"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	CreatedBy string             `bson:"created_by"`
	UpdatedBy string             `bson:"updated_by"`
}

func (a *Album) ToEntity() (*entities.Album, error) {
	var album entities.Album
	err := copier.Copy(&album, a)
	return &album, err
}

type Albums []Album

func (as Albums) ToEntity() []entities.Album {
	var albums []entities.Album
	for _, v := range as {
		album, _ := v.ToEntity()
		albums = append(albums, *album)
	}
	return albums
}
