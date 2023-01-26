package dtos

import (
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
	"gitlab.com/chaihanij/evat/app/types"
)

type ResetPasswordByOTPRequestJSON struct {
	Email           string `json:"email" validate:"required,email" example:"test@gmail.com"`
	Otp             string `json:"otp" validate:"required" example:"123456"`
	NewPassword     string `json:"newPassword" validate:"required" example:"newP@ssw0rd"`
	ConfirmPassword string `json:"confirmPassword" validate:"required" example:"newP@ssw0rd"`
}

func (model *ResetPasswordByOTPRequestJSON) Parse(c *gin.Context) (*ResetPasswordByOTPRequestJSON, error) {
	if err := c.ShouldBindJSON(model); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}

	validationErr := types.Validate.Struct(model)
	if validationErr != nil {
		return nil, errors.ParameterError{Message: validationErr.Error()}
	}

	return model, nil
}

func (model *ResetPasswordByOTPRequestJSON) ToEntity() *entities.ResetPassword {
	return &entities.ResetPassword{
		Email:           model.Email,
		Otp:             model.Otp,
		NewPassword:     model.NewPassword,
		ConfirmPassword: model.ConfirmPassword,
	}
}

type ResetPasswordByOTPResponseSwagger struct {
	StatusCode    int         `json:"statusCode" example:"1000"`
	StatusMessage string      `json:"statusMessage" example:"Success"`
	Timestamp     time.Time   `json:"timestamp" example:"2015-06-30T21:59:59Z"`
	Data          interface{} `json:"data,omitempty"`
}
