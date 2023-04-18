package dtos

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

type AllScoreRequestJSON struct {
	Name string `json:"name"`
}

// type AllScoreResponseJSON struct {
// 	Considerations []Consideration `json:"considerations"`
// }

func (req *AllScoreRequestJSON) Parse(c *gin.Context) (*AllScoreRequestJSON, error) {

	if err := c.ShouldBind(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}

	return req, nil

}

func (req *AllScoreRequestJSON) ToEntity() *entities.AllScoreFilter {
	return &entities.AllScoreFilter{
		// Name: req.Name,
	}
}

type AllScoreResponseJSON struct {
	ID                string              `json:"_id" bson:"id"`
	Title             string              `json:"title" bson:"title"`
	Total             float64             `json:"total" bson:"total"`
	AllConsiderations []AllConsiderations `json:"considerations" bson:"considerations"`
}
type AllConsiderations struct {
	Title string  `json:"title" bson:"title" `
	Score float64 `json:"score" bson:"score" `
}

type AllScoresResponseJSON []AllScoreResponseJSON

func (m *AllScoresResponseJSON) Parse(c *gin.Context, data []entities.AllScore) *AllScoresResponseJSON {

	var allScores AllScoresResponseJSON = AllScoresResponseJSON{}

	fmt.Println("-data :", data)

	for _, value := range data {

		// fmt.Println("data", value.Allconsiderations)

		var allConsideration []AllConsiderations
		for _, vl := range value.Allconsiderations {

			allScore := &AllConsiderations{
				Title: vl.Title,
				Score: vl.Score,
			}

			allConsideration = append(allConsideration, *allScore)

		}

		allScore := &AllScoreResponseJSON{
			ID:                value.ID,
			Title:             value.Title,
			Total:             value.Total,
			AllConsiderations: allConsideration,
		}

		allScores = append(allScores, *allScore)

	}

	return &allScores

}
