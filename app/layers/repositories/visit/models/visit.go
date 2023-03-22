package models

import (
	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Visit struct {
	ID           primitive.ObjectID `json:"_id"`
	TotalVisited int                `json:"total_visited"`
	TodayVisit   int                `json:"today_visit"`
}

func (at *Visit) ToEntity() (*entities.Visited, error) {
	var visit entities.Visited
	err := copier.Copy(&visit, at)
	return &visit, err
}
