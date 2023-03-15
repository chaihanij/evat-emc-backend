package dtos

import (
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

type FindAllScoreRequesJSON struct {
	Page     *int64 `form:"page" validate:"omitempty,gte=1" example:"1"`
	PageSize *int64 `form:"pageSize" validate:"omitempty,gte=1" example:"20"`
}

func (req *FindAllScoreRequesJSON) Parse(c *gin.Context) (*FindAllScoreRequesJSON, error) {

	err := c.ShouldBindQuery(req)
	if err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	return req, nil
}

func (req *FindAllScoreRequesJSON) ToEntity() *entities.ScoreFilter {
	return &entities.ScoreFilter{
		// Year:     req.Year,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
}

type FindAllScoreResponseJSON []ScoreResponse

func (m *FindAllScoreResponseJSON) Parse(data []entities.Score) *FindAllScoreResponseJSON {
	var scores FindAllScoreResponseJSON = FindAllScoreResponseJSON{}
	for _, value := range data {
		score := &ScoreResponse{
			// UID:            value.UUID,
			NameTeam:       value.NameTeam,
			FirstTeam:      value.FirstTeam,
			SecondTeam:     value.SecondTeam,
			First_Stadium:  value.First_Stadium,
			Second_Stadium: value.Second_Stadium,
			Third_Stadium:  value.Third_Stadium,
			Fourth_Stadium: value.Fourth_Stadium,
			Sum_Score:      value.Sum_Score,
			No:             value.No,
			CreateDate:     time.Now(),
			Lastupdate:     time.Now(),
			CreateBy:       value.CreateBy,
			LastUpdateBy:   value.LastUpdateBy,
		}
		scores = append(scores, *score)
	}
	return &scores
}
