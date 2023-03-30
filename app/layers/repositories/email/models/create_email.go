package models

import (
	"time"

	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateEmail struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Email     string             `json:"email" bson:"email" `
	Create_at time.Time          `json:"create_at" bson:"create_at" `
}

func (at *CreateEmail) ToEntity() (*entities.Email, error ) {
	var email entities.Email
	err := copier.Copy(&email, at)
	return &email, err
}