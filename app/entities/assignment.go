package entities

import "time"

type Assignment struct {
	ID            string
	UUID          string
	TeamUUID      string
	No            int
	Title         string
	Description   string
	Image         interface{}
	Document      interface{}
	FullScore     float64
	IsActive      bool
	DueDate       time.Time
	Year          string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	CreatedBy     string
	UpdatedBy     string
	SendDoc       bool
	Consideration []ConsiderationAssignment
}

type ConsiderationAssignment struct {
	Name  string
	Score float64
}

type TeamAssignment struct {
	ID        string
	UUID      string
	TeamUUID  string
	Title     string
	FullScore float64
	CreatedAt time.Time
}

type Assignments []Assignment

type AssignmentPartialUpdate struct {
	ID            *string
	UUID          string
	No            *int
	Title         *string
	Description   *string
	Image         interface{}
	Document      interface{}
	FullScore     *float64
	IsActive      *bool
	DueDate       *time.Time
	Year          *string
	UpdatedBy     string
	SendDoc       *bool
	Consideration *[]interface{}
}

type AssignmentFilter struct {
	ID       *string
	UUID     *string
	TeamUUID *string
	No       *int
	Title    *string
	Year     *string
	Page     *int64
	PageSize *int64
}
