package models

import (
	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/entities"
)

type ExportAssignmentTopic struct {
	Topic []string `json:"topic" bson:"topic" `
}

func (am *ExportAssignmentTopic) ToEntity() (*entities.ExportAssignmentTopic, error) {
	var exportTopic entities.ExportAssignmentTopic
	err := copier.Copy(&exportTopic, am)
	return &exportTopic, err
}
