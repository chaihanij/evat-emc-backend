package dtos

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
	"gitlab.com/chaihanij/evat/app/types"
)

type CreateSuperAdminRequestJSON struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (req *CreateSuperAdminRequestJSON) Parse(c *gin.Context) (*CreateSuperAdminRequestJSON, error) {

	err := c.ShouldBindJSON(req)
	if err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	err = types.Validate.Struct(req)
	if err != nil {
		if errValidate := types.HandleValidateError(err, req); errValidate != nil {
			return nil, errors.ParameterError{Message: errValidate.Error()}
		}
		return nil, errors.ParameterError{Message: err.Error()}
	}

	return req, nil
}

func (req *CreateSuperAdminRequestJSON) ToEntity() *entities.User {
	return &entities.User{
		Email:    req.Email,
		Password: req.Password,
		Role:     string(types.UserRoleSuperAdmin),
	}
}

type CreateSuperAdminResponseJSON User

func (req *CreateSuperAdminResponseJSON) Parse(data *entities.User) *CreateSuperAdminResponseJSON {
	copier.Copy(req, data)
	return req
}

type CreateSuperAdminResponseSwagger struct {
	StatusCode    int                          `json:"statusCode" example:"1000"`
	StatusMessage string                       `json:"statusMessage" example:"Success"`
	Timestamp     time.Time                    `json:"timestamp" example:"2015-06-30T21:59:59Z"`
	Data          CreateSuperAdminResponseJSON `json:"data,omitempty"`
}
