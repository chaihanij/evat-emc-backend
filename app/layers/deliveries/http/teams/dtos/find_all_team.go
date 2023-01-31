package dtos

import (
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

type FindAllTeamRequestJSON struct {
	Year     *string `form:"year" example:"2023"`
	Page     *int64  `form:"page" validate:"omitempty,gte=1" example:"1"`
	PageSize *int64  `form:"pageSize" validate:"omitempty,gte=1" example:"20"`
}

func (req *FindAllTeamRequestJSON) Parse(c *gin.Context) (*FindAllTeamRequestJSON, error) {

	err := c.ShouldBindQuery(req)
	if err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	return req, nil
}

func (req *FindAllTeamRequestJSON) ToEntity() *entities.TeamFilter {
	return &entities.TeamFilter{
		Year:     req.Year,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
}

type FindAllTeamResponseJSON []TeamResponse

func (m *FindAllTeamResponseJSON) Parse(data []entities.Team) *FindAllTeamResponseJSON {
	var teams FindAllTeamResponseJSON = FindAllTeamResponseJSON{}
	for _, value := range data {
		team := &TeamResponse{
			UUID:      value.UUID,
			Code:      value.Code,
			Name:      value.Name,
			TeamType:  value.TeamType,
			Academy:   value.Academy,
			Detail:    value.Detail,
			Year:      value.Year,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
			CreatedBy: value.CreatedBy,
			UpdatedBy: value.UpdatedBy,
		}
		teams = append(teams, *team)
	}
	return &teams
}

type FindAllTeamResponseSwagger struct {
	StatusCode    int                     `json:"statusCode" example:"1000"`
	StatusMessage string                  `json:"statusMessage" example:"Success"`
	Timestamp     time.Time               `json:"timestamp" example:"2015-06-30T21:59:59Z"`
	Data          FindAllTeamResponseJSON `json:"data,omitempty"`
	MetaData      MetaDataResponse        `json:"metaData,omitempty"`
}
