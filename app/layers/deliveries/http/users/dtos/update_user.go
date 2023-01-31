package dtos

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
	"gitlab.com/chaihanij/evat/app/types"
)

type UpdateUserRequestJSON struct {
	UID       string  `json:"-" uri:"uid" validate:"required"`
	Username  *string `json:"username" validate:"required"`
	Email     *string `json:"email" validate:""`
	FirstName *string `json:"firstname" validate:""`
	LastName  *string `json:"lastname" validate:""`
	Address   *string `json:"address" validate:""`
	Tel       *string `json:"tel" validate:""`
	Role      *string `json:"role" validate:"userRole"`
	Year      *string `json:"year" validate:""`
	TeamUUID  *string `json:"teamUUID" validate:""`
	IsActive  *bool   `json:"isActive" validate:""`
	UpdatedBy string  `json:"-" validate:""`
}

type UpdateUserResponseJSON struct {
	UID       string    `json:"uid"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Address   string    `json:"address"`
	Tel       string    `json:"tel"`
	Role      string    `json:"role"`
	Year      string    `json:"year"`
	TeamUUID  string    `json:"teamUUID"`
	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedBy string    `json:"createdBy"`
	UpdatedBy string    `json:"updatedBy"`
}

func (req *UpdateUserRequestJSON) Parse(c *gin.Context) (*UpdateUserRequestJSON, error) {
	if err := c.ShouldBindUri(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	if err := c.ShouldBindJSON(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}

	if err := types.Validate.Struct(req); err != nil {
		if err := types.HandleValidateError(err, req); err != nil {
			return nil, errors.ParameterError{Message: err.Error()}
		}
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
	req.UpdatedBy = jwtData.UID

	return req, nil
}

func (req *UpdateUserRequestJSON) ToEntity() *entities.UserPartialUpdate {
	return &entities.UserPartialUpdate{
		UID:       req.UID,
		Username:  req.Username,
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Address:   req.Address,
		Tel:       req.Tel,
		Role:      req.Role,
		Year:      req.Year,
		TeamUUID:  req.TeamUUID,
		IsActive:  req.IsActive,
		UpdatedBy: &req.UpdatedBy,
	}
}

func (m *UpdateUserResponseJSON) Parse(data *entities.User) *UpdateUserResponseJSON {
	_ = copier.Copy(m, data)
	return m
}

type UpdateUserResponseSwagger struct {
	StatusCode    int                    `json:"statusCode" example:"1000"`
	StatusMessage string                 `json:"statusMessage" example:"Success"`
	Timestamp     time.Time              `json:"timestamp" example:"2015-06-30T21:59:59Z"`
	Data          UpdateUserResponseJSON `json:"data,omitempty"`
}
