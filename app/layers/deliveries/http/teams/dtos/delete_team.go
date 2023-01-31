package dtos

import (
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

type DeleteTeamRequest struct {
	UUID string `uri:"team_uuid"`
}

func (req *DeleteTeamRequest) Parse(c *gin.Context) (*DeleteTeamRequest, error) {
	if err := c.ShouldBindUri(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	return req, nil
}

func (req *DeleteTeamRequest) ToEntity() *entities.TeamFilter {
	return &entities.TeamFilter{
		UUID: &req.UUID,
	}
}

type DeleteTeamResponseSwagger struct {
	StatusCode    int       `json:"statusCode" example:"1000"`
	StatusMessage string    `json:"statusMessage" example:"Success"`
	Timestamp     time.Time `json:"timestamp" example:"2015-06-30T21:59:59Z"`
}
