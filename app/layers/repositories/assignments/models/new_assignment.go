package models

import (
	"time"

	"github.com/google/uuid"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// type Assignment struct {
// 	ID          primitive.ObjectID `bson:"id"`
// 	UUID        string             `bson:"uuid"`
// 	No          int                `bson:"no"`
// 	Title       string             `bson:"title"`
// 	Description string             `bson:"description"`
// 	Image       string             `bson:"image"`
// 	File        string             `bson:"file"`
// 	FullScore   float64            `bson:"full_score"`
// 	IsActive    bool               `bson:"is_active"`
// 	DueDate     time.Time          `bson:"due_date"`
// 	Year        string             `bson:"year"`
// 	CreatedAt   time.Time          `bson:"created_at"`
// 	UpdatedAt   time.Time          `bson:"updated_at"`
// 	CreatedBy   string             `bson:"created_by"`
// 	UpdatedBy   string             `bson:"updated_by"`
// }

func NewAssignment(input *entities.Assignment) *Assignment {
	var image string
	if val, ok := input.Image.(string); ok {
		image = val
	}
	var document string
	if val, ok := input.Document.(string); ok {
		document = val
	}
	now := time.Now()
	return &Assignment{
		ID:          primitive.NewObjectID(),
		UUID:        uuid.NewString(),
		No:          input.No,
		Title:       input.Title,
		Description: input.Description,
		Image:       image,
		Document:    document,
		FullScore:   input.FullScore,
		IsActive:    input.IsActive,
		DueDate:     input.DueDate,
		Year:        input.Year,
		CreatedAt:   now,
		UpdatedAt:   now,
		CreatedBy:   input.CreatedBy,
		UpdatedBy:   input.UpdatedBy,
	}
}
