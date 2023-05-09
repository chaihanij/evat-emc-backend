package dtos

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
	"gitlab.com/chaihanij/evat/app/types"
)

type CreateMemberRequestJSON struct {
	FirstName    string    `json:"firstname" validate:"required"`
	LastName     string    `json:"lastname"  validate:"required"`
	Address      string    `json:"address"`
	Email        string    `json:"email"  validate:"email"`
	Tel          string    `json:"tel"`
	Academy      string    `json:"academy"`
	Major        string    `json:"major"`
	Year         string    `json:"year"`
	MemberType   string    `json:"memberType" validate:"required,memberType" example:"MEMBER, MENTOR"`
	TeamUUID     string    `json:"teamUUID" validate:"required"`
	IsTeamLeader bool      `json:"isTeamLeader"`
	CreatedBy    string    `json:"-" swaggerignore:"true"`
	BirthDay     time.Time `json:"birth_day" validate:"required" `
	NationalId   string    `json:"national_id" validate:"required" `
}

type CreateMemberResponseJSON MemberResponse

func (req *CreateMemberRequestJSON) Parse(c *gin.Context) (*CreateMemberRequestJSON, error) {
	if err := c.ShouldBindJSON(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	err := types.Validate.Struct(req)
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

func (req *CreateMemberRequestJSON) ToEntity() *entities.Member {
	member := &entities.Member{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Address:      req.Address,
		Email:        req.Email,
		Tel:          req.Tel,
		Academy:      req.Academy,
		Major:        req.Major,
		Year:         req.Year,
		MemberType:   req.MemberType,
		TeamUUID:     req.TeamUUID,
		IsTeamLeader: req.IsTeamLeader,
		CreatedBy:    req.CreatedBy,
		NationalId:   req.NationalId,
		BirthDay:     req.BirthDay,
	}
	return member
}

func (res *CreateMemberResponseJSON) Parse(input *entities.Member) *CreateMemberResponseJSON {
	copier.Copy(res, input)
	if val, ok := input.Image.(entities.File); ok {
		log.WithField("value", val).Debug("Parse image ")
	} else {
		res.Image = nil
	}
	if val, ok := input.Documents.(entities.Files); ok {
		log.WithField("value", val).Debug("Parse documents")
	} else {
		res.Documents = &FilesResponse{}
	}
	return res
}

type CreateMemberResponseSwagger struct {
	StatusCode    int                      `json:"statusCode" example:"1000"`
	StatusMessage string                   `json:"statusMessage" example:"Success"`
	Timestamp     time.Time                `json:"timestamp" example:"2015-06-30T21:59:59Z"`
	Data          CreateMemberResponseJSON `json:"data,omitempty"`
}
