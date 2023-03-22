package models

import (
	"time"

	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/entities"
)

// type Consideration struct {
// 	ID        string    `json:"_id" bson:"_id"`
// 	Score     float64   `score:"score" bson:"score"`
// 	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
// 	No        int    `json:"no"`
// }

type Consideration struct {
	ID             string           `json:"_id" bson:"_id"`
	TotalScore     float64          `json:"total_score" bson:"total_score" `
	UpdatedAt      time.Time        `json:"update_at"  bson:"updated_at" `
	No             int              `json:"no" bson:"no"`
	IndivdualScore []IndivdualScore `json:"indivdual_score" bson:"indivdual_score" `
}

type IndivdualScore struct {
	Title  string  `json:"title"`
	Score float64 `json:"score"`
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
