package models

import (
	"time"

	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FieldRaceTeam struct {
	ID            primitive.ObjectID `bson:"_id"`
	FieldRaceUUID string             `bson:"field_race_uuid"`
	TeamUUID      string             `bson:"team_uuid"`
	Score         float64            `bson:"score"`
	CreatedAt     time.Time          `bson:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at"`
	CreatedBy     string             `bson:"created_by"`
	UpdatedBy     string             `bson:"update_by"`
}

func (at *FieldRaceTeam) ToEntity() (*entities.FieldRaceTeam, error) {
	var fieldRaceTeam entities.FieldRaceTeam
	err := copier.Copy(&fieldRaceTeam, at)
	return &fieldRaceTeam, err
}

type FieldRaceTeams []FieldRaceTeam

func (fTeams FieldRaceTeams) ToEntity() []entities.FieldRaceTeam {
	var fieldRaceTeams []entities.FieldRaceTeam
	for _, v := range fTeams {
		fieldRaceTeam, _ := v.ToEntity()
		fieldRaceTeams = append(fieldRaceTeams, *fieldRaceTeam)
	}
	return fieldRaceTeams
}
