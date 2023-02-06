package entities

import "time"

type Announcement struct {
	ID          string
	UUID        string
	Title       string
	Description string
	Year        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CreatedBy   string
	UpdatedBy   string
}

type AnnouncementPartialUpdate struct {
	ID          *string
	UUID        string
	Title       *string
	Description *string
	Year        *string
	UpdatedBy   *string
}

type AnnouncementFilter struct {
	ID          *string
	UUID        *string
	Title       *string
	Description *string
	Year        *string
	Page        *int64
	PageSize    *int64
}
