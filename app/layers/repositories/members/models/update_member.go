package models

import (
	"time"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateMember(input *entities.Member) *bson.D {
	updateFields := bson.D{
		bson.E{Key: "firstname", Value: input.FirstName},
		bson.E{Key: "lastname", Value: input.LastName},
		bson.E{Key: "address", Value: input.Address},
		bson.E{Key: "email", Value: input.Email},
		bson.E{Key: "tel", Value: input.Tel},
		bson.E{Key: "academy", Value: input.Academy},
		bson.E{Key: "major", Value: input.Major},
		bson.E{Key: "year", Value: input.Year},
		bson.E{Key: "member_type", Value: input.Year},
		bson.E{Key: "is_team_leader", Value: input.IsTeamLeader},
		bson.E{Key: "team_uuid", Value: input.TeamUUID},
		bson.E{Key: "updated_at", Value: time.Now()},
		bson.E{Key: "birth_day", Value: input.BirthDay},
		bson.E{Key: "national_id", Value: input.NationalId},
	}

	if val, ok := input.Image.(string); ok {
		updateFields = append(updateFields, bson.E{Key: "image", Value: val})
	}
	if val, ok := input.Documents.([]string); ok {
		updateFields = append(updateFields, bson.E{Key: "documents", Value: val})
	}
	update := bson.D{{Key: "$set", Value: updateFields}}
	log.WithField("value", update).Debugln("models.UpdateMember")
	return &update
}
