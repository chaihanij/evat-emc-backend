package models

import (
	"time"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson"
)

func PartialUpdateMember(input *entities.MemberPartialUpdate) *bson.D {

	updateFields := bson.D{
		bson.E{Key: "firstname", Value: input.FirstName},
		bson.E{Key: "lastname", Value: input.LastName},
		bson.E{Key: "address", Value: input.Address},
		bson.E{Key: "email", Value: input.Email},
		bson.E{Key: "tel", Value: input.Tel},
		bson.E{Key: "academy", Value: input.Academy},
		bson.E{Key: "year", Value: input.Year},
		bson.E{Key: "member_type", Value: input.MemberType},
		bson.E{Key: "is_team_leader", Value: input.IsTeamLeader},
		bson.E{Key: "team_uuid", Value: input.TeamUUID},
		bson.E{Key: "updated_at", Value: time.Now()},
	}

	if input.FirstName != nil {
		updateFields = append(updateFields, bson.E{Key: "firstname", Value: input.FirstName})
	}

	if input.LastName != nil {
		updateFields = append(updateFields, bson.E{Key: "lastname", Value: input.LastName})
	}

	if input.Address != nil {
		updateFields = append(updateFields, bson.E{Key: "address", Value: input.Academy})
	}

	if input.Email != nil {
		updateFields = append(updateFields, bson.E{Key: "email", Value: input.Email})
	}

	if input.Tel != nil {
		updateFields = append(updateFields, bson.E{Key: "tel", Value: input.Tel})
	}

	if input.Academy != nil {
		updateFields = append(updateFields, bson.E{Key: "academy", Value: input.Academy})
	}

	if input.Year != nil {
		updateFields = append(updateFields, bson.E{Key: "year", Value: input.Year})
	}

	if input.MemberType != nil {
		updateFields = append(updateFields, bson.E{Key: "member_type", Value: input.MemberType})
	}
	if input.IsTeamLeader != nil {
		updateFields = append(updateFields,
			bson.E{Key: "is_team_leader", Value: input.IsTeamLeader})
	}
	if input.TeamUUID != nil {
		updateFields = append(updateFields, bson.E{Key: "team_uuid", Value: input.TeamUUID})
	}

	if val, ok := input.Image.(*string); ok {
		updateFields = append(updateFields, bson.E{Key: "image", Value: val})
	}

	if val, ok := input.Documents.(*[]string); ok {
		updateFields = append(updateFields, bson.E{Key: "document", Value: val})
	}
	update := bson.D{{Key: "$set", Value: updateFields}}
	log.WithField("value", update).Debugln("models.PartialUpdateMember")
	return &update
}
