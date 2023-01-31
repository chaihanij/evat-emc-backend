package models

import (
	"time"

	"github.com/google/uuid"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewTeam(input *entities.Team) *Team {
	var members []string
	if val, ok := input.Members.([]string); ok {
		members = val
	}
	return &Team{
		ID:        primitive.NewObjectID(),
		UUID:      uuid.NewString(),
		Code:      input.Code,
		Name:      input.Name,
		TeamType:  input.TeamType,
		Academy:   input.Academy,
		Detail:    input.Detail,
		Year:      input.Year,
		Members:   members,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: input.CreatedBy,
	}
}
