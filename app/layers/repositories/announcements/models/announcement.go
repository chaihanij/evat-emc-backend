package models

import (
	"time"

	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Announcement struct {
	ID          primitive.ObjectID `bson:"_id"`
	UUID        string             `bson:"uuid"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Year        string             `bson:"year"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
	CreatedBy   string             `bson:"created_by"`
	UpdatedBy   string             `bson:"update_by"`
}

func (am *Announcement) ToEntity() (*entities.Announcement, error) {
	var announcement entities.Announcement
	err := copier.Copy(&announcement, am)
	return &announcement, err
}

type Announcements []Announcement

func (as Announcements) ToEntity() []entities.Announcement {
	var announcements []entities.Announcement
	for _, v := range as {
		announcement, _ := v.ToEntity()
		announcements = append(announcements, *announcement)
	}
	return announcements
}
