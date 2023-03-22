package models

import (
	"time"

	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Team struct {
	ID        primitive.ObjectID `bson:"_id"`
	UUID      string             `bson:"uuid"`
	Code      string             `bson:"code"`
	Name      string             `bson:"name"`
	TeamType  string             `bson:"team_type"`
	Academy   string             `bson:"academy"`
	Detail    string             `bson:"detail"`
	Members   []string           `bson:"members"`
	Year      string             `bson:"year"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	CreatedBy string             `bson:"created_by"`
	UpdatedBy string             `bson:"updated_by"`
}

func (t *Team) ToEntity() (*entities.Team, error) {
	var team entities.Team
	err := copier.Copy(&team, t)
	return &team, err
}

type Teams []Team

func (tList Teams) ToEntity() []entities.Team {
	var teams []entities.Team
	for _, v := range tList {
		team, _ := v.ToEntity()
		teams = append(teams, *team)
	}
	return teams
}

type TeamSearch struct {
	ID       primitive.ObjectID `bson:"_id"`
	UUID     string             `bson:"uuid"`
	Code     string             `bson:"code"`
	Name     string             `bson:"name"`
	TeamType string             `bson:"team_type"`
	Academy  string             `bson:"academy"`
	Tel      string             `bson:"tel"`
	Contact  string             `bson:"contact"`
}

func (t *TeamSearch) ToEntityTeamSearch() (*entities.TeamSearch, error) {
	var TeamSearch entities.TeamSearch
	err := copier.Copy(&TeamSearch, t)
	return &TeamSearch, err
}

type TeamsSearch []TeamSearch

func (ts TeamsSearch) ToEntityTeamSearch() []entities.TeamSearch {
	var teamsSearch []entities.TeamSearch
	for _, v := range ts {
		teamssSearch, _ := v.ToEntityTeamSearch()
		teamsSearch = append(teamsSearch, *teamssSearch)
	}
	return teamsSearch
}
