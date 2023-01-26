package models

import (
	"encoding/base64"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	UID           string             `bson:"uid"`
	Username      string             `bson:"username"`
	Email         string             `bson:"email"`
	FirstName     string             `bson:"first_name"`
	LastName      string             `bson:"last_name"`
	Address       string             `bson:"address"`
	Tel           string             `bson:"tel"`
	UserType      string             `bson:"user_type"`
	Role          string             `bson:"role"`
	Password      string             `bson:"password"`
	Year          string             `bson:"year"`
	TeamUID       string             `bson:"team_uid"`
	IsEmailVerify bool               `bson:"is_email_verify"`
	ActivateCode  string             `bson:"activate_code"`
	AccessToken   string             `bson:"access_token"`
	IsActive      bool               `bson:"is_active"`
	LastLogin     time.Time          `bson:"last_login"`
	CreatedAt     time.Time          `bson:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at"`
}

func (u *User) ToEntity() (*entities.User, error) {
	var user entities.User
	err := copier.Copy(&user, u)
	return &user, err
}

type Users []User

func (u Users) ToEntiire() ([]entities.User, error) {
	var users []entities.User
	for _, v := range u {
		user, err := v.ToEntity()
		if err != nil {
			return nil, err
		}
		users = append(users, *user)
	}
	return users, nil
}

func (u *User) ParseToMinimalModel(e *entities.UserMinimalCreate) *User {
	hash, _ := bcrypt.GenerateFromPassword([]byte(e.Password), bcrypt.DefaultCost)
	password := string(hash)
	u.ID = primitive.NewObjectID()
	u.UID = uuid.NewString()
	u.Email = e.Email
	u.Password = password
	now := time.Now()
	u.CreatedAt = now
	u.UpdatedAt = now
	return u
}

func (u *User) ParseToModel(e *entities.UserCreate) *User {
	now := time.Now()
	user := &User{
		ID:            primitive.NewObjectID(),
		UID:           uuid.NewString(),
		Username:      e.Username,
		Email:         e.Email,
		IsEmailVerify: false,
		ActivateCode:  base64.StdEncoding.EncodeToString([]byte(uuid.New().String())),
		IsActive:      false,
		CreatedAt:     now,
		UpdatedAt:     now,
	}
	if e.FirstName != nil {
		user.FirstName = *e.FirstName
	}
	if e.LastName != nil {
		user.LastName = *e.LastName
	}
	if e.Address != nil {
		user.Address = *e.Address
	}
	if e.Tel != nil {
		user.Tel = *e.Tel
	}
	if e.Role != nil {
		role := *e.Role
		user.Role = string(role)
	}
	if e.Password != nil {
		hash, _ := bcrypt.GenerateFromPassword([]byte(*e.Password), bcrypt.DefaultCost)
		user.Password = string(hash)
	}
	if e.Year != nil {
		user.Year = *e.Year
	}
	if e.TeamUID != nil {
		user.TeamUID = *e.TeamUID
	}
	if e.IsEmailVerify != nil {
		user.IsEmailVerify = *e.IsEmailVerify
	}
	if e.IsActive != nil {
		user.IsActive = *e.IsActive
	}
	return user
}

func ToUpdateUserFields(input *entities.UserPartialUpdate) *bson.D {
	updateFields := bson.D{
		bson.E{Key: "updated_at", Value: time.Now()},
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
	if input.Role != nil {
		updateFields = append(updateFields, bson.E{Key: "role", Value: input.Role})
	}
	if input.Password != nil {
		hash, _ := bcrypt.GenerateFromPassword([]byte(*input.Password), bcrypt.DefaultCost)
		updateFields = append(updateFields, bson.E{Key: "password", Value: string(hash)})
	}
	if input.TeamUID != nil {
		updateFields = append(updateFields, bson.E{Key: "team_uid", Value: input.TeamUID})
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
	statement := bson.D{{Key: "$set", Value: updateFields}}
	return &statement
}
