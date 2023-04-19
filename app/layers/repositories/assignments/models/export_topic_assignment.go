package models

import (
	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/entities"
)

type ExportAssignmentTopic struct {
	Title           string   `json:"title" bson:"title"`
	Topic           []string `json:"topic" bson:"topic" `
	ExportTeamTopic []ExportTeamTopic
}

type ExportTeamTopic struct {
	Code      string `json:"code" bson:"code" `
	Name      string `json:"name" bson:"name"`
	Team_type string `json:"team_type" bson:"team_type" `
}

func (am *ExportAssignmentTopic) ToEntity() (*entities.ExportAssignmentTopic, error) {
	var exportTopic entities.ExportAssignmentTopic
	err := copier.Copy(&exportTopic, am)
	return &exportTopic, err
}
