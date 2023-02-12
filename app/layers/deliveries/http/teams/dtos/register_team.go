package dtos

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	IsConsent bool   `json:"isConsent" validate:"required"`
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
			Username:  uuid.NewString(),
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Email:     req.Email,
			Tel:       req.Tel,
			Year:      req.Year,
		}
}

type RegisterTeamResponseSwagger struct {
	StatusCode    int       `json:"statusCode" example:"1000"`
	StatusMessage string    `json:"statusMessage" example:"Success"`
	Timestamp     time.Time `json:"timestamp" example:"2015-06-30T21:59:59Z"`
}
