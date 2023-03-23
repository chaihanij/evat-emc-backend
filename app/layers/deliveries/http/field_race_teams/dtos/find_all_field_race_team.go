package dtos

import (
	"github.com/gin-gonic/gin"
	// log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

type FindAllField_race_teamsRequestJSON struct {
	// Year     *string `form:"year" example:"2023"`
	Page      *int64  `form:"page" validate:"omitempty,gte=1" example:"1"`
	PageSize  *int64  `form:"pageSize" validate:"omitempty,gte=1" example:"20"`
	Team_uuid *string `uri:"team_uuid"`
}

func (req *FindAllField_race_teamsRequestJSON) Parse(c *gin.Context) (*FindAllField_race_teamsRequestJSON, error) {

	err := c.ShouldBindQuery(req)
	if err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}

	if err := c.ShouldBindUri(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	


	return req, nil
}

func (req *FindAllField_race_teamsRequestJSON) ToEntity() *entities.FieldRaceTeamFilter {
	return &entities.FieldRaceTeamFilter{
		// Year:     req.Year,
		TeamUUID: req.Team_uuid,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
}

type FindAllField_race_teamsResponseJSON []FieldRaceTeam

func (m *FindAllField_race_teamsResponseJSON) Parse(data []entities.FieldRaceTeam) *FindAllField_race_teamsResponseJSON {
	var field_race_teams FindAllField_race_teamsResponseJSON = FindAllField_race_teamsResponseJSON{}
	for _, value := range data {

		// var fieldRaces []FieldRace

		// for _, valueFieldRace := range value.FieldRaces {
		// 	// log.Debugln("dddd", valueFieldRace)
		// 	fieldRace := &FieldRace{
		// 		Title:       valueFieldRace.Title,
		// 		Description: valueFieldRace.Description,
		// 		File:        valueFieldRace.File,
		// 		Image:       valueFieldRace.Image,
		// 		Year:        valueFieldRace.Year,
		// 		FullScore:   valueFieldRace.FullScore,
		// 	}
		// 	fieldRaces = append(fieldRaces, *fieldRace)
		// }

		field_race_team := &FieldRaceTeam{
			FieldRaceUUID: value.FieldRaceUUID,
			TeamUUID:      value.TeamUUID,
			Description:   value.Description,
			Score:         value.Score,
			CreatedAt:     value.CreatedAt,
			UpdatedAt:     value.UpdatedAt,
			CreatedBy:     value.CreatedBy,
			UpdatedBy:     value.UpdatedBy,
			Name:          value.Name,
			Code:          value.Code,
			Type:          value.Type,
			// FieldRaces:    fieldRaces,
		}
		field_race_teams = append(field_race_teams, *field_race_team)
	}
	return &field_race_teams
}
