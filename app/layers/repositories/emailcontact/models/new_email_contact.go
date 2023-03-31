package models

import (
	"time"

	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// type

type CreateEmailContact struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Title       string             `json:"title" bson:"title" `
	Email       string             `json:"email" bson:"email"`
	FirstName   string             `json:"firstname" bson:"firstname" `
	LastName    string             `json:"lastname" bson:"lastname" `
	Description string             `json:"description" bson:"description"`
	Create_at   time.Time          `json:"create_at" bson:"create_at" `
	Status      bool               `json:"status" bson:"status" `
}

func (at *CreateEmailContact) ToEntity() (*entities.CreateContactEmail, error) {
	var emailcontact entities.CreateContactEmail
	err := copier.Copy(&emailcontact, at)

	return &emailcontact, err

}

func NewEmailContact(data *entities.CreateContactEmail) *CreateEmailContact {
	return &CreateEmailContact{
		ID:          primitive.NewObjectID(),
		Email:       data.Email,
		Title:       data.Title,
		Description: data.Description,
		FirstName:   data.FirstName,
		LastName:    data.LastName,
		Create_at:   time.Now(),
		Status:      false,
	}
}
