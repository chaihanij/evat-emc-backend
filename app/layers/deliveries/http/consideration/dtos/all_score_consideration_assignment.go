package dtos

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

type AllScoreConsiderationAssignmentRequestJSON struct {
	AssignmentUUID string `uri:"assignment_UUID"`
}

func (req *AllScoreConsiderationAssignmentRequestJSON) Parse(c *gin.Context) (*AllScoreConsiderationAssignmentRequestJSON, error) {

	if err := c.ShouldBindUri(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	// if err := c.ShouldBindUri(req); err != nil { //a2cdf901-c976-46ff-98be-ef0fcaf4d0f2
	// 	return nil, errors.ParameterError{Message: err.Error()}
	// }
	return req, nil

}

func (req *AllScoreConsiderationAssignmentRequestJSON) ToEntity() *entities.AllScoreFilter {
	// fmt.Println("req :", req.AssignmentUUID)
	return &entities.AllScoreFilter{
		UUID: req.AssignmentUUID,
	}
}

type AllScoreConsiderationAssignmentResponseJSON []AllScoreResponseJSON

func (m *AllScoreConsiderationAssignmentResponseJSON) Parse(c *gin.Context, data []entities.AllScore) *AllScoreConsiderationAssignmentResponseJSON {

	var allScores AllScoreConsiderationAssignmentResponseJSON = AllScoreConsiderationAssignmentResponseJSON{}
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
