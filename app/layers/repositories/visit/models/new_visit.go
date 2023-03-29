package models

import (
	"time"

	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewVisit(data *entities.UpdateVisit) *CreateVisit {
	return &CreateVisit{
		ID:        primitive.NewObjectID(),
		IP:        data.IP,
		Create_at: time.Now(),
	}
}
