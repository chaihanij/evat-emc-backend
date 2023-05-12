package entities

import "time"

type Config struct {
	ID             string
	UUID           string
	RegisterConfig interface{}
	StartProject   interface{}
}

type DateRegisterConfig struct {
	Start_date time.Time
	End_date   time.Time
}

type DateStartProject struct {
	Start_date time.Time
	End_date   time.Time
}
