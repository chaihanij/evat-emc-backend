package entities

import "time"

type Assignment struct {
	ID            string
	UUID          string
	IsShowMenu    bool
	TeamUUID      string
	No            int
	Title         string
	Description   string
	Image         interface{}
	Document      interface{}
	FullScore     float64
	IsActive      bool
	DueDate       time.Time
	Year          string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	CreatedBy     string
	UpdatedBy     string
	SendDoc       bool
	Consideration []ConsiderationAssignment
	// OverDue      time.Time
	DeliveryTime   time.Time
	Fileassignment []Fileassignment
	Topic          []string
}

type Fileassignment struct {
	// UUID       *string
	FileName   string
	CreateDate time.Time
	FileUrl    string
	CreateBy   string
}

type ConsiderationAssignment struct {
	ID       string
	Title    string
	NameTeam string
	TeamType string
	// No       int
	Score float64
}

type TeamAssignment struct {
	ID        string
	UUID      string
	TeamUUID  string
	Title     string
	FullScore float64
	CreatedAt time.Time
}

type Assignments []Assignment

type AssignmentPartialUpdate struct {
	ID            *string
	UUID          string
	No            *int
	Title         *string
	Description   *string
	Image         interface{}
	Document      interface{}
	FullScore     *float64
	IsActive      *bool
	DueDate       *time.Time
	Year          *string
	UpdatedBy     string
	SendDoc       *bool
	Consideration *[]interface{}
	DeliveryTime  *time.Time
	IsShowMenu    *bool
	// UploadFile    []UploadFile
}

type AssignmentFilter struct {
	ID        *string
	UUID      *string
	User_UUID *string
	TeamUUID  *string
	No        *int
	Title     *string
	Year      *string
	Page      *int64
	PageSize  *int64
}

type ExportAssignmentTopic struct {
	Title           string
	Topic           []string
	Path            string
	ExportTeamTopic []ExportTeamTopic
}

type ExportTeamTopic struct {
	Code      string
	Name      string
	Team_type string
}

type AssignmentPartialUpdateScore struct {
	AssignmentUUID string
	UploadFile     Fileassignment
}
