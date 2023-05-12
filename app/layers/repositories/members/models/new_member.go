package models

import (
	"time"

	"github.com/google/uuid"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewMember(input *entities.Member) *Member {
	now := time.Now()
	is_Leader := false
	if input.IsTeamLeader != false {
		is_Leader = true
	}
	member := &Member{
		ID:           primitive.NewObjectID(),
		UUID:         uuid.NewString(),
		FirstName:    input.FirstName,
		LastName:     input.LastName,
		Address:      input.Address,
		Email:        input.Email,
		Tel:          input.Tel,
		Academy:      input.Academy,
		Major:        input.Major,
		Year:         input.Year,
		TeamUUID:     input.TeamUUID,
		MemberType:   input.MemberType,
		IsTeamLeader: is_Leader,
		Documents:    []string{},
		CreatedAt:    now,
		UpdatedAt:    now,
		CreatedBy:    input.CreatedBy,
		NationalId:   input.NationalId,
		BirthDay:     input.BirthDay,
	}
	if val, ok := input.Image.(string); ok {
		member.Image = val
	}
	if val, ok := input.Documents.([]string); ok {
		member.Documents = val
	}
	return member
}
