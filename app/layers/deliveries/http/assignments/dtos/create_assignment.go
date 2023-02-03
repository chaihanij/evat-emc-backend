package dtos

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
	"gitlab.com/chaihanij/evat/app/types"
)

type CreateAssignmentRequestJSON struct {
	No          int       `json:"no" validate:"required"  binding:"required"`
	Title       string    `json:"title" validate:"required"  binding:"required"`
	Description string    `json:"description"`
	FullScore   float64   `json:"fullScore"`
	IsActive    bool      `json:"isActive"`
	DueDate     time.Time `json:"dueDate" validate:"required"  binding:"required"`
	Year        string    `json:"year" validate:"required"  binding:"required"`
	CreatedBy   string    `json:"-" swaggerignore:"true"`
}

func (req *CreateAssignmentRequestJSON) Parse(c *gin.Context) (*CreateAssignmentRequestJSON, error) {
	err := c.ShouldBindJSON(req)
	if err != nil {
		if errValidate := types.HandleValidateError(err, req); errValidate != nil {
			return nil, errors.ParameterError{Message: errValidate.Error()}
		}
		return nil, errors.ParameterError{Message: err.Error()}
	}
	err = types.Validate.Struct(req)
	if err != nil {
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
	req.CreatedBy = jwtData.UID

	return req, nil
}

func (req *CreateAssignmentRequestJSON) ToEntity() *entities.Assignment {
	return &entities.Assignment{
		No:          req.No,
		Title:       req.Title,
		Description: req.Description,
		FullScore:   req.FullScore,
		IsActive:    req.IsActive,
		DueDate:     req.DueDate,
		Year:        req.Year,
		CreatedBy:   req.CreatedBy,
	}
}

type CreateAssignmentResponseJSON AssignmentResponse

func (m *CreateAssignmentResponseJSON) Parse(c *gin.Context, data *entities.Assignment) *CreateAssignmentResponseJSON {
	copier.Copy(m, data)
	if val, ok := data.Document.(*entities.File); ok {
		m.Document = new(FileResponse).Parse(c, val)
	}
	if val, ok := data.Image.(*entities.File); ok {
		m.Image = new(FileResponse).Parse(c, val)
	}
	return m
}

type CreateAssignmentResponseSwagger struct {
	StatusCode    int                          `json:"statusCode" example:"1000"`
	StatusMessage string                       `json:"statusMessage" example:"Success"`
	Timestamp     time.Time                    `json:"timestamp" example:"2015-06-30T21:59:59Z"`
	Data          CreateAssignmentResponseJSON `json:"data,omitempty"`
}
