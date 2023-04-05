package entities

import "time"

type AssignmentTeam struct {
	ID             string
	AssignmentUUID string
	TeamUUID       string
	Description    string
	Documents      interface{}
	IsConfirmed    bool
	Score          float64
	CreatedAt      time.Time
	UpdatedAt      time.Time
	CreatedBy      string
	UpdatedBy      string
}

type AssignmentTeamPartialUpdate struct {
	ID             *string
	AssignmentUUID *string
	TeamUUID       *string
	Description    *string
	IsConfirmed    *bool
	Documents      interface{}
	Score          *float64
	UpdatedBy      string
}

type AssignmentTeamFilter struct {
	AssignmentUUID *string
	TeamUUID       *string
	Year           *string
	Page           *int64
	PageSize       *int64
}

type AssignmentTeamScore struct {
	Title string
	Score float64
}
