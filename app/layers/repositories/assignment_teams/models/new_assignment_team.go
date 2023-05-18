package models

import (
	"time"

	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewAssignmentTeam(input *entities.AssignmentTeam) *AssignmentTeam {
	// var documents []string = []string{}
	// if val, ok := input.Documents.([]string); ok {
	// 	documents = val
	// }
	now := time.Now()
	return &AssignmentTeam{
		ID:             primitive.NewObjectID(),
		AssignmentUUID: input.AssignmentUUID,
		TeamUUID:       input.TeamUUID,
		Description:    input.Description,
		IsConfirmed:    false,
		// Documents:      documents,
		Score:          input.Score,
		CreatedAt:      now,
		UpdatedAt:      now,
		CreatedBy:      input.CreatedBy,
		UpdatedBy:      input.UpdatedBy,
	}
}
