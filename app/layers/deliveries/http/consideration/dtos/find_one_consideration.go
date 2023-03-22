package dtos

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

type FindOneConsiderationRequestJSON struct {
	UUID string `uri:"team_uuid"`
}

func (req *FindOneConsiderationRequestJSON) Parse(c *gin.Context) (*FindOneConsiderationRequestJSON, error) {

	if err := c.ShouldBindUri(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	return req, nil
}

func (req *FindOneConsiderationRequestJSON) ToEntity() *entities.ConsiderationFilter {
	return &entities.ConsiderationFilter{
		TeamUUID: &req.UUID,
	}
}

type FindOneConsiderationResponseJSON ConsiderationResponse

func (m *FindOneConsiderationResponseJSON) Parse(c *gin.Context, input *entities.Consideration) *FindOneConsiderationResponseJSON {
	
	var indivdualScores []IndivdualScore

	for _, data := range input.IndivdualScore {

		indivdualScore := IndivdualScore{
			Title:  data.Title,
			Score: data.Score,
		}

		indivdualScores = append(indivdualScores, indivdualScore)

	}

	consideration := &FindOneConsiderationResponseJSON{
		ID:             input.ID,
		TotalScore:     input.TotalScore,
		No:             input.No,
		UpdatedAt:      input.UpdatedAt,
		IndivdualScore: indivdualScores,
	}

	return consideration
}
