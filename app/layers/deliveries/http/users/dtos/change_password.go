package dtos

import (
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

type ChangePasswordRequestJSON struct {
	UID             string `json:"-"`
	OldPassword     string `json:"oldPassword" biding:"required"  example:"oldP@ssw0rd"`
	NewPassword     string `json:"newPassword" biding:"required"  example:"newP@ssw0rd"`
	ConfirmPassword string `json:"confirmPassword" biding:"required"  example:"newP@ssw0rd"`
}

func (model *ChangePasswordRequestJSON) Parse(c *gin.Context) (*ChangePasswordRequestJSON, error) {
	if err := c.ShouldBindJSON(model); err != nil {
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

	model.UID = jwtData.UID

	return model, nil

}

func (model *ChangePasswordRequestJSON) ToEntity() *entities.ResetPassword {
	return &entities.ResetPassword{
		UID:             model.UID,
		OldPassword:     model.OldPassword,
		NewPassword:     model.NewPassword,
		ConfirmPassword: model.ConfirmPassword,
	}
}

type ChangePasswordResponseSwagger struct {
	StatusCode    int         `json:"statusCode" example:"1000"`
	StatusMessage string      `json:"statusMessage" example:"Success"`
	Timestamp     time.Time   `json:"timestamp" example:"2015-06-30T21:59:59Z"`
	Data          interface{} `json:"data,omitempty"`
}
