package models

import (
	"fmt"

	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/entities"
)

type AllScoreConsideration struct {
	ID                string              `json:"_id" bson:"_id"`
	Title             string              `json:"title" bson:"title" `
	Total             float64             `json:"total" bson:"total"`
	AllConsiderations []AllConsiderations `json:"considerations" bson:"considerations"`
}

type AllConsiderations struct {
	Title string  `json:"title" bson:"title" `
	Score float64 `json:"score" bson:"score" `
}

func (a *AllScoreConsideration) ToEntity() (*entities.AllScore, error) {
	var allScore entities.AllScore
	err := copier.Copy(&allScore, a)
	return &allScore, err
}

type AllScoreConsiderations []AllScoreConsideration

func (as AllScoreConsiderations) ToEntity() []entities.AllScore {
	var allScoreConsiderations []entities.AllScore

	fmt.Println("as", as)
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
			Allconsiderations: allScoresconsiderations,
		}

		fmt.Println("alsc ;", alsc)

		// allScore, _ := value.ToEntity()
		allScoreConsiderations = append(allScoreConsiderations, alsc)
	}
	fmt.Println("allScoreConsiderations :", allScoreConsiderations)
	return allScoreConsiderations
}
