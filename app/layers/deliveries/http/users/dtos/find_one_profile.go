package dtos

import (
	"time"

	"github.com/AlekSi/pointer"
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

type FindOneUserProfileRequestJSON struct {
	UID *string `json:"-"`
}

type FindOneUserProfileResponseJSON struct {
	UID       string     `json:"uid" example:"40d7b15e-cec9-4735-a829-5c69fc508e1e"`
	Username  string     `json:"username"`
	Email     string     `json:"email" example:"test@gmail.com"`
	FirstName string     `json:"firstname"`
	LastName  string     `json:"lastname"`
	Address   string     `json:"address"  example:"address"`
	Tel       string     `json:"tel"`
	Role      string     `json:"role"`
	LastLogin *time.Time `json:"lastLogin"  example:"2022-05-22T00:00:00Z,null"`
	UpdatedAt time.Time  `json:"updatedAt"  example:"2022-05-22T00:00:00Z"`
	CreatedAt time.Time  `json:"createdAt"  example:"2022-05-22T00:00:00Z"`
}

func (req *FindOneUserProfileRequestJSON) Parse(c *gin.Context) (*FindOneUserProfileRequestJSON, error) {

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

	req.UID = &jwtData.UID

	return req, nil

}

func (req *FindOneUserProfileRequestJSON) ToEntity() *entities.UserFilter {
	return &entities.UserFilter{
		UID:      req.UID,
		IsActive: pointer.ToBool(true),
	}
}

func (m *FindOneUserProfileResponseJSON) Parse(data *entities.User) *FindOneUserProfileResponseJSON {
	var lastLogin *time.Time
	if !data.LastLogin.IsZero() {
		lastLogin = pointer.ToTime(data.LastLogin)
	} else {
		lastLogin = nil
	}
	return &FindOneUserProfileResponseJSON{
		UID:       data.UID,
		Username:  data.Username,
		Email:     data.Email,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Address:   data.Address,
		Tel:       data.Tel,
		Role:      string(data.Role),
		LastLogin: lastLogin,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

type FindOneUserProfileResponseSwagger struct {
	StatusCode    int                            `json:"statusCode" example:"1000"`
	StatusMessage string                         `json:"statusMessage" example:"Success"`
	Timestamp     time.Time                      `json:"timestamp" example:"2015-06-30T21:59:59Z"`
	Data          FindOneUserProfileResponseJSON `json:"data,omitempty"`
}
