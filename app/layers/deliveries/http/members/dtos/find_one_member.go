package dtos

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

type FindOneMemberRequest struct {
	UUID string `uri:"member_uuid"`
}

func (req *FindOneMemberRequest) Parse(c *gin.Context) (*FindOneMemberRequest, error) {
	if err := c.ShouldBindUri(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	return req, nil
}

func (req *FindOneMemberRequest) ToEntity() *entities.MemberFilter {
	return &entities.MemberFilter{
		UUID: &req.UUID,
	}
}

type FindOneMemberResponseJSON MemberResponse

func (res *FindOneMemberResponseJSON) Parse(input *entities.Member) *FindOneMemberResponseJSON {
	copier.Copy(res, input)
	if val, ok := input.Image.(entities.File); ok {
		var image FileResponse
		copier.Copy(image, val)
		res.Image = &image
	} else {
		res.Image = nil
	}
	if val, ok := input.Documents.(entities.Files); ok {
		var documents FilesResponse
		for _, value := range val {
			var document FileResponse
			copier.Copy(document, value)
			documents = append(documents, document)
		}
		res.Documents = &documents
	} else {
		res.Documents = &FilesResponse{}
	}
	log.WithField("value", res).Debug("FindOneMemberResponseJSON Parse")
	return res
}

type FindOneMemberResponseSwagger struct {
	StatusCode    int                       `json:"statusCode" example:"1000"`
	StatusMessage string                    `json:"statusMessage" example:"Success"`
	Timestamp     time.Time                 `json:"timestamp" example:"2015-06-30T21:59:59Z"`
	Data          FindOneMemberResponseJSON `json:"data,omitempty"`
}
