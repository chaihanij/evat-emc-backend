package models

import (
	"time"

	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Member struct {
	ID             primitive.ObjectID `bson:"_id"`
	UUID           string             `bson:"uuid"`
	FirstName      string             `bson:"firstname"`
	LastName       string             `bson:"lastname"`
	Address        string             `bson:"address"`
	Email          string             `bson:"email"`
	Tel            string             `bson:"tel"`
	Academy        string             `bson:"academy"`
	Major          string             `bson:"major"`
	Year           string             `bson:"year"`
	MemberType     string             `bson:"member_type"`
	IsTeamLeader   bool               `bson:"is_team_leader"`
	TeamUUID       string             `bson:"team_uuid"`
	Image          string             `bson:"image"`
	Documents      []string           `bson:"documents"`
	CreatedAt      time.Time          `bson:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at"`
	CreatedBy      string             `bson:"created_by"`
	UpdatedBy      string             `bson:"updated_by"`
	BirthDay       time.Time          `bson:"birth_day" `
	NationalId     string             `bson:"national_id" `
	Is_checkin     bool               `bson:"is_checkin"`
	Checkin_date   time.Time          `bson:"checkin_date"`
	Check_national bool               `bson:"is_national"`
	Is_check_data  bool               `bson:"is_data"`
	Is_Check_image bool               `bson:"is_image"`
	Prefix         string             `bson:"prefix"`
}

func (m *Member) ToEntity() (*entities.Member, error) {
	var member entities.Member
	err := copier.Copy(&member, m)
	return &member, err
}

type Members []Member

func (ms Members) ToEntity() []entities.Member {
	var members []entities.Member
	for _, v := range ms {
		member, _ := v.ToEntity()
		members = append(members, *member)

	}

	return members
}
