package dtos

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
	"gitlab.com/chaihanij/evat/app/types"
)

type CreateSuperAdminRequestJSON struct {
	Email     string `json:"email" validate:"required,email"`
	FirstName string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname" validate:"required"`
	Password  string `json:"password" validate:"required,passwordComplexity"`
}

type CreateSuperAdminResponseJSON struct {
	UID       string         `json:"uud"`
	Email     string         `json:"email"`
	FirstName string         `json:"firstname"`
	LastName  string         `json:"lastname"`
	UpdatedAt types.DateTime `json:"updatedAt,omitempty"`
	CreatedAt types.DateTime `json:"createdAt"`
}

func (req *CreateSuperAdminRequestJSON) Parse(c *gin.Context) (*CreateSuperAdminRequestJSON, error) {

	err := c.ShouldBindJSON(req)
	if err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	err = types.Validate.Struct(req)
	if err != nil {
		if strings.Contains(err.Error(), "passwordComplexity") {
			return nil, errors.ParameterError{Message: constants.WeakPassword}
		}
		return nil, errors.ParameterError{Message: err.Error()}
	}

	return req, nil
}

func (req *CreateSuperAdminRequestJSON) ToEntity() *entities.UserCreate {
	superAdmin := types.UserRoleSuperAdmin
	return &entities.UserCreate{
		Email:     req.Email,
		FirstName: &req.FirstName,
		LastName:  &req.LastName,
		Password:  &req.Password,
		Role:      &superAdmin,
	}
}

func (m *CreateSuperAdminResponseJSON) Parse(data *entities.User) *CreateSuperAdminResponseJSON {
	_ = copier.Copy(&m, data)
	return m
}

type CreateUserResponseSwagger struct {
	StatusCode    int                          `json:"statusCode" example:"1000"`
	StatusMessage string                       `json:"statusMessage" example:"Success"`
	Timestamp     time.Time                    `json:"timestamp" example:"2015-06-30T21:59:59Z"`
	Data          CreateSuperAdminResponseJSON `json:"data,omitempty"`
}
