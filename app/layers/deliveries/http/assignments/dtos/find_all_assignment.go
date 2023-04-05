package dtos

import (
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

type FindAllAssignmentRequestJSON struct {
	Year     *string `form:"year" example:"2023"`
	Page     *int64  `form:"page" validate:"omitempty,gte=1" example:"1"`
	PageSize *int64  `form:"pageSize" validate:"omitempty,gte=1" example:"20"`
}

func (req *FindAllAssignmentRequestJSON) Parse(c *gin.Context) (*FindAllAssignmentRequestJSON, error) {

	err := c.ShouldBindQuery(req)
	if err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	return req, nil
}

func (req *FindAllAssignmentRequestJSON) ToEntity() *entities.AssignmentFilter {
	return &entities.AssignmentFilter{
		Year:     req.Year,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
}

type FindAllAssignmentResponseJSON []AssignmentResponse

func (m *FindAllAssignmentResponseJSON) Parse(data []entities.Assignment) *FindAllAssignmentResponseJSON {
	var assignments FindAllAssignmentResponseJSON = FindAllAssignmentResponseJSON{}
	for _, value := range data {
		assignment := &AssignmentResponse{
			UUID:         value.UUID,
			No:           value.No,
			Title:        value.Title,
			Description:  value.Description,
			FullScore:    value.FullScore,
			IsActive:     value.IsActive,
			DueDate:      value.DueDate,
			Year:         value.Year,
			CreatedAt:    value.CreatedAt,
			UpdatedAt:    value.UpdatedAt,
			CreatedBy:    value.CreatedBy,
			UpdatedBy:    value.UpdatedBy,
			DeliveryTime: value.DeliveryTime,
		}
		assignments = append(assignments, *assignment)
	}
	return &assignments
}

// type FindTeamAssignmentResponseJSON []TeamAssignmentResponse

// func (a *FindTeamAssignmentResponseJSON) Parse(data []entities.TeamAssignment) *FindTeamAssignmentResponseJSON {
// 	var teamAssignments FindTeamAssignmentResponseJSON = FindTeamAssignmentResponseJSON{}
// 	for _, value := range data {
// 		teamAssignment := &TeamAssignmentResponse{
// 			UUID:      value.UUID,
// 			TeamUUID:  value.TeamUUID,
// 			Title:     value.Title,
// 			FullScore: value.FullScore,
// 		}
// 		teamAssignments = append(teamAssignments, *teamAssignment)
// 	}
// 	return &teamAssignments

// }

type FindAllAssignmentResponseSwagger struct {
	StatusCode    int                           `json:"statusCode" example:"1000"`
	StatusMessage string                        `json:"statusMessage" example:"Success"`
	Timestamp     time.Time                     `json:"timestamp" example:"2015-06-30T21:59:59Z"`
	Data          FindAllAssignmentResponseJSON `json:"data,omitempty"`
	MetaData      MetaDataResponse              `json:"metaData,omitempty"`
}
