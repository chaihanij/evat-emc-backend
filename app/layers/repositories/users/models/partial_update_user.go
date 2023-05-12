package models

import (
	"time"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func PartialUpdateUser(input *entities.UserPartialUpdate) *bson.D {
	updateFields := bson.D{
		bson.E{Key: "updated_at", Value: time.Now()},
		bson.E{Key: "updated_by", Value: input.UpdatedBy},
	}
	if input.Username != nil {
		updateFields = append(updateFields, bson.E{Key: "username", Value: input.Username})
	}
	if input.Email != nil {
		updateFields = append(updateFields, bson.E{Key: "email", Value: input.Email})
	}
	if input.FirstName != nil {
		updateFields = append(updateFields, bson.E{Key: "first_name", Value: input.FirstName})
	}
	if input.LastName != nil {
		updateFields = append(updateFields, bson.E{Key: "last_name", Value: input.LastName})
	}
	if input.Address != nil {
		updateFields = append(updateFields, bson.E{Key: "address", Value: input.Address})
	}
	if input.Tel != nil {
		updateFields = append(updateFields, bson.E{Key: "tel", Value: input.Tel})
	}
	if input.Academy != nil {
		updateFields = append(updateFields, bson.E{Key: "academy", Value: input.Academy})
	}
	if input.Occupation != nil {
		updateFields = append(updateFields, bson.E{Key: "occupation", Value: input.Occupation})
	}
	if input.Role != nil {
		updateFields = append(updateFields, bson.E{Key: "role", Value: input.Role})
	}
	if input.Password != nil {
		hash, _ := bcrypt.GenerateFromPassword([]byte(*input.Password), bcrypt.DefaultCost)
		updateFields = append(updateFields, bson.E{Key: "password", Value: string(hash)})
	}
	if input.TeamUUID != nil {
		updateFields = append(updateFields, bson.E{Key: "team_uuid", Value: input.TeamUUID})
	}
	if input.Year != nil {
		updateFields = append(updateFields, bson.E{Key: "year", Value: input.Year})
	}
	if input.IsEmailVerify != nil {
		updateFields = append(updateFields, bson.E{Key: "is_email_verify", Value: input.IsEmailVerify})
	}
	if input.ActivateCode != nil {
		updateFields = append(updateFields, bson.E{Key: "activate_code", Value: input.ActivateCode})
	}
	if input.AccessToken != nil {
		updateFields = append(updateFields, bson.E{Key: "access_token", Value: input.AccessToken})
	}
	if input.IsActive != nil {
		updateFields = append(updateFields, bson.E{Key: "is_active", Value: input.IsActive})
	}
	if input.LastLogin != nil {
		updateFields = append(updateFields, bson.E{Key: "last_login", Value: input.LastLogin})
	}
	if input.Prefix != nil {
		updateFields = append(updateFields, bson.E{Key: "prefix", Value: input.Prefix})
	}

	update := bson.D{{Key: "$set", Value: updateFields}}
	log.WithField("value", update).Debugln("models.PartialUpdateUser")
	return &update
}
