package dtos

import (
	_errors "errors"
	"fmt"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
	"gitlab.com/chaihanij/evat/app/types"
)

type LoginRequestJSON struct {
	Email    string `json:"email"   binding:"required,email" example:"testuser@gmail.com"`
	Password string `json:"password" binding:"required" example:"P@ssw0rd@!@#$"`
}

type LoginResponseJSON struct {
	UID         string     `json:"uid"`
	Email       string     `json:"email"`
	FirstName   string     `json:"firstname"`
	LastName    string     `json:"lastname"`
	Username    string     `json:"username"`
	Role        string     `json:"role"`
	TeamUID     string     `json:"teamUID"`
	LastLogin   *time.Time `json:"lastLogin"`
	AccessToken string     `json:"accessToken"`
}

func (req *LoginRequestJSON) Parse(c *gin.Context) (*LoginRequestJSON, error) {
	err := c.ShouldBindJSON(req)
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

func (req *LoginRequestJSON) ToEntity() *entities.Login {
	return &entities.Login{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (m *LoginResponseJSON) Parse(data *entities.User) *LoginResponseJSON {
	_ = copier.Copy(&m, data)
	m.TeamUID = data.TeamUUID
	m.LastLogin = &data.LastLogin
	return m
}

type LoginResponseSwagger struct {
	StatusCode    int               `json:"statusCode" example:"1000"`
	StatusMessage string            `json:"statusMessage" example:"Success"`
	Timestamp     time.Time         `json:"timestamp" example:"2015-06-30T21:59:59Z"`
	Data          LoginResponseJSON `json:"data,omitempty"`
}
