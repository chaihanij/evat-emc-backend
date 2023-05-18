package models

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateFileMemberFilter(input string) *bson.M {
	var filter bson.M = bson.M{}

	if input != "" {
		filter["uuid"] = input
	}
	log.WithField("value", filter).Debugln("models.NewAssignmentFilter")
	return &filter
}

func UploadFileMember(input *entities.MemberUpdatePDF) *bson.M {

	// salt := make([]byte, 10)
	// rand.Read(salt)
	// saltString := base64.URLEncoding.EncodeToString(salt)
	modifiedLink := fmt.Sprintf("%s/v1/files/%s", env.BaseUrl, input.Files)

	// u, err := url.Parse(originalLink)
	// if err != nil {
	// 	fmt.Println("Error parsing URL:", err)

	// }
	// u.Path = u.Path + saltString

	
	// modifiedLink := u.String()

	updateFields := bson.M{
		"$set": bson.M{
			"files": modifiedLink,
		},
	}

	return &updateFields
}
