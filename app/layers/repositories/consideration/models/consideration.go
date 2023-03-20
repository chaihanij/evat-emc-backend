package models

import (
	"time"

	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/entities"
)

type Consideration struct {
	ID        string    `json:"_id" bson:"_id"`
	Score     float64   `score:"score" bson:"score"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	No        int    `json:"no"`
}

func (at *Consideration) ToEntity() (*entities.Consideration, error) {
	var consideration entities.Consideration
	err := copier.Copy(&consideration, at)
	return &consideration, err
}

type Considerations []Consideration

func (fTeams Considerations) ToEntity() []entities.Consideration {
	var considerations []entities.Consideration
	for _, v := range fTeams {
		fieldRaceTeam, _ := v.ToEntity()
		considerations = append(considerations, *fieldRaceTeam)
	}
	return considerations
}
