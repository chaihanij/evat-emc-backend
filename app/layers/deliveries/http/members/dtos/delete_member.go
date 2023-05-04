package dtos

import (
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

type DeleteMemberRequest struct {
	UUID      string `uri:"member_uuid"`
	User_UUID string `json:"user_uuid"`
}

func (req *DeleteMemberRequest) Parse(c *gin.Context) (*DeleteMemberRequest, error) {
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

func (req *DeleteMemberRequest) ToEntity() *entities.MemberFilter {
	return &entities.MemberFilter{
		UUID:      &req.UUID,
		User_UUID: &req.User_UUID,
	}
}

type DeleteMemberResponseSwagger struct {
	StatusCode    int       `json:"statusCode" example:"1000"`
	StatusMessage string    `json:"statusMessage" example:"Success"`
	Timestamp     time.Time `json:"timestamp" example:"2015-06-30T21:59:59Z"`
}
