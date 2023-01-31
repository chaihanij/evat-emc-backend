package dtos

import (
	"time"

	"github.com/AlekSi/pointer"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

type FindOneUserRequestJSON struct {
	UID *string `uri:"uid" binding:"required"`
}

func (req *FindOneUserRequestJSON) Parse(c *gin.Context) (*FindOneUserRequestJSON, error) {
	if err := c.ShouldBindUri(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	return req, nil
}

func (req *FindOneUserRequestJSON) ToEntity() *entities.UserFilter {
	log.WithField("value", req).Debugln("FindOneUserRequestJSON ToEntity")
	return &entities.UserFilter{UID: req.UID}
}

type FindOneUserResponseJSON UserResponse

func (m *FindOneUserResponseJSON) Parse(data *entities.User) *FindOneUserResponseJSON {

	copier.Copy(m, data)
	if !data.LastLogin.IsZero() {
		m.LastLogin = pointer.ToTime(data.LastLogin)
	} else {
		m.LastLogin = nil
	}
	return m
}

type FindOneUserResponseSwagger struct {
	StatusCode    int                     `json:"statusCode" example:"1000"`
	StatusMessage string                  `json:"statusMessage" example:"Success"`
	Timestamp     time.Time               `json:"timestamp" example:"2015-06-30T21:59:59Z"`
	Data          FindOneUserResponseJSON `json:"data,omitempty"`
}
