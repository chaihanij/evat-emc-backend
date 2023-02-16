package entities

import "time"

type FieldRaceTeam struct {
	ID            string
	FieldRaceUUID string
	TeamUUID      string
	Description   string
	Score         float64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	CreatedBy     string
	UpdatedBy     string
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
