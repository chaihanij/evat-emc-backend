package dtos

import (
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

type DeleteMemberRequest struct {
	UUID string `uri:"member_uuid"`
}

func (req *DeleteMemberRequest) Parse(c *gin.Context) (*DeleteMemberRequest, error) {
	if err := c.ShouldBindUri(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	return req, nil
}

func (req *DeleteMemberRequest) ToEntity() *entities.MemberFilter {
	return &entities.MemberFilter{
		UUID: &req.UUID,
	}
}

type DeleteMemberResponseSwagger struct {
	StatusCode    int       `json:"statusCode" example:"1000"`
	StatusMessage string    `json:"statusMessage" example:"Success"`
	Timestamp     time.Time `json:"timestamp" example:"2015-06-30T21:59:59Z"`
}
