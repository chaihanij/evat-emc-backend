package dtos

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

type FindAllMemberRequest struct {
	UUID     *string `form:"uuid" example:"2023"`
	TeamUUID *string `form:"teamUUID" example:"2023"`
	Year     *string `form:"year" example:"2023"`
	Sort     *string `form:"sort" example:"2023"`
	Page     *int64  `form:"page" validate:"omitempty,gte=1" example:"1"`
	PageSize *int64  `form:"pageSize" validate:"omitempty,gte=1" example:"20"`
}

func (req *FindAllMemberRequest) Parse(c *gin.Context) (*FindAllMemberRequest, error) {
	if err := c.ShouldBindQuery(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	return req, nil
}

func (req *FindAllMemberRequest) ToEntity() *entities.MemberFilter {
	return &entities.MemberFilter{
		UUID:     req.UUID,
		TeamUUID: req.TeamUUID,
		Year:     req.Year,
		Sort:     req.Sort,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
}

type FindAllMemberResponseJSON []MemberResponse

func (res *FindAllMemberResponseJSON) Parse(input []entities.Member) *FindAllMemberResponseJSON {
	var members FindAllMemberResponseJSON
	for _, value := range input {
		var member MemberResponse
		copier.Copy(&member, value)
		if val, ok := value.Image.(entities.File); ok {
			log.WithField("value", val).Debug("Parse image ")
		} else {
			member.Image = nil
		}
		if val, ok := value.Documents.(entities.Files); ok {
			log.WithField("value", val).Debug("Parse documents")
		} else {
			member.Documents = &FilesResponse{}
		}
		members = append(members, member)
	}
	return &members
}

type FindAllMemberResponseSwagger struct {
	StatusCode    int                       `json:"statusCode" example:"1000"`
	StatusMessage string                    `json:"statusMessage" example:"Success"`
	Timestamp     time.Time                 `json:"timestamp" example:"2015-06-30T21:59:59Z"`
	Data          FindAllMemberResponseJSON `json:"data,omitempty"`
	MetaData      MetaDataResponse          `json:"metaData,omitempty"`
}
