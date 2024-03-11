package dtos

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

type FindAllTeamSearchRequestJSON struct {
	Year     *string `form:"year" example:"2024"`
	Page     *int64  `form:"page" validate:"omitempty,gte=1" example:"1"`
	PageSize *int64  `form:"pageSize" validate:"omitempty,gte=1" example:"20"`
}

func (req *FindAllTeamRequestJSON) ParseTeamSearch(c *gin.Context) (*FindAllTeamRequestJSON, error) {

	err := c.ShouldBindQuery(req)
	if err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	return req, nil
}

func (req *FindAllTeamRequestJSON) ToEntityTeamSearch() *entities.TeamFilter {
	return &entities.TeamFilter{
		Year:     req.Year,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
}

type FindAllTeamSearchResponseJSON []TeamSearchResponse

func (m *FindAllTeamSearchResponseJSON) ParseTeamSearch(data []entities.TeamSearch) *FindAllTeamSearchResponseJSON {
	var teams FindAllTeamSearchResponseJSON = FindAllTeamSearchResponseJSON{}
	for _, value := range data {
		team := &TeamSearchResponse{
			UUID:     value.UUID,
			Code:     value.Code,
			Name:     value.Name,
			TeamType: value.TeamType,
			Academy:  value.Academy,
			Tel:      value.Tel,
			Contact:  value.Contact,
		}
		teams = append(teams, *team)
	}
	return &teams
}
