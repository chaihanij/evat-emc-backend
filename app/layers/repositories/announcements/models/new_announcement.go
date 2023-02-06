package models

import (
	"time"

	"github.com/google/uuid"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewAnnouncement(input *entities.Announcement) *Announcement {
	now := time.Now()
	return &Announcement{
		ID:          primitive.NewObjectID(),
		UUID:        uuid.NewString(),
		Title:       input.Title,
		Description: input.Description,
		Year:        input.Year,
		CreatedAt:   now,
		UpdatedAt:   now,
		CreatedBy:   input.CreatedBy,
		UpdatedBy:   input.UpdatedBy,
	}
}
