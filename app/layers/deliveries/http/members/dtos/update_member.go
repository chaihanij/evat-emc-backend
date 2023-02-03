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

type UpdateMemberRequest struct {
	// uri
	UUID string `json:"-" uri:"member_uuid" validate:"required" binding:"required,uuid"`
	// json
	FirstName    string `json:"firstname" validate:"required"`
	LastName     string `json:"lastname"  validate:"required"`
	Address      string `json:"address"`
	Email        string `json:"email"  validate:"email"`
	Tel          string `json:"tel"`
	Academy      string `json:"academy"`
	Year         string `json:"year"`
	MemberType   string `json:"memberType" validate:"required,memberType" example:"MEMBER, MENTOR"`
	TeamUUID     string `json:"teamUUID" validate:"required"`
	IsTeamLeader bool   `json:"isTeamLeader"`
	//
	UpdatedBy string `json:"-" swaggerignore:"true"`
}

func (req *UpdateMemberRequest) Parse(c *gin.Context) (*UpdateMemberRequest, error) {
	if err := c.ShouldBindUri(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
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
	req.UpdatedBy = jwtData.UID
	return req, nil
}

func (req *UpdateMemberRequest) ToEntity() *entities.MemberPartialUpdate {
	member := &entities.MemberPartialUpdate{
		UUID:         req.UUID,
		FirstName:    &req.FirstName,
		LastName:     &req.LastName,
		Address:      &req.Address,
		Email:        &req.Email,
		Tel:          &req.Tel,
		Academy:      &req.Academy,
		Year:         &req.Year,
		MemberType:   &req.MemberType,
		TeamUUID:     &req.TeamUUID,
		IsTeamLeader: &req.IsTeamLeader,
		UpdatedBy:    &req.UpdatedBy,
	}
	return member
}

type UpdateMemberResponseJSON MemberResponse

func (res *UpdateMemberResponseJSON) Parse(c *gin.Context, input *entities.Member) *UpdateMemberResponseJSON {
	copier.Copy(res, input)
	if val, ok := input.Image.(entities.File); ok {
		res.Image = new(FileResponse).Parse(c, &val)
	} else {
		res.Image = nil
	}
	if val, ok := input.Documents.([]entities.File); ok {
		var documents FilesResponse
		for _, value := range val {
			document := new(FileResponse).Parse(c, &value)
			documents = append(documents, *document)
		}
		res.Documents = &documents
	} else {
		res.Documents = &FilesResponse{}
	}
	return res
}

type UpdateMemberResponseSwagger struct {
	StatusCode    int                      `json:"statusCode" example:"1000"`
	StatusMessage string                   `json:"statusMessage" example:"Success"`
	Timestamp     time.Time                `json:"timestamp" example:"2015-06-30T21:59:59Z"`
	Data          UpdateMemberResponseJSON `json:"data,omitempty"`
}
