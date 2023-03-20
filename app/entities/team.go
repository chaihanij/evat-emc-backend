package entities

import "time"

type Team struct {
	ID        *string
	UUID      string
	Code      string
	Name      string
	TeamType  string
	Academy   string
	Detail    string
	Members   interface{} // of []string or member type
	Year      string
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy string
	UpdatedBy string
}

type TeamPartialUpdate struct {
	ID        *string
	UUID      string
	Code      *string
	Name      *string
	TeamType  *string
	Academy   *string
	Detail    *string
	Members   interface{}
	Year      *string
	UpdatedBy *string
}

type TeamFilter struct {
	ID       *string
	UUID     *string
	Year     *string
	Sort     *string
	Page     *int64
	PageSize *int64
}

type TeamSearch struct {
	ID       string
	UUID     string
	Code     string
	Name     string
	Academy  string
	Tel      string
	Contact  string
	TeamType string
}
