package dtos

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

type FindTeamAssignmentRequestJSON struct {
	Year     *string `form:"year" example:"2024"`
	Page     *int64  `form:"page" validate:"omitempty,gte=1" example:"1"`
	PageSize *int64  `form:"pageSize" validate:"omitempty,gte=1" example:"20"`
	TeamUUID *string `uri:"team_uuid"`
}

func (req *FindTeamAssignmentRequestJSON) Parse(c *gin.Context) (*FindTeamAssignmentRequestJSON, error) {

	err := c.ShouldBindQuery(req)
	if err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	if err := c.ShouldBindUri(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	return req, nil
}

type FindTeamAssignmentResponseJSON []TeamAssignmentResponse

func (a *FindTeamAssignmentResponseJSON) Parse(data []entities.TeamAssignment) *FindTeamAssignmentResponseJSON {
	var teamAssignments FindTeamAssignmentResponseJSON = FindTeamAssignmentResponseJSON{}
	for _, value := range data {
		teamAssignment := &TeamAssignmentResponse{
			UUID:      value.UUID,
			TeamUUID:  value.TeamUUID,
			Title:     value.Title,
			FullScore: value.FullScore,
		}
		teamAssignments = append(teamAssignments, *teamAssignment)
	}
	return &teamAssignments

}

func (req *FindTeamAssignmentRequestJSON) ToEntity() *entities.AssignmentFilter {
	return &entities.AssignmentFilter{
		Year:     req.Year,
		Page:     req.Page,
		PageSize: req.PageSize,
		TeamUUID: req.TeamUUID,
	}
}
