package models

import (
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson"
)

func FilterUserCriteria(input *entities.UserFilter) *bson.M {
	filter := bson.M{}

	if input.UID != nil {
		filter["uid"] = *input.UID
	}

	if input.Email != nil {
		filter["email"] = *input.Email
	}

	if input.Year != nil {
		filter["year"] = *input.Year
	}

	if input.IsEmailVerify != nil {
		filter["is_email_verify"] = *input.IsEmailVerify
	}

	if input.IsActive != nil {
		filter["is_active"] = *input.IsActive
	}

	if input.ActivateCode != nil {
		filter["activate_code"] = *input.ActivateCode
	}

	if input.AccessToken != nil {
		filter["access_token"] = *input.AccessToken
	}

	return &filter
}
