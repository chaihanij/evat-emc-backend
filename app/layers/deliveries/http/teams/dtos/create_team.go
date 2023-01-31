package dtos

import (
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
	"gitlab.com/chaihanij/evat/app/types"
)

type CreateTeamRequestJSON struct {
	Code      string `json:"code" validate:"required"`
	Name      string `json:"name" validate:"required"`
	TeamType  string `json:"teamType" validate:"required,teamType"`
	Academy   string `json:"academy" validate:""`
	Detail    string `json:"detali" validate:""`
	Year      string `json:"year" validate:"required"`
	CreatedBy string `json:"-" swaggerignore:"true"`
}

type CreateTeamResponseJSON struct {
	UUID      string    `json:"uuid"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	TeamType  string    `json:"teamType"`
	Academy   string    `json:"academy"`
	Detail    string    `json:"detail"`
	Year      string    `json:"year"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedBy string    `json:"createdBy"`
	UpdatedBy string    `json:"updatedBy"`
}

func (req *CreateTeamRequestJSON) Parse(c *gin.Context) (*CreateTeamRequestJSON, error) {

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
	req.CreatedBy = jwtData.UID

	return req, nil
}

func (req *CreateTeamRequestJSON) ToEntity() *entities.Team {
	return &entities.Team{
		Code:      req.Code,
		Name:      req.Name,
		TeamType:  req.TeamType,
		Academy:   req.Academy,
		Detail:    req.Detail,
		Year:      req.Year,
		CreatedBy: req.CreatedBy,
	}
}

func (m *CreateTeamResponseJSON) Parse(input *entities.Team) *CreateTeamResponseJSON {
	return &CreateTeamResponseJSON{
		UUID:      input.UUID,
		Code:      input.Code,
		Name:      input.Name,
		TeamType:  input.TeamType,
		Academy:   input.Academy,
		Detail:    input.Detail,
		Year:      input.Year,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
		CreatedBy: input.CreatedBy,
		UpdatedBy: input.UpdatedBy,
	}
}

type CreateTeamResponseSwagger struct {
	StatusCode    int                    `json:"statusCode" example:"1000"`
	StatusMessage string                 `json:"statusMessage" example:"Success"`
	Timestamp     time.Time              `json:"timestamp" example:"2015-06-30T21:59:59Z"`
	Data          CreateTeamResponseJSON `json:"data,omitempty"`
}
