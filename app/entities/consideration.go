package entities

import "time"

type Consideration struct {
	ID string
	// UUID          string
	// FieldRaceUUID string
	// TeamUUID      string
	// Description   string
	Score float64
	// CreatedAt     time.Time
	UpdatedAt time.Time
	No        int
	// CreatedBy     string
	// UpdatedBy     string
}
type ConsiderationFilter struct {
	ID            *string
	FieldRaceUUID *string
	TeamUUID      *string
	Page          *int64
	PageSize      *int64
}
