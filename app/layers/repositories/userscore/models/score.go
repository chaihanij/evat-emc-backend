package medles

import (
	"time"

	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Score struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id"`
	UID            string             `json:"uid" bson:"uid"`
	NameTeam       string             `json:"nameteam" bson:"nameteam"`
	FirstTeam      string             `json:"firstteam" bson:"firstteam"`
	SecondTeam     string             `json:"secondteam" bson:"secondteam"`
	First_Stadium  string             `json:"firststadium" bson:"firststadium"`
	Second_Stadium string             `json:"secondstadium" bson:"secondstadium"`
	Third_Stadium  string             `json:"thirdstadium" bson:"thirdstadium"`
	Fourth_Stadium string             `json:"fourthstadium" bson:"fourthstadium"`
	Sum_Secore     string             `json:"sumscore" bson:"sumscore"`
	No             string             `json:"no" bson:"no"`
	CreateDate     time.Time          `json:"createdate" bson:"createdate"`
	Lastupdate     time.Time          `json:"lastupdate" bson:"lastupdate"`
	CreateBy       string             `json:"createby" bson:"createby"`
	LastUpdateBy   string             `json:"lastupdateby" bson:"lastupdateby"`
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
