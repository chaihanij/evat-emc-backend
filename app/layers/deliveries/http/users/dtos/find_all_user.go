package dtos

import (
	"time"

	"github.com/AlekSi/pointer"
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

type FindAllUserRequestJSON struct {
	Year     *string `form:"year" example:"2023"`
	Page     *int64  `form:"page" validate:"omitempty,gte=1" example:"1"`
	PageSize *int64  `form:"pageSize" validate:"omitempty,gte=1" example:"20"`
}

type FindAllUserResponseJSON []User

func (req *FindAllUserRequestJSON) Parse(c *gin.Context) (*FindAllUserRequestJSON, error) {

	err := c.ShouldBindQuery(req)
	if err != nil {
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
			CreatedBy: value.CreatedBy,
			UpdatedBy: value.UpdatedBy,
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
