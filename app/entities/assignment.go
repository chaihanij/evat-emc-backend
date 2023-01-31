package entities

import "time"

type Assignment struct {
	ID          string
	UUID        string
	No          int
	Title       string
	Description string
	Image       interface{}
	Document    interface{}
	FullScore   float64
	IsActive    bool
	DueDate     time.Time
	Year        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CreatedBy   string
	UpdatedBy   string
}

type Assignments []Assignment

type AssignmentPartialUpdate struct {
	ID          *string
	UUID        string
	No          *int
	Title       *string
	Description *string
	Image       interface{}
	Document    interface{}
	FullScore   *float64
	IsActive    *bool
	DueDate     *time.Time
	Year        *string
	UpdatedBy   string
}

type AssignmentFilter struct {
	ID       *string
	UUID     *string
	No       *int
	Title    *string
	Year     *string
	Page     *int64
	PageSize *int64
}
