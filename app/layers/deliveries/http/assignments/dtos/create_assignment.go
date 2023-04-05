package dtos

import (
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
	"gitlab.com/chaihanij/evat/app/types"
)

type CreateAssignmentRequestJSON struct {
	No           int       `json:"no" validate:"required"  binding:"required"`
	Title        string    `json:"title" validate:"required"  binding:"required"`
	Description  string    `json:"description"`
	FullScore    float64   `json:"full_score"`
	IsActive     bool      `json:"isActive"`
	DueDate      time.Time `json:"dueDate" validate:"required"  binding:"required"`
	Year         string    `json:"year" validate:"required"  binding:"required"`
	CreatedBy    string    `json:"-" swaggerignore:"true"`
	TeamUUID     string    `json:"team_uuid"`
	SendDoc      bool      `json:"senddoc"`
	//OverDue      time.Time `json:"overdue" bson:"overdue"`
	DeliveryTime time.Time `json:"delivery_time" bson:"delivery_time" `

	// Consideration []ConsiderationAssignment `json:"consideration"`
}

type ConsiderationAssignment struct {
	Name  string  `json:"name"`
	Score float64 `json:"score"`
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

type ConsiderationAssignments []ConsiderationAssignment

func (req *CreateAssignmentRequestJSON) ToEntity() *entities.Assignment {

	// var ConsiderationAssignments []entities.ConsiderationAssignment

	// for _, value := range req.Consideration {
	// 	considerationAssignment := &entities.ConsiderationAssignment{
	// 		Name:  value.Name,
	// 		Score: value.Score,
	// 	}
	// 	ConsiderationAssignments = append(ConsiderationAssignments, *considerationAssignment)
	// }

	return &entities.Assignment{
		No:           req.No,
		Title:        req.Title,
		Description:  req.Description,
		FullScore:    req.FullScore,
		IsActive:     req.IsActive,
		DueDate:      req.DueDate,
		Year:         req.Year,
		CreatedBy:    req.CreatedBy,
		TeamUUID:     req.TeamUUID,
		SendDoc:      req.SendDoc,
		// OverDue:      req.OverDue,
		DeliveryTime: req.DeliveryTime,
		// Consideration: ConsiderationAssignments,
	}
}

type CreateAssignmentResponseJSON AssignmentResponse

func (m *CreateAssignmentResponseJSON) Parse(c *gin.Context, input *entities.Assignment) *CreateAssignmentResponseJSON {
	assignment := &CreateAssignmentResponseJSON{
		UUID:         input.UUID,
		No:           input.No,
		Title:        input.Title,
		Description:  input.Description,
		FullScore:    input.FullScore,
		IsActive:     input.IsActive,
		DueDate:      input.DueDate,
		Year:         input.Year,
		CreatedAt:    input.CreatedAt,
		UpdatedAt:    input.UpdatedAt,
		CreatedBy:    input.CreatedBy,
		UpdatedBy:    input.UpdatedBy,
		TeamUUID:     input.TeamUUID,
		SendDoc:      input.SendDoc,
		// OverDue:      input.OverDue,
		DeliveryTime: input.DeliveryTime,
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

type CreateAssignmentResponseSwagger struct {
	StatusCode    int                          `json:"statusCode" example:"1000"`
	StatusMessage string                       `json:"statusMessage" example:"Success"`
	Timestamp     time.Time                    `json:"timestamp" example:"2015-06-30T21:59:59Z"`
	Data          CreateAssignmentResponseJSON `json:"data,omitempty"`
}
