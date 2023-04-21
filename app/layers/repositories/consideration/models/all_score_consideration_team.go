package models

import (
	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/entities"
)

type AllScoreConsideration struct {
	ID                string              `json:"team" bson:"team"`
	Title             string              `json:"title" bson:"title" `
	Total             float64             `json:"total" bson:"total"`
	Code              string              `json:"code" bson:"code" `
	No                int                 `json:"no" bson:"no"`
	AllConsiderations []AllConsiderations `json:"considerations" bson:"considerations"`
}

type AllConsiderations struct {
	Title string  `json:"title" bson:"title" `
	Score float64 `json:"total" bson:"total" `
}

func (a *AllScoreConsideration) ToEntity() (*entities.AllScore, error) {
	var allScore entities.AllScore
	err := copier.Copy(&allScore, a)
	return &allScore, err
}

type AllScoreConsiderations []AllScoreConsideration

func (as AllScoreConsiderations) ToEntity() []entities.AllScore {
	var allScoreConsiderations []entities.AllScore

	for _, value := range as {

		var allScoresconsiderations []entities.AllConsideration

		for _, vl := range value.AllConsiderations {
			allScoreconsideration := entities.AllConsideration{
				Title: vl.Title,
				Score: vl.Score,
			}
			allScoresconsiderations = append(allScoresconsiderations, allScoreconsideration)
		}

		alsc := entities.AllScore{
			ID:                value.ID,
			Title:             value.Title,
			Total:             value.Total,
			Code:              value.Code,
			No:                value.No,
			Allconsiderations: allScoresconsiderations,
		}

		// allScore, _ := value.ToEntity()
		allScoreConsiderations = append(allScoreConsiderations, alsc)
	}
	return allScoreConsiderations
}
