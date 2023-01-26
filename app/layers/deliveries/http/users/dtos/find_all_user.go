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

type MetaDataResponse struct {
	TotalRecords uint `json:"totalRecords" example:"10"`
	Page         uint `json:"page" example:"1"`
	PageSize     uint `json:"pageSize" example:"20"`
}

func (m *MetaDataResponse) Parse(page *int64, pageSize *int64, totalRecords *int64) *MetaDataResponse {
	m.TotalRecords = uint(*totalRecords)

	if page != nil {
		m.Page = uint(*page)
	}
	if pageSize != nil {
		m.PageSize = uint(*pageSize)
	}

	return m
}

type FindAllUserRequestJSON struct {
	Year     *string `form:"year" example:"2023"`
	Page     *int64  `form:"page" validate:"omitempty,gte=1" example:"1"`
	PageSize *int64  `form:"pageSize" validate:"omitempty,gte=1" example:"20"`
}

type User struct {
	UID       string     `json:"uid"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	FirstName string     `json:"firstname"`
	LastName  string     `json:"lastname"`
	Address   string     `json:"address"`
	Role      string     `json:"role"`
	Year      string     `json:"year"`
	IsActive  bool       `json:"isActive"`
	LastLogin *time.Time `json:"lastLogin"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}

type FindAllUserResponseJSON []User

func (req *FindAllUserRequestJSON) Parse(c *gin.Context) (*FindAllUserRequestJSON, error) {

	err := c.ShouldBindQuery(req)
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

func (req *FindAllUserRequestJSON) ToEntity() *entities.UserFilter {
	return &entities.UserFilter{
		Year:     req.Year,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
}

func (m *FindAllUserResponseJSON) Parse(data []entities.User) *FindAllUserResponseJSON {
	var users FindAllUserResponseJSON
	for _, value := range data {
		var lastLogin *time.Time
		if !value.LastLogin.IsZero() {
			lastLogin = pointer.ToTime(value.LastLogin)
		} else {
			lastLogin = nil
		}
		user := &User{
			UID:       value.UID,
			Username:  value.Username,
			Email:     value.Email,
			FirstName: value.FirstName,
			LastName:  value.LastName,
			Role:      string(value.Role),
			Year:      value.Year,
			IsActive:  value.IsActive,
			LastLogin: lastLogin,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		}
		users = append(users, *user)
	}
	return &users
}

type FindAllUserResponseSwagger struct {
	StatusCode    int                     `json:"statusCode" example:"1000"`
	StatusMessage string                  `json:"statusMessage" example:"Success"`
	Timestamp     time.Time               `json:"timestamp" example:"2015-06-30T21:59:59Z"`
	Data          FindAllUserResponseJSON `json:"data,omitempty"`
	MetaData      MetaDataResponse        `json:"metaData,omitempty"`
}
