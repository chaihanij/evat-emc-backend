package medles

import (
	"time"

	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewScore(input *entities.Score) *Score {

	now := time.Now()

	return &Score{
		ID: primitive.NewObjectID(),
		// NameTeam:       input.NameTeam,
		// FirstTeam:      input.FirstTeam,
		// SecondTeam:     input.SecondTeam,
		// First_Stadium:  input.First_Stadium,
		// Second_Stadium: input.Second_Stadium,
		// Third_Stadium:  input.Third_Stadium,
		// Fourth_Stadium: input.Fourth_Stadium,
		// Sum_Secore:     input.Sum_Score,
		// No:             input.No,
		UUID:            input.UUID,
		Field_race_uuid: input.Field_race_uuid,
		Team_uuid:       input.Team_uuid,
		Score:           input.Score,
		CreateDate:      now,
		Lastupdate:      now,
		CreateBy:        input.CreateBy,
		UpdateBy:        input.Updated_by,
		// LastUpdateBy:input.Updated_by,
		// LastUpdateBy:   input.LastUpdateBy,
	}

}
