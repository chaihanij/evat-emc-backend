package models

import (
	"time"

	"github.com/google/uuid"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewMember(input *entities.Member) *Member {
	now := time.Now()
	member := &Member{
		ID:         primitive.NewObjectID(),
		UUID:       uuid.NewString(),
		FirstName:  input.FirstName,
		LastName:   input.LastName,
		Address:    input.Address,
		Email:      input.Email,
		Tel:        input.Tel,
		Academy:    input.Academy,
		Year:       input.Year,
		MemberType: input.MemberType,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
	if val, ok := input.Image.(string); ok {
		member.Image = val
	}
	if val, ok := input.Documents.([]string); ok {
		member.Documents = val
	}
	return member
}
