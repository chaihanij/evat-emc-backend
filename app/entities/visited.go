package entities

import "time"

type Visited struct {
	ID           string
	TotalVisited int
	TodayVisit   int
}

type UpdateVisit struct {
	ID        string
	IP        string
	Create_at time.Time
}
