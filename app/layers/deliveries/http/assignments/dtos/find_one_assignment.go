package dtos

import (
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

type FindOneAssignmentRequestJSON struct {
	UUID string `uri:"assignment_uuid"`
}

func (req *FindOneAssignmentRequestJSON) Parse(c *gin.Context) (*FindOneAssignmentRequestJSON, error) {
	if err := c.ShouldBindUri(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	return req, nil
}

func (req *FindOneAssignmentRequestJSON) ToEntity() *entities.AssignmentFilter {
	return &entities.AssignmentFilter{
		UUID: &req.UUID,
	}
}

type FindOneAssignmentResponseJSON AssignmentResponse

func (m *FindOneAssignmentResponseJSON) Parse(c *gin.Context, input *entities.Assignment) *FindOneAssignmentResponseJSON {
	assignment := &FindOneAssignmentResponseJSON{
		UUID:        input.UUID,
		No:          input.No,
		Title:       input.Title,
		Description: input.Description,
		FullScore:   input.FullScore,
		IsActive:    input.IsActive,
		DueDate:     input.DueDate,
		Year:        input.Year,
		CreatedAt:   input.CreatedAt,
		UpdatedAt:   input.UpdatedAt,
		CreatedBy:   input.CreatedBy,
		UpdatedBy:   input.UpdatedBy,
	}
	if val, ok := input.Document.(entities.File); ok {
		assignment.Document = new(FileResponse).Parse(c, &val)
	} else {
		assignment.Document = nil
	}
	if val, ok := input.Image.(entities.File); ok {
		assignment.Image = new(FileResponse).Parse(c, &val)
	} else {
		assignment.Image = nil
	}
	return assignment
}

type FindOneAssignmentResponseSwagger struct {
	StatusCode    int                           `json:"statusCode" example:"1000"`
	StatusMessage string                        `json:"statusMessage" example:"Success"`
	Timestamp     time.Time                     `json:"timestamp" example:"2015-06-30T21:59:59Z"`
	Data          FindOneAssignmentResponseJSON `json:"data,omitempty"`
}
