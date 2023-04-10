package models

import (
	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/entities"
)

type FieldRaceTeamScore struct {
	Considerations []Consideration `json:"considerations" bson:"considerations" `
	Total          float64         `json:"total" bson:"total"`
}

type FieldRaceTeamScores []FieldRaceTeamScore

func (f *FieldRaceTeamScore) ToEntity() (*entities.FieldRaceTeamScore, error) {
	var FieldRaceTeamScores entities.FieldRaceTeamScore
	err := copier.Copy(&FieldRaceTeamScores, f)
	return &FieldRaceTeamScores, err
}

func (as FieldRaceTeamScores) ToEntity() []entities.FieldRaceTeamScore {
	var FieldRaceTeamScores []entities.FieldRaceTeamScore
	for _, v := range as {
		fieldraceteam, _ := v.ToEntity()
		FieldRaceTeamScores = append(FieldRaceTeamScores, *fieldraceteam)
	}
	return FieldRaceTeamScores
}
