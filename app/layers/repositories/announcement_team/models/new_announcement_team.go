package models

import (
	"time"

	"github.com/google/uuid"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewAnnouncementTeam(input *entities.AnnouncementTeam) *AnnouncementTeam {
	return &AnnouncementTeam{
		ID:               primitive.NewObjectID(),
		UUID:             uuid.NewString(),
		AnnouncementUUID: input.AnnouncementUUID,
		TeamUUID:         input.TeamUUID,
		IsRead:           false,
		CreatedAt:        time.Now(),
	}
}
