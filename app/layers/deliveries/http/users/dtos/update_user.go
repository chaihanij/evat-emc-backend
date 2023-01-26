package dtos

import (
	_errors "errors"
	"fmt"
	"reflect"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
	"gitlab.com/chaihanij/evat/app/types"
)

type UpdateUserRequestJSON struct {
	UID       string         `json:"uid" validate:"required"`
	Username  string         `json:"username" validate:"required"`
	Email     string         `json:"email" validate:""`
	FirstName string         `json:"firstname" validate:""`
	LastName  string         `json:"lastname" validate:""`
	Address   string         `json:"address" validate:""`
	Tel       string         `json:"tel" validate:""`
	Role      types.UserRole `json:"role" validate:"userRole"`
	Year      string         `json:"year" validate:""`
	TeamUID   string         `json:"teamUID" validate:""`
	IsActive  bool           `json:"isActive" validate:""`
}

type UpdateUserResponseJSON struct {
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

func (req *UpdateUserRequestJSON) Parse(c *gin.Context) (*UpdateUserRequestJSON, error) {

	err := c.ShouldBindJSON(req)
	if err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	err = types.Validate.Struct(req)
	if err != nil {
		var ve validator.ValidationErrors
		if _errors.As(err, &ve) {
			for _, fe := range ve {
				fieldName := fe.Field()
				field, ok := reflect.TypeOf(req).Elem().FieldByName(fieldName)
				if ok {
					fieldName, ok := field.Tag.Lookup("json")
					if ok {
						msg := fmt.Sprintf("%s %s", fieldName, types.MsgForTag(fe))
						return nil, errors.ParameterError{Message: msg}
					}
				} else {
					msg := fmt.Sprintf("%s %s", fieldName, types.MsgForTag(fe))
					return nil, errors.ParameterError{Message: msg}
				}
			}
		}
		return nil, errors.ParameterError{Message: err.Error()}
	}

	return req, nil
}

func (req *UpdateUserRequestJSON) ToEntity() *entities.UserPartialUpdate {

	return &entities.UserPartialUpdate{
		UID:       &req.UID,
		Username:  &req.Username,
		Email:     &req.Email,
		FirstName: &req.FirstName,
		LastName:  &req.LastName,
		Address:   &req.Address,
		Tel:       &req.Tel,
		Role:      &req.Role,
		Year:      &req.Year,
		TeamUID:   &req.TeamUID,
		IsActive:  &req.IsActive,
	}
}

func (m *UpdateUserResponseJSON) Parse(data *entities.User) *UpdateUserResponseJSON {
	var lastLogin *time.Time
	if !data.LastLogin.IsZero() {
		lastLogin = pointer.ToTime(data.LastLogin)
	} else {
		lastLogin = nil
	}
	return &UpdateUserResponseJSON{
		UID:       data.UID,
		Email:     data.Email,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Role:      string(data.Role),
		Address:   data.Address,
		Tel:       data.Tel,
		Year:      data.Year,
		TeamUID:   data.TeamUID,
		IsActive:  data.IsActive,
		LastLogin: lastLogin,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

type UpdateUserResponseSwagger struct {
	StatusCode    int                    `json:"statusCode" example:"1000"`
	StatusMessage string                 `json:"statusMessage" example:"Success"`
	Timestamp     time.Time              `json:"timestamp" example:"2015-06-30T21:59:59Z"`
	Data          UpdateUserResponseJSON `json:"data,omitempty"`
}
