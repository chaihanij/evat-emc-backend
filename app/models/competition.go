package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Competition struct {
	ID   primitive.ObjectID `bson:"_id"`
	UID  string             `bson:"uid"`
	Year string             `bson:"year"`
}
