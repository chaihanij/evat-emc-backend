package models

import (
	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/entities"
)

type AllScoreConsideration struct {
	ID                string              `json:"_id" bson:"name"`
	Name              string              `json:"name" bson:"name"`
	AllConsiderations []AllConsiderations `json:"considerations" bson:"considerations"`
}

type AllConsiderations struct {
	Title string  `json:"title" bson:"title" `
	Score float64 `json:"score" bson:"score" `
	Type  string  `json:"type" bson:"type"`
}

func (a *AllScoreConsideration) ToEntity() (*entities.AllScore, error) {
	var allScore entities.AllScore
	err := copier.Copy(&allScore, a)
	return &allScore, err
}

type AllScoreConsiderations []AllScoreConsideration

func(as AllScoreConsiderations) ToEntity() []entities.AllScore {
	var allScoreConsiderations []entities.AllScore
	for _, value := range as {
		allScore, _ := value.ToEntity()
		allScoreConsiderations = append(allScoreConsiderations, *allScore)
	}
	return allScoreConsiderations
}