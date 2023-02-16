package dtos

import (
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
	"gitlab.com/chaihanij/evat/app/types"
)

type CreateAnnouncementRequestJSON struct {
	Title       string `json:"title" validate:"required"  binding:"required"`
	Description string `json:"description" validate:"required"  binding:"required"`
	Year        string `json:"year" validate:"required"  binding:"required"`
	CreatedBy   string `json:"-" swaggerignore:"true"`
}

func (req *CreateAnnouncementRequestJSON) Parse(c *gin.Context) (*CreateAnnouncementRequestJSON, error) {
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

func (req *CreateAnnouncementRequestJSON) ToEntity() *entities.Announcement {
	return &entities.Announcement{
		Title:       req.Title,
		Description: req.Description,
		Year:        req.Year,
		CreatedBy:   req.CreatedBy,
	}
}

type CreateAnnouncementResponseJSON AnnouncementResponse

func (m *CreateAnnouncementResponseJSON) Parse(c *gin.Context, input *entities.Announcement) *CreateAnnouncementResponseJSON {
	assignment := &CreateAnnouncementResponseJSON{
		UUID:        input.UUID,
		Title:       input.Title,
		Description: input.Description,
		Year:        input.Year,
		CreatedAt:   input.CreatedAt,
		UpdatedAt:   input.UpdatedAt,
		CreatedBy:   input.CreatedBy,
		UpdatedBy:   input.UpdatedBy,
	}
	return assignment
}

type CreateAnnouncementResponseSwagger struct {
	StatusCode    int                            `json:"statusCode" example:"1000"`
	StatusMessage string                         `json:"statusMessage" example:"Success"`
	Timestamp     time.Time                      `json:"timestamp" example:"2015-06-30T21:59:59Z"`
	Data          CreateAnnouncementResponseJSON `json:"data,omitempty"`
}
