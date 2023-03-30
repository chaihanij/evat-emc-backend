package dtos

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/entities"
)

type CreateEmailRequestJSON struct {
	Email     string `json:"email" bson:"email" `
	Create_at string `json:"create_at" bson:"create_at" `
}

type CreateEmailResponseJSON struct {
	Email     string `json:"email" bson:"email" `
}

func (req *CreateEmailRequestJSON) Parse(c *gin.Context) (*CreateEmailRequestJSON, error) {
	err := c.ShouldBindJSON(req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (req *CreateEmailRequestJSON) ToEntity() *entities.Email {
	return &entities.Email{
		Email:     req.Email,
		Create_at: req.Create_at,
	}
}

func (m *CreateEmailResponseJSON) Parse(data *entities.Email) *CreateEmailResponseJSON {
	email := &CreateEmailResponseJSON{
		Email: data.Email,
	}
	return email
}
