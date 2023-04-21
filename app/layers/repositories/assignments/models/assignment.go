package models

import (
	"time"

	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Assignment struct {
	ID            primitive.ObjectID        `bson:"id"`
	UUID          string                    `bson:"uuid"`
	TeamUUID      string                    `bson:"team_uuid"`
	No            int                       `bson:"no"`
	Title         string                    `bson:"title"`
	Description   string                    `bson:"description"`
	Image         string                    `bson:"image"`
	Document      string                    `bson:"document"`
	FullScore     float64                   `bson:"full_score"`
	IsActive      bool                      `bson:"is_active"`
	DueDate       time.Time                 `bson:"due_date"`
	Year          string                    `bson:"year"`
	CreatedAt     time.Time                 `bson:"created_at"`
	UpdatedAt     time.Time                 `bson:"updated_at"`
	CreatedBy     string                    `bson:"created_by"`
	UpdatedBy     string                    `bson:"updated_by"`
	SendDoc       bool                      `bson:"senddoc"`
	Consideration []ConsiderationAssignment `bson:"consideration"`
	DeliveryTime  time.Time                 `bson:"delivery_time"`
	IsShowMenu    bool                      `bson:"isShowMenu"`
}

type ConsiderationAssignment struct {
	ID        string  `bson:"id"`
	Title     string  `bson:"title"`
	Nameteam  string  `bson:"nameteam"`
	Team_type string  `bson:"team_type"`
	Score     float64 `bson:"score"`
	No        int     `bson:"no"`
}

func (am *Assignment) ToEntity() (*entities.Assignment, error) {
	var assignment entities.Assignment
	err := copier.Copy(&assignment, am)
	return &assignment, err
}

type Assignments []Assignment

func (as Assignments) ToEntity() []entities.Assignment {
	var assignments []entities.Assignment
	for _, v := range as {
		assignment, _ := v.ToEntity()
		assignments = append(assignments, *assignment)
	}
	return assignments
}

type TeamAssignment struct {
	ID        primitive.ObjectID `bson:"id"`
	TeamUUID  string             `bson:"team_uuid"`
	UUID      string             `bson:"uuid"`
	Title     string             `bson:"title"`
	FullScore float64            `bson:"full_score"`
}
type TeamAssignments []TeamAssignment

func (am *TeamAssignment) ToEntity() (*entities.TeamAssignment, error) {
	var teamassignment entities.TeamAssignment
	err := copier.Copy(&teamassignment, am)
	return &teamassignment, err
}

func (as TeamAssignments) ToEntity() []entities.TeamAssignment {
	var teamassignments []entities.TeamAssignment
	for _, v := range as {
		teamassignment, _ := v.ToEntity()
		teamassignments = append(teamassignments, *teamassignment)
	}
	return teamassignments
}
