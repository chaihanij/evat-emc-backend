package models

import (
	"time"

	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewCreateLogSetting(input *entities.LogSetting) *LogSetting {

	return &LogSetting{
		ID:          primitive.NewObjectID(),
		UUID_User:   input.UUID_User,
		NewData:     input.NewData,
		OldData:     input.OldData,
		Create_at:   time.Now(),
		Discription: input.Discription,
	}

}
