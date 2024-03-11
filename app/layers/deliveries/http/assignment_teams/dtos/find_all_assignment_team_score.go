package dtos

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

type FindAllAssignmentRequestTeamJSON struct {
	Year     *string `form:"year" example:"2024"`
	Page     *int64  `form:"page" validate:"omitempty,gte=1" example:"1"`
	PageSize *int64  `form:"pageSize" validate:"omitempty,gte=1" example:"20"`
	TeamUUID *string `uri:"team_uuid"`
}

func (req *FindAllAssignmentRequestTeamJSON) Parse(c *gin.Context) (*FindAllAssignmentRequestTeamJSON, error) {

	err := c.ShouldBindQuery(req)
	if err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	if err := c.ShouldBindUri(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	return req, nil
}
func (req *FindAllAssignmentRequestTeamJSON) ToEntity() *entities.AssignmentTeamFilter {
	return &entities.AssignmentTeamFilter{
		Year:     req.Year,
		Page:     req.Page,
		PageSize: req.PageSize,
		TeamUUID: req.TeamUUID,
	}
}

type FindAllAssignmentTeamResponseJSON []AssignmentTeamResponse

func (m *FindAllAssignmentTeamResponseJSON) Parse(data []entities.AssignmentTeamScore) *FindAllAssignmentTeamResponseJSON {
	var assignmentteamscores FindAllAssignmentTeamResponseJSON = FindAllAssignmentTeamResponseJSON{}
	for _, value := range data {
		assignment := &AssignmentTeamResponse{
			Score: value.Score,
			Title: value.Title,
		}
		assignmentteamscores = append(assignmentteamscores, *assignment)
	}
	return &assignmentteamscores
}
