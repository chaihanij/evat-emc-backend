package models

import (
	"time"

	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LastLogin struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id" `
	Email     string             `json:"email" bson:"email" `
	IP        string             `json:"ip" bson:"ip"`
	Create_at time.Time          `json:"create_at" bson:"create_at" `
}

func (u *LastLogin) ToEntity() (*entities.LastLogin, error) {
	var lastlogin entities.LastLogin
	err := copier.Copy(&lastlogin, u)
	return &lastlogin, err
}
