package dtos

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

type AllScoreRequestJSON struct {
	Name string `json:"name"`
}

type AllScoreResponseJSON struct {
	Considerations []Consideration `json:"considerations"`
}

func (req *AllScoreRequestJSON) Parse(c *gin.Context) (*AllScoreRequestJSON, error) {

	if err := c.ShouldBind(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}

	return req, nil

}

func (req *AllScoreRequestJSON) ToEntity() *entities.AllScoreFilter {
	return &entities.AllScoreFilter{
		Name: req.Name,
	}
}

type AllScoresResponseJSON []AllScoreResponseJSON

func (m *AllScoresResponseJSON) Parse(c *gin.Context, data []entities.AllScore) *AllScoresResponseJSON {

	var allScores AllScoresResponseJSON = AllScoresResponseJSON{}

	for _, value := range data {

		var allConsideration []Consideration
		for _, vl := range value.Allconsiderations {

			allScore := &Consideration{
				ID:       vl.ID,
				TeamName: vl.TeamName,
				Title:    vl.Title,
				Score:    vl.Score,
			}

			allConsideration = append(allConsideration, *allScore)

		}

		allScore := &AllScoreResponseJSON{
			Considerations: allConsideration,
		}

		allScores = append(allScores, *allScore)

	}

	return &allScores

}
