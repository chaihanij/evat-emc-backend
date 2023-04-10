package entities

import "time"

type Config struct {
	ID             string
	RegisterConfig interface{}
}

type RegisterConfig struct {
	Start_date time.Time
	End_date   time.Time
}
