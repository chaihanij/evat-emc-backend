package dtos

import (
	_errors "errors"
	"fmt"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

func (req *UpdateUserProfileRequestJSON) ToEntity() *entities.UserPartialUpdate {

	return &entities.UserPartialUpdate{
		UID:       &req.UID,
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
