package models

import (
	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson"
)

func NewUserFilter(input interface{}) *bson.M {
	var filter bson.M = bson.M{}
	if val, ok := input.(*entities.User); ok {
		if val.UID != "" {
			filter["uid"] = val.UID
		}
	}
	if val, ok := input.(*entities.UserPartialUpdate); ok {
		if val.UID != "" {
			filter["uid"] = val.UID
		}
	}
	if val, ok := input.(*entities.UserFilter); ok {
		if val.UID != nil {
			filter["uid"] = val.UID
		}
		if val.Email != nil {
			filter["email"] = val.Email
		}
		if val.Year != nil {
			filter["year"] = val.Year
		}
		if val.IsEmailVerify != nil {
			filter["is_email_verify"] = val.IsEmailVerify
		}
		if val.IsActive != nil {
			filter["is_active"] = val.IsActive
		}
		if val.ActivateCode != nil {
			filter["activate_code"] = val.ActivateCode
		}
		if val.IsActive != nil {
			filter["is_active"] = val.IsActive
		}
	}
	log.WithField("value", filter).Debugln("models.NewUserFilter")
	return &filter
}
