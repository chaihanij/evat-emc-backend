package entities

import "time"

type Score struct {
	ID              string
	UUID            string
	Field_race_uuid string
	Team_uuid       string
	Score           float64
	CreateDate      time.Time
	Lastupdate      time.Time
	CreateBy        string
	Updated_by      string
}

type DropDown struct {
	ID   string
	Name string
}

type ScoreFilter struct {
	ID *string
	// UUID *string
	// Year     *string
	Sort     *string
	Page     *int64
	PageSize *int64
}
