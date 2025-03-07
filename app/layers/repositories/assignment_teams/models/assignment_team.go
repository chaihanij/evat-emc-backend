package models

import (
	"time"

	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AssignmentTeam struct {
	ID             primitive.ObjectID `bson:"_id"`
	AssignmentUUID string             `bson:"assignment_uuid"`
	TeamUUID       string             `bson:"team_uuid"`
	Description    string             `bson:"description"`
	// Documents       []string           `bson:"documenta"`
	IsConfirmed     bool       `bson:"is_confirmed"`
	Score           float64    `bson:"score"`
	CreatedAt       time.Time  `bson:"created_at"`
	UpdatedAt       time.Time  `bson:"updated_at"`
	CreatedBy       string     `bson:"created_by"`
	UpdatedBy       string     `bson:"update_by"`
	AssignmentTopic string     `bson:"assignment_topic"`
	Document        []Document `bson:"documents"`
}

type Document struct {
	DocumentUUID    string `bson:"documentUUID"`
	AssignmentTopic string `bson:"AssignmentTopic"`
}

type DliveryTime struct {
	DliveryTime time.Time `json:"delivery_time" bson:"delivery_time"`
	DueDate     time.Time `json:"due_date" bson:"due_date"`
	TimeNow     time.Time
}

type DocumentAssignment struct {
	FileUUID        string `json:"documentUUID" bson:"documentUUID" `
	AssignmentTopic string `json:"AssignmentTopic" bson:"AssignmentTopic"`
}

func (at *AssignmentTeam) ToEntity() (*entities.AssignmentTeam, error) {
	var assignmentTeam entities.AssignmentTeam
	err := copier.Copy(&assignmentTeam, at)
	return &assignmentTeam, err
}

type AssignmentTeams []AssignmentTeam

func (aTeams AssignmentTeams) ToEntity() []entities.AssignmentTeam {
	var assignmentTeams []entities.AssignmentTeam
	for _, v := range aTeams {
		assignmentTeam, _ := v.ToEntity()
		assignmentTeams = append(assignmentTeams, *assignmentTeam)
	}
	return assignmentTeams
}

type AssignmentTeamScore struct {
	Title string  `bson:"title"`
	Score float64 `bson:"full_score"`
}

func (at *AssignmentTeamScore) ToEntity() (*entities.AssignmentTeamScore, error) {
	var assignmentTeamScore entities.AssignmentTeamScore
	err := copier.Copy(&assignmentTeamScore, at)
	return &assignmentTeamScore, err
}

type AssignmentTeamScores []AssignmentTeamScore

func (aTeams AssignmentTeamScores) ToEntity() []entities.AssignmentTeamScore {
	var assignmentTeamScores []entities.AssignmentTeamScore
	for _, v := range aTeams {
		assignmentTeam, _ := v.ToEntity()
		assignmentTeamScores = append(assignmentTeamScores, *assignmentTeam)
	}
	return assignmentTeamScores
}
