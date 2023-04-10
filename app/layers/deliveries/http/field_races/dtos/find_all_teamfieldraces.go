package dtos

import (

	// log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

type FindAllTeamFieldRacestRequestJSON struct {
	Year     *string `form:"year" example:"2023"`
	Page     *int64  `form:"page" validate:"omitempty,gte=1" example:"1"`
	PageSize *int64  `form:"pageSize" validate:"omitempty,gte=1" example:"20"`
	UUID     *string `uri:"uuid"`
}

func (req *FindAllTeamFieldRacestRequestJSON) Parse(c *gin.Context) (*FindAllTeamFieldRacestRequestJSON, error) {
	err := c.ShouldBindQuery(req)
	if err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	return req, nil
}
func (req *FindAllTeamFieldRacestRequestJSON) ToEntity() *entities.FieldRaceFilter {
	return &entities.FieldRaceFilter{
		Year:     req.Year,
		Page:     req.Page,
		PageSize: req.PageSize,
		UUID:     req.UUID,
	}
}

type FindAllTeamFieldRecestResponseJSON []FieldRaces

func (m *FindAllTeamFieldRecestResponseJSON) Parse(data []entities.FieldRace) *FindAllTeamFieldRecestResponseJSON {
	var assignments FindAllTeamFieldRecestResponseJSON = FindAllTeamFieldRecestResponseJSON{}
	for _, value := range data {
		assignment := &FieldRaces{
			UUID:        value.UUID,
			No:          value.No,
			Title:       value.Title,
			Description: value.Description,
			FullScore:   value.FullScore,
			IsActive:    value.IsActive,
			Year:        value.Year,
			CreatedAt:   value.CreatedAt,
			UpdatedAt:   value.UpdatedAt,
			CreatedBy:   value.CreatedBy,
			UpdatedBy:   value.UpdatedBy,
		}
		assignments = append(assignments, *assignment)
	}
	return &assignments
}
