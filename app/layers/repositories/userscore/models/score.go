package medles

import (
	"time"

	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Score struct {
	// ID             primitive.ObjectID `json:"_id" bson:"_id"`
	// UID            string             `json:"uid" bson:"uid"`
	// NameTeam       string             `json:"nameteam" bson:"nameteam"`
	// FirstTeam      string             `json:"firstteam" bson:"firstteam"`
	// SecondTeam     string             `json:"secondteam" bson:"secondteam"`
	// First_Stadium  string             `json:"firststadium" bson:"firststadium"`
	// Second_Stadium string             `json:"secondstadium" bson:"secondstadium"`
	// Third_Stadium  string             `json:"thirdstadium" bson:"thirdstadium"`
	// Fourth_Stadium string             `json:"fourthstadium" bson:"fourthstadium"`
	// Sum_Secore     string             `json:"sumscore" bson:"sumscore"`
	// No             string             `json:"no" bson:"no"`
	// CreateDate     time.Time          `json:"createdate" bson:"createdate"`
	// Lastupdate     time.Time          `json:"lastupdate" bson:"lastupdate"`
	// CreateBy       string             `json:"createby" bson:"createby"`
	// LastUpdateBy   string             `json:"lastupdateby" bson:"lastupdateby"`
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	UUID            string             `json:"uuid" bson:"uuid"`
	Field_race_uuid string             `json:"field_race_uuid" bson:"field_race_uuid"`
	Team_uuid       string             `json:"team_uuid" bson:"team_uuid"`
	Score           float64            `json:"score" bson:"score"`
	CreateDate      time.Time          `json:"create_at" bson:"create_at"`
	Lastupdate      time.Time          `json:"update_at" bson:"update_at"`
	CreateBy        string             `json:"create_by" bson:"create_by"`
	UpdateBy        string             `json:"update_by" bson:"update_by"`
}

func (sc *Score) ToEntity() (*entities.Score, error) {
	var _score entities.Score
	err := copier.Copy(&_score, sc)
	return &_score, err
}

type Scores []Score

func (as Scores) ToEntity() []entities.Score {
	var scores []entities.Score
	for _, v := range as {
		score, _ := v.ToEntity()
		scores = append(scores, *score)
	}
	return scores
}
