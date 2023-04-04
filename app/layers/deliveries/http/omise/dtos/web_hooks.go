package dtos

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/omise/omise-go"
	"gitlab.com/chaihanij/evat/app/errors"
)

type WebHooksRequestJSON omise.Event

func (req *WebHooksRequestJSON) Parse(c *gin.Context) (*WebHooksRequestJSON, error) {

	err := c.ShouldBindJSON(req)
	if err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}

	return req, nil
}

func (req *WebHooksRequestJSON) ToEntity() *omise.Event {
	var event omise.Event
	_ = copier.Copy(&event, req)
	return &event
}
