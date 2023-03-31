package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmailContact struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Email     string             `json:"email" bson:"email"`
	FirstName string             `json:"firstname" bson:"firstname" `
	LastName  string             `json:"lastname" bson:"lastname" `
	Create_at time.Time          `json:"create_at" bson:"create_at" `
	Status    bool               `json:"status" bson:"status" `
}

