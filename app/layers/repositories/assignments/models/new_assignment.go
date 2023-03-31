package models

import (
	"time"

	"github.com/google/uuid"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewAssignment(input *entities.Assignment) *Assignment {
	var image string
	if val, ok := input.Image.(string); ok {
		image = val
	}
	var document string
	if val, ok := input.Document.(string); ok {
		document = val
	}
	now := time.Now()
	return &Assignment{
		ID:          primitive.NewObjectID(),
		UUID:        uuid.NewString(),
		TeamUUID:    input.TeamUUID,
		No:          input.No,
		Title:       input.Title,
		Description: input.Description,
		Image:       image,
		Document:    document,
		FullScore:   input.FullScore,
		IsActive:    input.IsActive,
		DueDate:     input.DueDate,
		Year:        input.Year,
		CreatedAt:   now,
		UpdatedAt:   now,
		CreatedBy:   input.CreatedBy,
		UpdatedBy:   input.UpdatedBy,
		SendDoc:     input.SendDoc,
	}
}
