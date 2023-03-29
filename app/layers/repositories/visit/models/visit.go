package models

import (
	"time"

	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Visit struct {
	TotalVisited int `json:"total_visited"`
	TodayVisit   int `json:"today_visit"`
}

func (at *Visit) ToEntity() (*entities.Visited, error) {
	var visit entities.Visited
	err := copier.Copy(&visit, at)
	return &visit, err
}

type CreateVisit struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	IP        string             `json:"ip" bson:"ip"`
	Create_at time.Time          `json:"create_at" bson:"create_at"`
}

func (at *CreateVisit) ToEntity() (*entities.UpdateVisit, error) {
	var visit entities.UpdateVisit
	err := copier.Copy(&visit, at)
	return &visit, err
}
