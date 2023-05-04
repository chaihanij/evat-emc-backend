package entities

import "time"

type Login struct {
	Email    string
	Password string
}

type LastLogin struct {
	ID        string
	Email string
	IP        string
	create_at time.Time
}
