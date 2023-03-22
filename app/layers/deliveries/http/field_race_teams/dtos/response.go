package dtos

import (
	"time"
)

type FieldRaceTeam struct {
	FieldRaceUUID string      `json:"field_race_uuid"`
	TeamUUID      string      `json:"team_uuid"`
	Description   string      `json:"description"`
	Score         float64     `json:"score"`
	CreatedAt     time.Time   `json:"create_at"`
	UpdatedAt     time.Time   `json:"update_at"`
	CreatedBy     string      `json:"created_by"`
	UpdatedBy     string      `json:"updated_by"`
	Name          string      `json:"name"`
	Code          string      `json:"code"`
	Type          string      `json:"type"`
	FieldRaces    []FieldRace `json:"field_races" bson:"field_races"`
}

type FieldRace struct {
	// Title       string  `json:"title"`
	// Description string  `json:"description"`
	// Image       string  `json:"image"`
	// File        string  `json:"file"`
	// Year        string  `json:"year"`
	// FullScore   float64 `json:"full_score"`
	Title       string  `json:"title" bson:"title"`
	Description string  `json:"description" bson:"description" `
	Image       string  `json:"image" bson:"image"`
	File        string  `json:"file" bson:"file" `
	Year        string  `json:"year" bson:"year" `
	FullScore   float64 `json:"full_score" bson:"full_score"`
}

type FieldRaceTeamResponse struct {
	FieldRaceUUID string      `json:"field_race_uuid"`
	TeamUUID      string      `json:"team_uuid"`
	Description   string      `json:"description"`
	Score         float64     `json:"score"`
	CreatedAt     time.Time   `json:"create_at"`
	UpdatedAt     time.Time   `json:"update_at"`
	CreatedBy     string      `json:"created_by"`
	UpdatedBy     string      `json:"updated_by"`
	Name          string      `json:"name"`
	Code          string      `json:"code"`
	Type          string      `json:"type"`
	FieldRaces    []FieldRace `json:"field_races"`
}

type MetaDataResponse struct {
	TotalRecords uint `json:"totalRecords" example:"10"`
	Page         uint `json:"page" example:"1"`
	PageSize     uint `json:"pageSize" example:"20"`
}

func (m *MetaDataResponse) Parse(page *int64, pageSize *int64, totalRecords *int64) *MetaDataResponse {
	m.TotalRecords = uint(*totalRecords)

	if page != nil {
		m.Page = uint(*page)
	}
	if pageSize != nil {
		m.PageSize = uint(*pageSize)
	}

	return m
}
