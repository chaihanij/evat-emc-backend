package entities

import "time"

type Consideration struct {
	ID string
	// // UUID          string
	// // FieldRaceUUID string
	// // TeamUUID      string
	// // Description   string
	// Score float64
	// // CreatedAt     time.Time
	UpdatedAt time.Time
	No        int
	// // CreatedBy     string
	// UpdatedBy     string
	TotalScore     float64
	IndivdualScore []IndivdualScore
}

type IndivdualScore struct {
	Title  string
	Score float64
}
type ConsiderationFilter struct {
	ID            *string
	FieldRaceUUID *string
	TeamUUID      *string
	Page          *int64
	PageSize      *int64
}
