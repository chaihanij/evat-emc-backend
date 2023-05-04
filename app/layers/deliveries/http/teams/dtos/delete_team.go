package dtos

import (
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

type DeleteTeamRequest struct {
	UUID      string `uri:"team_uuid"`
	User_UUID string `json:"user_uuid"`
}

func (req *DeleteTeamRequest) Parse(c *gin.Context) (*DeleteTeamRequest, error) {
	if err := c.ShouldBindUri(req); err != nil {
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

	req.User_UUID = jwtData.UID

	return req, nil
}

func (req *DeleteTeamRequest) ToEntity() *entities.TeamFilter {
	return &entities.TeamFilter{
		UUID:      &req.UUID,
		User_UUID: &req.User_UUID,
	}
}

type DeleteTeamResponseSwagger struct {
	StatusCode    int       `json:"statusCode" example:"1000"`
	StatusMessage string    `json:"statusMessage" example:"Success"`
	Timestamp     time.Time `json:"timestamp" example:"2015-06-30T21:59:59Z"`
}
