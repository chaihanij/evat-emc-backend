package dtos

import (
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

type DeleteUserRequestJSON struct {
	UID string `uri:"uid"`
}

func (req *DeleteUserRequestJSON) Parse(c *gin.Context) (*DeleteUserRequestJSON, error) {
	err := c.ShouldBindUri(req)
	if err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	return req, nil
}

func (req *DeleteUserRequestJSON) ToEntity() *entities.UserFilter {
	return &entities.UserFilter{UID: &req.UID}
}

type DeleteUserResponseSwagger struct {
	StatusCode    int         `json:"statusCode" example:"1000"`
	StatusMessage string      `json:"statusMessage" example:"Success"`
	Timestamp     time.Time   `json:"timestamp" example:"2015-06-30T21:59:59Z"`
	Data          interface{} `json:"data,omitempty"`
}
