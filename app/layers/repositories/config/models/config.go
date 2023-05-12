package models

import (
	"time"

	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Config struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id" `
	RegisterConfig RegisterConfig     `json:"registerconfig" bson:"registerconfig"`
	StartProject   StartProject       `json:"startproject"`
}
type StartProject struct {
	StartDate time.Time `json:"start_date" bson:"start_date"`
	EndDate   time.Time `json:"end_date" bson:"end_date" `
}

type RegisterConfig struct {
	StartDate time.Time `json:"start_date" bson:"start_date"`
	EndDate   time.Time `json:"end_date" bson:"end_date" `
}

func (t *Config) ToEntity() (*entities.Config, error) {
	var Config entities.Config
	err := copier.Copy(&Config, t)
	return &Config, err

}
