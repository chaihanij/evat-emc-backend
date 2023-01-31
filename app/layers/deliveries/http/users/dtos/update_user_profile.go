package dtos

import (
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
	"gitlab.com/chaihanij/evat/app/types"
)

type UpdateUserProfileRequestJSON struct {
	UID       string  `json:"uid" validate:"required"`
	Username  *string `json:"username"`
	Email     *string `json:"email"`
	FirstName *string `json:"firstname"`
	LastName  *string `json:"lastname"`
	Address   *string `json:"address"`
	Tel       *string `json:"tel"`
}

func (req *UpdateUserProfileRequestJSON) Parse(c *gin.Context) (*UpdateUserProfileRequestJSON, error) {

	if err := c.ShouldBindJSON(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	if err := types.Validate.Struct(req); err != nil {
		if err := types.HandleValidateError(err, req); err != nil {
			return nil, errors.ParameterError{Message: err.Error()}
		}
		return nil, errors.ParameterError{Message: err.Error()}
	}

	return req, nil
}

func (req *UpdateUserProfileRequestJSON) ToEntity() *entities.UserPartialUpdate {
	return &entities.UserPartialUpdate{
		UID:       req.UID,
		Username:  req.Username,
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Address:   req.Address,
		Tel:       req.Tel,
	}
}

type UpdateUserProfileResponseSwagger struct {
	StatusCode    int                            `json:"statusCode" example:"1000"`
	StatusMessage string                         `json:"statusMessage" example:"Success"`
	Timestamp     time.Time                      `json:"timestamp" example:"2015-06-30T21:59:59Z"`
	Data          FindOneUserProfileResponseJSON `json:"data,omitempty"`
}
