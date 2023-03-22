package models

import (
	"time"

	"github.com/google/uuid"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewFieldRaceTeam(input *entities.FieldRaceTeam) *FieldRaceTeam {
	// var fieldRaces []FieldRaces
	// for _, valueFieldRace := range input.FieldRaces {
	// 	fieldRace := &FieldRaces{
	// 		Title:       valueFieldRace.Title,
	// 		Description: valueFieldRace.Description,
	// 		File:        valueFieldRace.File,
	// 		Image:       valueFieldRace.Image,
	// 		Year:        valueFieldRace.Year,
	// 		FullScore:   valueFieldRace.FullScore,
	// 	}
	// 	fieldRaces = append(fieldRaces, *fieldRace)
	// }
	return &FieldRaceTeam{
		ID:            primitive.NewObjectID(),
		UUID:          uuid.NewString(),
		FieldRaceUUID: input.FieldRaceUUID,
		TeamUUID:      input.TeamUUID,
		Score:         input.Score,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		CreatedBy:     input.CreatedBy,
		UpdatedBy:     input.UpdatedBy,
		// FieldRaces:    fieldRaces,
	}

}

func NewFieldRaces(input *entities.FieldRaces) *FieldRaces {
	var fieldRaces []FieldRaces
	for _, valueFieldRace := range fieldRaces {
		fieldRace := &FieldRaces{
			Title:       valueFieldRace.Title,
			Description: valueFieldRace.Description,
			File:        valueFieldRace.File,
			Image:       valueFieldRace.Image,
			Year:        valueFieldRace.Year,
			FullScore:   valueFieldRace.FullScore,
		}
		fieldRaces = append(fieldRaces, *fieldRace)
	}
	return &FieldRaces{}
}
