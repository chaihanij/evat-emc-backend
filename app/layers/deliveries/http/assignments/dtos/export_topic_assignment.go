package dtos

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

type FindTopicRequestJSON struct {
	UUID string `uri:"assignment_uuid"`
}

func (req *FindTopicRequestJSON) Parse(c *gin.Context) (*FindTopicRequestJSON, error) {
	if err := c.ShouldBindUri(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	return req, nil
}

func (req *FindTopicRequestJSON) ToEntity() *entities.AssignmentFilter {
	return &entities.AssignmentFilter{
		UUID: &req.UUID,
	}
}

type ExportAssignmentTopicResponseJSON struct {
	Topic []string `json:"topic" bson:"topic"`
}

func (u ExportAssignmentTopicResponseJSON) Parse(c *gin.Context, input *entities.ExportAssignmentTopic) *ExportAssignmentTopicResponseJSON {
	topic := &ExportAssignmentTopicResponseJSON{
		Topic: input.Topic,
	}
	// f := excelize.NewFile()

	for index, value := range topic.Topic {

		fmt.Println(index+1, "value :", value)
		convertRow := fmt.Sprintf("A%d", index+1)
		fmt.Println("convertRow", convertRow)
		// f.SetCellValue("Sheet1", index, value)

	}

	return topic

}
