package dtos

import (
	"time"

	"github.com/AlekSi/pointer"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/omise/omise-go"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
	"gitlab.com/chaihanij/evat/app/types"
)

type RegisterTeamRequestJSON struct {
	TeamName  string `json:"teamName" validate:"required" `
	TeamType  string `json:"teamType" validate:"required,teamType"`
	FirstName string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Tel       string `json:"tel" validate:"required"`
	Year      string `json:"year" validate:"required"`
	Password  string `json:"password" validate:"required,passwordComplexity"`
	IsConsent bool   `json:"isConsent" validate:"required"`
	Username  string `json:"username"`
}

func (req *RegisterTeamRequestJSON) Parse(c *gin.Context) (*RegisterTeamRequestJSON, error) {

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

func (req *RegisterTeamRequestJSON) ToEntity() (*entities.Team, *entities.User) {

	return &entities.Team{
			Code:     uuid.NewString(),
			Name:     req.TeamName,
			TeamType: req.TeamType,
			Year:     req.Year,
		},
		&entities.User{
			Username:  req.Username,
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Email:     req.Email,
			Password:  req.Password,
			Tel:       req.Tel,
			Year:      req.Year,
			Role:      string(types.UserRoleUSER),
		}
}

type FindOneUserResponseJSON UserResponse

func (m *FindOneUserResponseJSON) Parse(data *entities.User) *FindOneUserResponseJSON {

	copier.Copy(m, data)
	if !data.LastLogin.IsZero() {
		m.LastLogin = pointer.ToTime(data.LastLogin)
	} else {
		m.LastLogin = nil
	}
	return m
}

type RegisterTeamResponseJSON struct {
	Team          *FindOneTeamResponseJSON `json:"team"`
	User          *FindOneUserResponseJSON `json:"user"`
	ScannableCode *omise.ScannableCode     `json:"scannable_code"`
}

func (registerTeamResponseJSON *RegisterTeamResponseJSON) Parse(c *gin.Context, team *entities.Team, user *entities.User, charge *entities.OmiseCharge) *RegisterTeamResponseJSON {
	return &RegisterTeamResponseJSON{
		Team:          new(FindOneTeamResponseJSON).Parse(c, team),
		User:          new(FindOneUserResponseJSON).Parse(user),
		ScannableCode: charge.Source.ScannableCode,
	}
}

type RegisterTeamResponseSwagger struct {
	StatusCode    int       `json:"statusCode" example:"1000"`
	StatusMessage string    `json:"statusMessage" example:"Success"`
	Timestamp     time.Time `json:"timestamp" example:"2015-06-30T21:59:59Z"`
}
