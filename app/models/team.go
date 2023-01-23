package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Team struct {
	ID        primitive.ObjectID `bson:"_id"`
	UID       string             `bson:"uid"`
	Name      string             `bson:"name"`
	Code      string             `bson:"code"`
	Academy   string             `bson:"academy"`
	Detail    string             `bson:"detail"`
	IsActive  bool               `bson:"is_active"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`

	// Relational
	Year        string `bson:"year"`
	TeamTypeUID string `bson:"team_uid"`
}
