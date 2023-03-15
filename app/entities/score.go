package entities

import "time"

type Score struct {
	ID             string
	// UUID           string
	NameTeam       string
	FirstTeam      string
	SecondTeam     string
	First_Stadium  string
	Second_Stadium string
	Third_Stadium  string
	Fourth_Stadium string
	Sum_Score      string
	No             string
	CreateDate     time.Time
	Lastupdate     time.Time
	CreateBy       string
	LastUpdateBy   string
}

type DropDown struct {
	ID   string
	Name string
}

type ScoreFilter struct {
	ID   *string
	// UUID *string
	// Year     *string
	Sort     *string
	Page     *int64
	PageSize *int64
}
