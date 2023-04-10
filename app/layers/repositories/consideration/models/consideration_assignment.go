package models

import (
	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/entities"
)

type AssignmentScore struct {
	// ID             string          `json:"_id" bson:"_id"`
	Considerations []Consideration `json:"considerations" bson:"considerations" `
	Total          float64         `json:"total" bson:"total"`
}

type Consideration struct {
	ID       string  `json:"id" bson:"id" `
	Title    string  `json:"title" bson:"title"`
	TeamName string  `json:"nameteam" bson:"nameteam"`
	Score    float64 `json:"score" bson:"score"`
}

type AssignmentScores []AssignmentScore

func (am *AssignmentScore) ToEntity() (*entities.AssignmentScore, error) {
	var assignmentScore entities.AssignmentScore
	err := copier.Copy(&assignmentScore, am)
	return &assignmentScore, err
}

func (as AssignmentScores) ToEntity() []entities.AssignmentScore {
	var assignmentScores []entities.AssignmentScore
	for _, v := range as {
		assignment, _ := v.ToEntity()
		assignmentScores = append(assignmentScores, *assignment)
	}
	return assignmentScores
}
