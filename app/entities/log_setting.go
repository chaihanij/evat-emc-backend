package entities

import "time"

type LogSetting struct {
	ID          string
	UUID_User   string
	NewData     interface{}
	OldData     interface{}
	Discription string
	Create_at   time.Time
}
