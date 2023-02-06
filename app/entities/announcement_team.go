package entities

import "time"

type AnnouncementTeam struct {
	ID               string
	UUID             string
	TeamUUID         string
	AnnouncementUUID string
	IsRead           bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type AnnouncementTeamPartialUpdate struct {
	UUID   *string
	IsRead *bool
}

type AnnouncementTeamFilter struct {
	UUID             *string
	TeamUUID         *string
	AnnouncementUUID *string
	IsRead           *bool
}
