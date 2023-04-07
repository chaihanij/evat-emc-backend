package dtos

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

type FindOneConsiderationRequestJSON struct {
	UUID string `uri:"assignment_UUID"`
	ID   string `uri:"id"`
}

func (req *FindOneConsiderationRequestJSON) Parse(c *gin.Context) (*FindOneConsiderationRequestJSON, error) {

	if err := c.ShouldBindUri(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	return req, nil
}

func (req *FindOneConsiderationRequestJSON) ToEntity() *entities.ConsiderationFilter {
	return &entities.ConsiderationFilter{
		AssignmentUUID: &req.UUID,
		ID:             &req.ID,
	}
}

// type FindOneConsiderationResponseJSON ConsiderationResponse

type FindOneConsiderationsResponseJSON []ConsiderationResponse

func (m *FindOneConsiderationsResponseJSON) Parse(c *gin.Context, data []entities.AssignmentScore) *FindOneConsiderationsResponseJSON {

	var assignmentsScores FindOneConsiderationsResponseJSON = FindOneConsiderationsResponseJSON{}
	for _, value := range data {

		var considerations []Consideration

		for _, valueConsideration := range value.Considerations {

			consideration := &Consideration{
				ID:       valueConsideration.ID,
				TeamName: valueConsideration.TeamName,
				Score:    valueConsideration.Score,
				Title:    valueConsideration.Title,
			}

			considerations = append(considerations, *consideration)

		}

		assignmentsScore := &ConsiderationResponse{
			// ID:             value.ID,
			Total:          value.Total,
			Considerations: considerations,
		}

		assignmentsScores = append(assignmentsScores, *assignmentsScore)

	}

	return &assignmentsScores
}
