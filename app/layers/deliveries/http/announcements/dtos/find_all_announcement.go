package dtos

import (
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

type FindAllAnnouncementRequestJSON struct {
	Year     *string `form:"year" example:"2024"`
	Page     *int64  `form:"page" validate:"omitempty,gte=1" example:"1"`
	PageSize *int64  `form:"pageSize" validate:"omitempty,gte=1" example:"20"`
}

func (req *FindAllAnnouncementRequestJSON) Parse(c *gin.Context) (*FindAllAnnouncementRequestJSON, error) {

	err := c.ShouldBindQuery(req)
	if err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	return req, nil
}

func (req *FindAllAnnouncementRequestJSON) ToEntity() *entities.AnnouncementFilter {
	return &entities.AnnouncementFilter{
		Year:     req.Year,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
}

type FindAllAnnouncementResponseJSON []AnnouncementResponse

func (m *FindAllAnnouncementResponseJSON) Parse(data []entities.Announcement) *FindAllAnnouncementResponseJSON {
	var announcements FindAllAnnouncementResponseJSON = FindAllAnnouncementResponseJSON{}
	for _, value := range data {
		announcement := &AnnouncementResponse{
			UUID:        value.UUID,
			Title:       value.Title,
			Description: value.Description,
			Year:        value.Year,
			CreatedAt:   value.CreatedAt,
			UpdatedAt:   value.UpdatedAt,
			CreatedBy:   value.CreatedBy,
			UpdatedBy:   value.UpdatedBy,
		}
		announcements = append(announcements, *announcement)
	}
	return &announcements
}

type FindAllAnnouncementResponseSwagger struct {
	StatusCode    int                             `json:"statusCode" example:"1000"`
	StatusMessage string                          `json:"statusMessage" example:"Success"`
	Timestamp     time.Time                       `json:"timestamp" example:"2015-06-30T21:59:59Z"`
	Data          FindAllAnnouncementResponseJSON `json:"data,omitempty"`
	MetaData      MetaDataResponse                `json:"metaData,omitempty"`
}
