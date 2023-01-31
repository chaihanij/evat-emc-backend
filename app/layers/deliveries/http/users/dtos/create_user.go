package dtos

import (
	"time"

	"github.com/AlekSi/pointer"
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
	"gitlab.com/chaihanij/evat/app/types"
)

type CreateUserRequestJSON struct {
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"email"`
	FirstName string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname" validate:"required"`
	Address   string `json:"address" validate:""`
	Tel       string `json:"tel" validate:""`
	Role      string `json:"role" validate:"userRole"`
	Year      string `json:"year" validate:"required"`
	TeamUID   string `json:"teamUID"`
	Password  string `json:"password" validate:"required,passwordComplexity"`
}

type CreateUserResponseJSON struct {
	UID       string     `json:"uid"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	FirstName string     `json:"firstname"`
	LastName  string     `json:"lastname"`
	Address   string     `json:"address"`
	Tel       string     `json:"tel"`
	Role      string     `json:"role"`
	Year      string     `json:"year"`
	TeamUID   string     `json:"teamUID"`
	IsActive  bool       `json:"isActive"`
	LastLogin *time.Time `json:"lastLogin"`
	UpdatedAt time.Time  `json:"updatedAt"`
	CreatedAt time.Time  `json:"createdAt"`
}

func (req *CreateUserRequestJSON) Parse(c *gin.Context) (*CreateUserRequestJSON, error) {

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

func (req *CreateUserRequestJSON) ToEntity() *entities.User {
	return &entities.User{
		Username:  req.Username,
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Address:   req.Address,
		Tel:       req.Tel,
		Role:      req.Role,
		Year:      req.Year,
		Password:  req.Password,
	}
}

func (m *CreateUserResponseJSON) Parse(data *entities.User) *CreateUserResponseJSON {
	var lastLogin *time.Time
	if !data.LastLogin.IsZero() {
		lastLogin = pointer.ToTime(data.LastLogin)
	} else {
		lastLogin = nil
	}
	user := &CreateUserResponseJSON{
		UID:       data.UID,
		Username:  data.Username,
		Email:     data.Email,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Role:      string(data.Role),
		Year:      data.Year,
		IsActive:  data.IsActive,
		LastLogin: lastLogin,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
	return user
}

type CreateUserResponseSwagger struct {
	StatusCode    int                    `json:"statusCode" example:"1000"`
	StatusMessage string                 `json:"statusMessage" example:"Success"`
	Timestamp     time.Time              `json:"timestamp" example:"2015-06-30T21:59:59Z"`
	Data          CreateUserResponseJSON `json:"data,omitempty"`
}
