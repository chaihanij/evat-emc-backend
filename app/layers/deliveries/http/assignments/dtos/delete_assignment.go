package dtos

import (
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

type DeleteAssignmentRequestJSON struct {
	UUID string `uri:"assignment_uuid"`
}

func (req *DeleteAssignmentRequestJSON) Parse(c *gin.Context) (*DeleteAssignmentRequestJSON, error) {
	if err := c.ShouldBindUri(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	return req, nil
}

func (req *DeleteAssignmentRequestJSON) ToEntity() *entities.AssignmentFilter {
	return &entities.AssignmentFilter{
		UUID: &req.UUID,
	}
}

type DeleteAssignmentResponseSwagger struct {
	StatusCode    int       `json:"statusCode" example:"1000"`
	StatusMessage string    `json:"statusMessage" example:"Success"`
	Timestamp     time.Time `json:"timestamp" example:"2015-06-30T21:59:59Z"`
}
