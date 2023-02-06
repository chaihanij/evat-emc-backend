package models

import (
	"time"

	"github.com/google/uuid"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewAlbum(input *entities.Album) *Album {
	now := time.Now()
	return &Album{
		ID:        primitive.NewObjectID(),
		UUID:      uuid.NewString(),
		Title:     input.Title,
		Year:      input.Year,
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: input.CreatedBy,
		UpdatedBy: input.UpdatedBy,
	}
}
