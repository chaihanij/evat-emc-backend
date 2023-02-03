package dtos

import (
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
	"gitlab.com/chaihanij/evat/app/types"
)

type UpdateTeamRequestJSON struct {
	UUID      string `json:"-" uri:"team_uuid" validate:"required"`
	Code      string `json:"code" validate:"required"`
	Name      string `json:"name" validate:"required"`
	TeamType  string `json:"teamType" validate:"required,teamType"`
	Academy   string `json:"academy" validate:""`
	Detail    string `json:"detail" validate:""`
	Year      string `json:"year" validate:"required"`
	UpdatedBy string `json:"-" swaggerignore:"true"`
}

func (req *UpdateTeamRequestJSON) Parse(c *gin.Context) (*UpdateTeamRequestJSON, error) {

	if err := c.ShouldBindUri(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}

	if err := c.ShouldBindJSON(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	err := types.Validate.Struct(req)
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
	req.UpdatedBy = jwtData.UID
	return req, nil
}

func (req *UpdateTeamRequestJSON) ToEntity() *entities.TeamPartialUpdate {
	return &entities.TeamPartialUpdate{
		UUID:      req.UUID,
		Code:      &req.Code,
		Name:      &req.Name,
		TeamType:  &req.TeamType,
		Academy:   &req.Academy,
		Detail:    &req.Detail,
		Year:      &req.Year,
		UpdatedBy: &req.UpdatedBy,
	}
}

type UpdateTeamResponseSwagger struct {
	StatusCode    int                     `json:"statusCode" example:"1000"`
	StatusMessage string                  `json:"statusMessage" example:"Success"`
	Timestamp     time.Time               `json:"timestamp" example:"2015-06-30T21:59:59Z"`
	Data          FindOneTeamResponseJSON `json:"data,omitempty"`
}
