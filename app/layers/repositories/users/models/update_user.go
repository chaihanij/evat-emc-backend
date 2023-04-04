package models

import (
	"time"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateUser(input *entities.User) *bson.D {
	updateFields := bson.D{
		bson.E{Key: "username", Value: input.Username},
		bson.E{Key: "email", Value: input.Email},
		bson.E{Key: "first_name", Value: input.FirstName},
		bson.E{Key: "last_name", Value: input.LastName},
		bson.E{Key: "address", Value: input.Address},
		bson.E{Key: "tel", Value: input.Tel},
		bson.E{Key: "academy", Value: input.Academy},
		bson.E{Key: "occupation", Value: input.Occupation},
		bson.E{Key: "role", Value: input.Role},
		bson.E{Key: "password", Value: input.Password},
		bson.E{Key: "Year", Value: input.Year},
		bson.E{Key: "team_uuid", Value: input.TeamUUID},
		bson.E{Key: "is_email_verify", Value: input.IsEmailVerify},
		bson.E{Key: "activate_code", Value: input.ActivateCode},
		bson.E{Key: "access_token", Value: input.AccessToken},
		bson.E{Key: "is_active", Value: input.IsActive},
		bson.E{Key: "last_login", Value: input.LastLogin},
		bson.E{Key: "access_token", Value: input.AccessToken},
		bson.E{Key: "updated_at", Value: time.Now()},
		bson.E{Key: "updated_by", Value: input.UpdatedBy},
	}
	update := bson.D{{Key: "$set", Value: updateFields}}
	log.WithField("value", update).Debugln("models.UpdateUser")
	return &update
}
