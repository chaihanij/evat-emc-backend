package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/entities"
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
	Role          string             `bson:"role"`
	Password      string             `bson:"password"`
	Year          string             `bson:"year"`
	TeamUUID      string             `bson:"team_uuid"`
	IsEmailVerify bool               `bson:"is_email_verify"`
	ActivateCode  string             `bson:"activate_code"`
	AccessToken   string             `bson:"access_token"`
	IsActive      bool               `bson:"is_active"`
	LastLogin     time.Time          `bson:"last_login"`
	CreatedAt     time.Time          `bson:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at"`
	CreatedBy     string             `bson:"created_by"`
	UpdatedBy     string             `bson:"updated_by"`
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
