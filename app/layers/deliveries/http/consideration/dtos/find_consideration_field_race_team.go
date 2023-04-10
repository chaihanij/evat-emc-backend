package dtos

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

type ConsiderationFieldRaceTeamRequestJSON struct {
	UUID string `uri:"rield_race_UUID"`
	ID   string `uri:"id"`
}

func (req *ConsiderationFieldRaceTeamRequestJSON) Parse(c *gin.Context) (*ConsiderationFieldRaceTeamRequestJSON, error) {

	if err := c.ShouldBindUri(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	return req, nil
}

func (req *ConsiderationFieldRaceTeamRequestJSON) ToEntity() *entities.ConsiderationFilter {
	return &entities.ConsiderationFilter{
		AssignmentUUID: &req.UUID,
		ID:             &req.ID,
	}
}

type ConsiderationFieldRaceTeamResponeJSON []ConsiderationFieldRaceTeamResponse

func (m *ConsiderationFieldRaceTeamResponeJSON) Parse(c *gin.Context, data []entities.FieldRaceTeamScore) *ConsiderationFieldRaceTeamResponeJSON {

	var fieldRaceTeamScores ConsiderationFieldRaceTeamResponeJSON = ConsiderationFieldRaceTeamResponeJSON{}

	for _, value := range data {
		var considerationsFieldRaceTeam []Consideration

		for _, valueConsideration := range value.Considerations {

			consideration := &Consideration{
				ID:       valueConsideration.ID,
				TeamName: valueConsideration.TeamName,
				Score:    valueConsideration.Score,
				Title:    valueConsideration.Title,
			}

			considerationsFieldRaceTeam = append(considerationsFieldRaceTeam, *consideration)

		}

		fieldRaceTeamScore := &ConsiderationFieldRaceTeamResponse{
			Total:          value.Total,
			Considerations: considerationsFieldRaceTeam,
		}

		fieldRaceTeamScores = append(fieldRaceTeamScores, *fieldRaceTeamScore)

	}

	return &fieldRaceTeamScores

}
