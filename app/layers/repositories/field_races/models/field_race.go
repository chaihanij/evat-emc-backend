package models

import (
	"time"

	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FieldRace struct {
	ID            primitive.ObjectID       `bson:"_id"`
	UUID          string                   `bson:"uuid"`
	No            int                      `bson:"no"`
	Title         string                   `bson:"title"`
	Description   string                   `bson:"description"`
	Image         string                   `bson:"image"`
	Document      string                   `bson:"document"`
	FullScore     float64                  `bson:"full_score"`
	IsActive      bool                     `bson:"is_active"`
	Year          string                   `bson:"year"`
	CreatedAt     time.Time                `bson:"created_at"`
	UpdatedAt     time.Time                `bson:"updated_at"`
	CreatedBy     string                   `bson:"created_by"`
	UpdatedBy     string                   `bson:"updated_by"`
	Consideration []ConsiderationFieldRace `bson:"consideration"`
}

// type UploadScoreFieldRaceTeam struct {
// 	ID            primitive.ObjectID       `bson:"_id"`
// 	UUID          string                   `bson:"uuid"`
// 	No            int                      `bson:"no"`
// 	Title         string                   `bson:"title"`
// 	Description   string                   `bson:"description"`
// 	Image         string                   `bson:"image"`
// 	Document      string                   `bson:"document"`
// 	FullScore     float64                  `bson:"full_score"`
// 	IsActive      bool                     `bson:"is_active"`
// 	Year          string                   `bson:"year"`
// 	CreatedAt     time.Time                `bson:"created_at"`
// 	UpdatedAt     time.Time                `bson:"updated_at"`
// 	CreatedBy     string                   `bson:"created_by"`
// 	UpdatedBy     string                   `bson:"updated_by"`
// 	Consideration []ConsiderationFieldRace `bson:"consideration"`
// }

type ConsiderationFieldRace struct {
	ID       string  `bson:"id"`
	Title    string  `bson:"title"`
	Nameteam string  `bson:"nameteam"`
	Score    float64 `bson:"score"`
}

func (fr *FieldRace) ToEntity() (*entities.FieldRace, error) {
	var fieldRace entities.FieldRace
	err := copier.Copy(&fieldRace, fr)
	return &fieldRace, err
}

type FieldRaces []FieldRace

func (fs FieldRaces) ToEntity() []entities.FieldRace {
	var fieldRaces []entities.FieldRace
	for _, v := range fs {
		fieldRace, _ := v.ToEntity()
		fieldRaces = append(fieldRaces, *fieldRace)
	}
	return fieldRaces
}
