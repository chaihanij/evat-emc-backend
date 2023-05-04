package entities

import "time"

type Team struct {
	ID            *string
	UUID          string
	Code          string
	Name          string
	TeamType      string
	Academy       string
	Major         string
	Detail        string
	Members       interface{} // of []string or member type
	Year          string
	Slip          interface{}
	IsPaid        bool
	PaymentMethod string
	IsVerify      bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
	CreatedBy     string
	UpdatedBy     string
}

type TeamPartialUpdate struct {
	ID            *string
	UUID          string
	Code          *string
	Name          *string
	TeamType      *string
	Academy       *string
	Major         *string
	Detail        *string
	Members       interface{}
	Slip          *string
	IsPaid        *bool
	PaymentMethod *string
	IsVerify      *bool
	Year          *string
	UpdatedBy     *string
}

type TeamFilter struct {
	ID        *string
	UUID      *string
	Year      *string
	IsVerify  *bool
	Sort      *string
	Page      *int64
	PageSize  *int64
	Name      *string
	TeamType  *string
	Tel       *string
	User_UUID *string
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
