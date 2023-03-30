package models

import (
	"time"

	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewEmail(data *entities.Email) *CreateEmail {
	return &CreateEmail{
		ID:        primitive.NewObjectID(),
		Email:     data.Email,
		Create_at: time.Now(),
	}
}
