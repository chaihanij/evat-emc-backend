package dtos

import (
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
	"gitlab.com/chaihanij/evat/app/types"
)

type UpdateAssignmentRequestJSON struct {
	UUID        string    `json:"-" uri:"assignment_uuid" validate:"required"`
	No          int       `json:"no" validate:"required"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description"`
	FullScore   float64   `json:"fullScore"`
	IsActive    bool      `json:"isActive"`
	DueDate     time.Time `json:"dueDate" validate:"required"`
	Year        string    `json:"year" validate:"required"`
	SendDoc     bool      `json:"senddoc"`
	UpdatedBy   string    `json:"-" swaggerignore:"true"`
}

func (req *UpdateAssignmentRequestJSON) Parse(c *gin.Context) (*UpdateAssignmentRequestJSON, error) {
	if err := c.ShouldBindUri(req); err != nil {
		if errValidate := types.HandleValidateError(err, req); errValidate != nil {
			return nil, errors.ParameterError{Message: errValidate.Error()}
		}
		return nil, errors.ParameterError{Message: err.Error()}
	}

	if err := c.ShouldBindJSON(req); err != nil {
		if errValidate := types.HandleValidateError(err, req); errValidate != nil {
			return nil, errors.ParameterError{Message: errValidate.Error()}
		}
		return nil, errors.ParameterError{Message: err.Error()}
	}

	if err := types.Validate.Struct(req); err != nil {
		if errValidate := types.HandleValidateError(err, req); errValidate != nil {
			return nil, errors.ParameterError{Message: errValidate.Error()}
		}
		return nil, errors.ParameterError{Message: err.Error()}
	}

	jwtRawData, ok := c.Get(constants.JWTDataKey)
	if !ok {
		return nil, errors.InternalError{Message: constants.JWTRestoreFail}
	}

	jwtData, ok := jwtRawData.(entities.JwtData)
	if !ok {
		return nil, errors.InternalError{Message: constants.JWTInvalidStructure}
	}

	if jwtData.UID == "" {
		return nil, errors.ParameterError{Message: constants.UserUIDMissing}
	}
	req.UpdatedBy = jwtData.UID

	return req, nil
}

func (req *UpdateAssignmentRequestJSON) ToEntity() *entities.AssignmentPartialUpdate {
	return &entities.AssignmentPartialUpdate{
		UUID:        req.UUID,
		Title:       &req.Title,
		Description: &req.Description,
		FullScore:   &req.FullScore,
		IsActive:    &req.IsActive,
		DueDate:     &req.DueDate,
		Year:        &req.Year,
		SendDoc:     &req.SendDoc,
		UpdatedBy:   req.UpdatedBy,
	}
}

type UpdateAssignmentResponseJSON AssignmentResponse

func (m *UpdateAssignmentResponseJSON) Parse(c *gin.Context, input *entities.Assignment) *UpdateAssignmentResponseJSON {
	assignment := &UpdateAssignmentResponseJSON{
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
		SendDoc:     input.SendDoc,
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

type UpdateAssignmentResponseSwagger struct {
	StatusCode    int                          `json:"statusCode" example:"1000"`
	StatusMessage string                       `json:"statusMessage" example:"Success"`
	Timestamp     time.Time                    `json:"timestamp" example:"2015-06-30T21:59:59Z"`
	Data          UpdateAssignmentResponseJSON `json:"data,omitempty"`
}
