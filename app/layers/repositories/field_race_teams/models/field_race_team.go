package models

import (
	"time"

	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FieldRaceTeam struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id"`
	UUID          string             `json:"uuid" bson:"uuid"`
	FieldRaceUUID string             `json:"field_race_uuid" bson:"field_race_uuid"`
	TeamUUID      string             `json:"team_uuid" bson:"team_uuid"`
	Score         float64            `score:"score" bson:"score"`
	CreatedAt     time.Time          `hson:"created_at" bson:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at" bson:"updated_at"`
	CreatedBy     string             `json:"created_by" bson:"created_by"`
	UpdatedBy     string             `json:"updated_by" bson:"updated_by"`
	Name          string             `json:"name" bson:"name"`
	Code          string             `json:"code" bson:"code" `
	Type          string             `json:"type" bson:"type" `
	FieldRaces    []FieldRaces       `json:"field_races" bson:"field_races"`
}

type CreateFieldRaceTeam struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id"`
	UUID          string             `json:"uuid" bson:"uuid"`
	FieldRaceUUID string             `json:"field_race_uuid" bson:"field_race_uuid"`
	TeamUUID      string             `json:"team_uuid" bson:"team_uuid"`
	Score         float64            `score:"score" bson:"score"`
	CreatedAt     time.Time          `hson:"created_at" bson:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at" bson:"updated_at"`
	CreatedBy     string             `json:"created_by" bson:"created_by"`
	UpdatedBy     string             `json:"updated_by" bson:"updated_by"`
	Name          string             `json:"name" bson:"name"`
	Code          string             `json:"code" bson:"code" `
	Type          string             `json:"type" bson:"type" `
	FieldRaces    []FieldRaces       `json:"field_races" bson:"field_races"`
}

type CreateFildRaceTeam struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id"`
	FieldRaceUUID string             `json:"field_race_uuid" bson:"field_race_uuid"`
	TeamUUID      string             `json:"team_uuid" bson:"team_uuid"`
	Score         float64            `score:"score" bson:"score"`
	CreatedAt     time.Time          `hson:"created_at" bson:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at" bson:"updated_at"`
	CreatedBy     string             `json:"created_by" bson:"created_by"`
	UpdatedBy     string             `json:"updated_by" bson:"updated_by"`
	FullScore     float64            `json:"full_score" bson:"full_score"`
}
type FieldRaces struct {
	Title       string  `json:"title" bson:"title"`
	Description string  `json:"description" bson:"description" `
	Image       string  `json:"image" bson:"image"`
	File        string  `json:"file" bson:"file" `
	Year        string  `json:"year" bson:"year" `
	FullScore   float64 `json:"full_score" bson:"full_score"`
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
