package dtos

import (
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
	"gitlab.com/chaihanij/evat/app/types"
)

type ResetPasswordRequestJSON struct {
	ActivateCode    string `json:"activateCode" biding:"required"`
	Password        string `json:"password" biding:"required"  example:"newP@ssw0rd"`
	ConfirmPassword string `json:"confirmPassword" biding:"required"  example:"newP@ssw0rd"`
}

func (model *ResetPasswordRequestJSON) Parse(c *gin.Context) (*ResetPasswordRequestJSON, error) {
	if err := c.ShouldBindJSON(model); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	err := types.Validate.Struct(model)
	if err != nil {
		if errValidate := types.HandleValidateError(err, model); errValidate != nil {
			return nil, errors.ParameterError{Message: errValidate.Error()}
		}
		return nil, errors.ParameterError{Message: err.Error()}
	}
	return model, nil
}

func (model *ResetPasswordRequestJSON) ToEntity() *entities.ResetPassword {
	return &entities.ResetPassword{
		ActivateCode:    model.ActivateCode,
		NewPassword:     model.Password,
		ConfirmPassword: model.ConfirmPassword,
	}
}

type ResetPasswordSwagger struct {
	StatusCode    int         `json:"statusCode" example:"1000"`
	StatusMessage string      `json:"statusMessage" example:"Success"`
	Timestamp     time.Time   `json:"timestamp" example:"2015-06-30T21:59:59Z"`
	Data          interface{} `json:"data,omitempty"`
}
