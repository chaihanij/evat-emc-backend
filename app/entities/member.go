package entities

import "time"

type Member struct {
	ID           string
	UUID         string
	FirstName    string
	LastName     string
	Address      string
	Email        string
	Tel          string
	Academy      string
	Major        string
	Image        interface{} // string or file
	Year         string
	MemberType   string
	Documents    interface{} // []string or []file
	IsTeamLeader bool
	TeamUUID     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	CreatedBy    string
	UpdatedBy    string
}

type Members []Member

type MemberPartialUpdate struct {
	ID           *string
	UUID         string
	FirstName    *string
	LastName     *string
	Address      *string
	Email        *string
	Tel          *string
	Academy      *string
	Major        *string
	Image        interface{} // string or file
	Year         *string
	MemberType   *string
	Documents    interface{} // []string or []file
	IsTeamLeader *bool
	TeamUUID     *string
	CreatedBy    *string
	UpdatedBy    *string
}

type MemberFilter struct {
	ID        *string
	UUID      *string
	Year      *string
	TeamUUID  *string
	Sort      *string
	Page      *int64
	PageSize  *int64
	User_UUID *string
}
