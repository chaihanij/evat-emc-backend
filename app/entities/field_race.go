package entities

import "time"

type FieldRace struct {
	ID          string
	UUID        string
	No          int
	Title       string
	Description string
	Image       interface{}
	Document    interface{}
	FullScore   float64
	IsActive    bool
	Year        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CreatedBy   string
	UpdatedBy   string
}

type FieldRacePartialUpdate struct {
	ID          *string
	UUID        string
	No          *int
	Title       *string
	Description *string
	Image       interface{}
	Document    interface{}
	FullScore   *float64
	IsActive    *bool
	Year        *string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	CreatedBy   *string
	UpdatedBy   *string
}

type FieldRaceFilter struct {
	ID          *string
	No          *string
	UUID        *string
	Title       *string
	Description *string
	Year        *string
	Page        *int64
	PageSize    *int64
}
