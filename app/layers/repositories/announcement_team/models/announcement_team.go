package models

import (
	"time"

	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AnnouncementTeam struct {
	ID               primitive.ObjectID `bson:"_id"`
	UUID             string             `bson:"uuid"`
	AnnouncementUUID string             `bson:"announcement_uuid"`
	TeamUUID         string             `bson:"team_uuid"`
	IsRead           bool               `bson:"is_read"`
	CreatedAt        time.Time          `bson:"created_at"`
}

func (am *AnnouncementTeam) ToEntity() (*entities.AnnouncementTeam, error) {
	var announcement entities.AnnouncementTeam
	err := copier.Copy(&announcement, am)
	return &announcement, err
}

type AnnouncementTeams []AnnouncementTeam

func (aTeams AnnouncementTeams) ToEntity() []entities.AnnouncementTeam {
	var announcementTeams []entities.AnnouncementTeam
	for _, v := range aTeams {
		announcementTeam, _ := v.ToEntity()
		announcementTeams = append(announcementTeams, *announcementTeam)
	}
	return announcementTeams
}
