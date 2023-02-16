package dtos

import (
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

type FindOneAnnouncementRequestJSON struct {
	UUID string `uri:"announcement_uuid"`
}

func (req *FindOneAnnouncementRequestJSON) Parse(c *gin.Context) (*FindOneAnnouncementRequestJSON, error) {
	if err := c.ShouldBindUri(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	return req, nil
}

func (req *FindOneAnnouncementRequestJSON) ToEntity() *entities.AnnouncementFilter {
	return &entities.AnnouncementFilter{
		UUID: &req.UUID,
	}
}

type FindOneAnnouncementResponseJSON AnnouncementResponse

func (m *FindOneAnnouncementResponseJSON) Parse(c *gin.Context, input *entities.Announcement) *FindOneAnnouncementResponseJSON {
	announcement := &FindOneAnnouncementResponseJSON{
		UUID:        input.UUID,
		Title:       input.Title,
		Description: input.Description,
		Year:        input.Year,
		CreatedAt:   input.CreatedAt,
		UpdatedAt:   input.UpdatedAt,
		CreatedBy:   input.CreatedBy,
		UpdatedBy:   input.UpdatedBy,
	}

	return announcement
}

type FindOneAnnouncementResponseSwagger struct {
	StatusCode    int                             `json:"statusCode" example:"1000"`
	StatusMessage string                          `json:"statusMessage" example:"Success"`
	Timestamp     time.Time                       `json:"timestamp" example:"2015-06-30T21:59:59Z"`
	Data          FindOneAnnouncementResponseJSON `json:"data,omitempty"`
}
