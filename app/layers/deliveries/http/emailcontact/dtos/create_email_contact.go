package dtos

import (
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/entities"
)

type CreateEmailContactRequestJSON struct {
	Title       string    `json:"title" bson:"title" `
	Email       string    `json:"email" bson:"email"`
	FirstName   string    `json:"firstname" bson:"firstname"`
	LastName    string    `json:"lastname" bson:"lastname" `
	Description string    `json:"description" bson:"description"`
	Create_at   time.Time `json:"create_at" bson:"create_at" `
	Status      bool      `json:"status" bson:"status" `
}

type CreateEmailContactResponseJSON struct {
	Title       string    `json:"title" bson:"title" `
	Email       string    `json:"email" bson:"email"`
	FirstName   string    `json:"firstname" bson:"firstname"`
	LastName    string    `json:"lastname" bson:"lastname" `
	Description string    `json:"description" bson:"description"`
	Create_at   time.Time `json:"create_at" bson:"create_at" `
	Status      bool      `json:"status" bson:"status" `
}

func (req *CreateEmailContactRequestJSON) Parse(c *gin.Context) (*CreateEmailContactRequestJSON, error) {
	err := c.ShouldBindJSON(req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (req *CreateEmailContactRequestJSON) ToEntity() *entities.CreateContactEmail {
	return &entities.CreateContactEmail{
		Title:       req.Title,
		Email:       req.Email,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Description: req.Description,
		Create_at:   req.Create_at,
	}
}

func (m *CreateEmailContactResponseJSON) Parse(data *entities.CreateContactEmail) *CreateEmailContactResponseJSON {
	email := &CreateEmailContactResponseJSON{
		Title:       data.Title,
		Email:       data.Email,
		FirstName:   data.FirstName,
		LastName:    data.LastName,
		Description: data.Description,
		Create_at:   data.Create_at,
	}
	return email
}
