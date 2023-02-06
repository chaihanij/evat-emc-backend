package models

import (
	"time"

	"github.com/google/uuid"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// type AssignmentTeam struct {
// 	ID             primitive.ObjectID `bson:"_id"`
// 	UUID           string             `bson:"uuid"`
// 	AssignmentUUID string             `bson:"assignment_uuid"`
// 	TeamUUID       string             `bson:"team_uuid"`
// 	Description    string             `bson:"uuid"`
// 	Documents      []string           `bson:"documents"`
// 	Score          float64            `bson:"ccore"`
// 	CreatedAt      time.Time          `bson:"created_at"`
// 	UpdatedAt      time.Time          `bson:"updated_at"`
// 	CreatedBy      string             `bson:"created_by"`
// 	UpdatedBy      string             `bson:"update_by"`
// }

func NewAssignmentTeam(input *entities.AssignmentTeam) *AssignmentTeam {
	var documents []string = []string{}
	if val, ok := input.Documents.([]string); ok {
		documents = val
	}
	now := time.Now()
	return &AssignmentTeam{
		ID:             primitive.NewObjectID(),
		UUID:           uuid.NewString(),
		AssignmentUUID: input.AssignmentUUID,
		TeamUUID:       input.TeamUUID,
		Description:    input.Description,
		Documents:      documents,
		Score:          input.Score,
		CreatedAt:      now,
		UpdatedAt:      now,
		CreatedBy:      input.CreatedBy,
		UpdatedBy:      input.UpdatedBy,
	}
}
