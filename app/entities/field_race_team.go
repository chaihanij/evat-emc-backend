package entities

import "time"

type FieldRaceTeam struct {
	ID            string
	UUID          string
	FieldRaceUUID string
	TeamUUID      string
	Description   string
	Score         float64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	CreatedBy     string
	UpdatedBy     string
	Name          string
	Code          string
	Type          string
	FieldRaces    []FieldRaces
}

type CreateFildRaceTeam struct {
	ID            string
	UUID          string
	FieldRaceUUID string
	TeamUUID      string
	Description   string
	Score         float64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	CreatedBy     string
	UpdatedBy     string
	FieldRaces    []FieldRaces
}

type FieldRaces struct {
	Title       string
	Description string
	Image       string
	File        string
	Year        string
	FullScore   float64
}

type FieldRaceTeamPartialUpdate struct {
	ID            *string
	FieldRaceUUID *string
	TeamUUID      *string
	Description   *string
	Score         *float64
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
	CreatedBy     *string
	UpdatedBy     *string
}

type FieldRaceTeamFilter struct {
	ID            *string
	FieldRaceUUID *string
	TeamUUID      *string
	Page          *int64
	PageSize      *int64
}
