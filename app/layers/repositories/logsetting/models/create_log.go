package models

import (
	"time"

	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LogSetting struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	UUID_User   string             `json:"uuid_user" bson:"uuid_user"`
	NewData     interface{}        `json:"new_data" bson:"new_data"`
	OldData     interface{}        `json:"old_data" bson:"old_data"`
	Create_at   time.Time          `json:"create_at" bson:"create_at" `
	Discription string             `json:"discription" bson:"discription"`
}

func (u *LogSetting) ToEntity() (*entities.LogSetting, error) {
	var logsetting entities.LogSetting
	err := copier.Copy(&logsetting, u)
	return &logsetting, err
}
