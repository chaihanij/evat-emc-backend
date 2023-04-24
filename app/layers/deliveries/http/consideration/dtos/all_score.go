package dtos

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

type AllScoreRequestJSON struct {
	Name     string `form:"name"`
	TeamType string `form:"teamtype"`
	Page     int    `form:"page" validate:"omitempty,gte=1" example:"1"`
	PageSize int    `form:"pageSize" validate:"omitempty,gte=1" example:"10"`
	Code     string `form:"code"`
}

// type AllScoreResponseJSON struct {
// 	Considerations []Consideration `json:"considerations"`
// }

func (req *AllScoreRequestJSON) Parse(c *gin.Context) (*AllScoreRequestJSON, error) {

	if err := c.ShouldBindQuery(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}

	// if err := c.ShouldBindUri(req); err != nil {
	// 	return nil, errors.ParameterError{Message: err.Error()}
	// }
	return req, nil

}

func (req *AllScoreRequestJSON) ToEntity() *entities.AllScoreFilter {
	fmt.Println("req :", req)
	return &entities.AllScoreFilter{
		Name:     req.Name,
		Teamtype: req.TeamType,
		Page:     req.Page,
		Pagesize: req.PageSize,
		Code:     req.Code,
	}
}

type AllScoreResponseJSON struct {
	ID                string              `json:"_id" bson:"id"`
	Title             string              `json:"title" bson:"title"`
	Total             float64             `json:"total" bson:"total"`
	Code              string              `json:"code" bson:"code"`
	No                int                 `json:"no" bson:"no"`
	Teamtype          string              `json:"teamtype" bson:"teamtype" `
	AllConsiderations []AllConsiderations `json:"considerations" bson:"considerations"`
}
type AllConsiderations struct {
	Title string  `json:"title" bson:"title" `
	Score float64 `json:"score" bson:"score" `
}

type AllScoresResponseJSON []AllScoreResponseJSON

func (m *AllScoresResponseJSON) Parse(c *gin.Context, data []entities.AllScore) *AllScoresResponseJSON {

	var allScores AllScoresResponseJSON = AllScoresResponseJSON{}
	// idx := 0
	// total := 0.0
	// for i := 0; i <= len(data); i++ {

	// idx += 1

	// }

	for _, value := range data {

		// fmt.Println("len data ", len(data))
		//
		// for i := 0; i <= len(data); i++ {
		// 	idx := 1

		// 	idx += 1
		// 	fmt.Println("idx :", idx)
		// }

		// if value.Total >= total {
		// 	idx += 1
		// }
		// sort.Sort(sort.Float64Slice{value.Allconsiderations[0].Score})

		// if value.Total == total {
		// 	idx  = idx
		// }

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
			No:                value.No,
			Title:             value.Title,
			Code:              value.Code,
			Teamtype:          value.TeamType,
			Total:             value.Total,
			AllConsiderations: allConsideration,
		}

		allScores = append(allScores, *allScore)

	}

	return &allScores

}
