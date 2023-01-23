package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TeamType struct {
	ID        primitive.ObjectID `bson:"_id"`
	UID       string             `bson:"uid"`
	Type      string             `bson:"type"`
	Content   string             `bson:"string"`
	IsActive  bool               `bson:"is_active"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`

	// Relational
	Year string `bson:"year"`
}
