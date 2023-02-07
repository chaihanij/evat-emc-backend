package dtos

import (
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

type SendAssignmentTeamRequestJSON struct {
	TeamUUID       string `json:"-" uri:"team_uuid"`
	AssignmentUUID string `json:"-" uri:"assignment_uuid"`
	Description    string `json:"description"`
	IsConfirmed    bool   `json:"isConfirmed"`
	UpdatedBy      string `json:"-"`
}

func (req *SendAssignmentTeamRequestJSON) Parse(c *gin.Context) (*SendAssignmentTeamRequestJSON, error) {
	if err := c.ShouldBindUri(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	if err := c.ShouldBindJSON(req); err != nil {
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

func (req *SendAssignmentTeamRequestJSON) ToEntity() *entities.AssignmentTeam {
	return &entities.AssignmentTeam{
		TeamUUID:       req.TeamUUID,
		AssignmentUUID: req.AssignmentUUID,
		Description:    req.Description,
		IsConfirmed:    req.IsConfirmed,
		UpdatedBy:      req.UpdatedBy,
	}
}

type SendAssignmentTeamResponseJSON struct {
	TeamUUID       string `json:"team_uuid"`
	AssignmentUUID string `json:"assignment_uuid"`
	Description    string `json:"description"`
	IsConfirmed    bool   `json:"isConfirmed"`
}

func (res *SendAssignmentTeamResponseJSON) Parse(input *entities.AssignmentTeam) *SendAssignmentTeamResponseJSON {
	return &SendAssignmentTeamResponseJSON{
		TeamUUID:       input.TeamUUID,
		AssignmentUUID: input.AssignmentUUID,
		Description:    input.Description,
		IsConfirmed:    input.IsConfirmed,
	}
}

type SendAssignmentTeamResponseSwagger struct {
	StatusCode    int                            `json:"statusCode" example:"1000"`
	StatusMessage string                         `json:"statusMessage" example:"Success"`
	Timestamp     time.Time                      `json:"timestamp" example:"2015-06-30T21:59:59Z"`
	Data          SendAssignmentTeamResponseJSON `json:"data,omitempty"`
}
