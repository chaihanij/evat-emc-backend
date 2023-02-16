package dtos

import (
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
	"gitlab.com/chaihanij/evat/app/types"
)

type UpdateAnnouncementRequestJSON struct {
	UUID        string `json:"-" uri:"announcement_uuid" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Year        string `json:"year" validate:"required"`
	UpdatedBy   string `json:"-" swaggerignore:"true"`
}

func (req *UpdateAnnouncementRequestJSON) Parse(c *gin.Context) (*UpdateAnnouncementRequestJSON, error) {
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

func (req *UpdateAnnouncementRequestJSON) ToEntity() *entities.AnnouncementPartialUpdate {
	return &entities.AnnouncementPartialUpdate{
		UUID:        req.UUID,
		Title:       &req.Title,
		Description: &req.Description,
		Year:        &req.Year,
		UpdatedBy:   &req.UpdatedBy,
	}
}

type UpdateAnnouncementResponseJSON AnnouncementResponse

func (m *UpdateAnnouncementResponseJSON) Parse(c *gin.Context, input *entities.Announcement) *UpdateAnnouncementResponseJSON {
	announcement := &UpdateAnnouncementResponseJSON{
		UUID:        input.UUID,
		Title:       input.Title,
		Description: input.Description,
		Year:        input.Year,
		CreatedAt:   input.CreatedAt,
		UpdatedAt:   input.UpdatedAt,
		CreatedBy:   input.CreatedBy,
		UpdatedBy:   input.UpdatedBy,
	}
	return announcement
}

type UpdateAnnouncementResponseSwagger struct {
	StatusCode    int                            `json:"statusCode" example:"1000"`
	StatusMessage string                         `json:"statusMessage" example:"Success"`
	Timestamp     time.Time                      `json:"timestamp" example:"2015-06-30T21:59:59Z"`
	Data          UpdateAnnouncementResponseJSON `json:"data,omitempty"`
}
