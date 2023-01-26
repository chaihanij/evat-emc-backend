package dtos

import (
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

type ForgotPasswordRequestJSON struct {
	Email string `json:"email" biding:"required"  example:"test@gmail.com"`
}

func (model *ForgotPasswordRequestJSON) Parse(c *gin.Context) (*ForgotPasswordRequestJSON, error) {
	if err := c.ShouldBindJSON(model); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}

	return model, nil
}

func (model *ForgotPasswordRequestJSON) ToEntity() *entities.ResetPassword {
	return &entities.ResetPassword{
		Email: model.Email,
	}
}

type ForgotPasswordResponseSwagger struct {
	StatusCode    int         `json:"statusCode" example:"1000"`
	StatusMessage string      `json:"statusMessage" example:"Success"`
	Timestamp     time.Time   `json:"timestamp" example:"2015-06-30T21:59:59Z"`
	Data          interface{} `json:"data,omitempty"`
}
