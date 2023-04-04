package models

import (
	"encoding/base64"
	"time"

	"github.com/google/uuid"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func NewUser(input *entities.User) *User {
	hash, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	password := string(hash)
	now := time.Now()
	return &User{
		ID:            primitive.NewObjectID(),
		UID:           uuid.NewString(),
		Username:      input.Username,
		Email:         input.Email,
		FirstName:     input.FirstName,
		LastName:      input.LastName,
		Address:       input.Address,
		Tel:           input.Tel,
		Academy:       input.Academy,
		Occupation:    input.Occupation,
		Role:          input.Role,
		Password:      password,
		Year:          input.Year,
		TeamUUID:      input.TeamUUID,
		IsEmailVerify: false,
		ActivateCode:  base64.StdEncoding.EncodeToString([]byte(uuid.New().String())),
		IsActive:      false,
		CreatedAt:     now,
		UpdatedAt:     now,
		CreatedBy:     input.CreatedBy,
	}
}
