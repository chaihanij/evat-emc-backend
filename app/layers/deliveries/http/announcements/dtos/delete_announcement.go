package dtos

import (
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

type DeleteAnnouncementRequestJSON struct {
	UUID string `uri:"announcement_uuid"`
}

func (req *DeleteAnnouncementRequestJSON) Parse(c *gin.Context) (*DeleteAnnouncementRequestJSON, error) {
	if err := c.ShouldBindUri(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	return req, nil
}

func (req *DeleteAnnouncementRequestJSON) ToEntity() *entities.AnnouncementFilter {
	return &entities.AnnouncementFilter{
		UUID: &req.UUID,
	}
}

type DeleteAnnouncementResponseSwagger struct {
	StatusCode    int       `json:"statusCode" example:"1000"`
	StatusMessage string    `json:"statusMessage" example:"Success"`
	Timestamp     time.Time `json:"timestamp" example:"2015-06-30T21:59:59Z"`
}
