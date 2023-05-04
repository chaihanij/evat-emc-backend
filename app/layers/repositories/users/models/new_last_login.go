package models

import (
	"time"

	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewLastLogin(input *entities.LastLogin) *LastLogin {
	return &LastLogin{
		ID:        primitive.NewObjectID(),
		Email: input.Email,
		IP:        input.IP,
		Create_at: time.Now(),
	}
}
