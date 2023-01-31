package models

import (
	"time"

	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Assignment struct {
	ID          primitive.ObjectID `bson:"id"`
	UUID        string             `bson:"uuid"`
	No          int                `bson:"no"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Image       string             `bson:"image"`
	Document    string             `bson:"document"`
	FullScore   float64            `bson:"full_score"`
	IsActive    bool               `bson:"is_active"`
	DueDate     time.Time          `bson:"due_date"`
	Year        string             `bson:"year"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
	CreatedBy   string             `bson:"created_by"`
	UpdatedBy   string             `bson:"updated_by"`
}

func (am *Assignment) ToEntity() (*entities.Assignment, error) {
	var assignment entities.Assignment
	err := copier.Copy(&assignment, am)
	return &assignment, err
}

type Assignments []Assignment

func (as Assignments) ToEntity() []entities.Assignment {
	var assignments []entities.Assignment
	for _, v := range as {
		assignment, _ := v.ToEntity()
		assignments = append(assignments, *assignment)
	}
	return assignments
}
