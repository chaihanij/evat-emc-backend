package dtos

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ScoreResponse struct {
	// ID             primitive.ObjectID `json:"_id" bson:"_id"`
	// UID            string    `json:"uid" bson:"uid"`
	// NameTeam       string    `json:"nameteam" bson:"nameteam"`
	// FirstTeam      string    `json:"firstteam" bson:"firstteam"`
	// SecondTeam     string    `json:"secondteam" bson:"secondteam"`
	// First_Stadium  string    `json:"firststadium" bson:"firststadium"`
	// Second_Stadium string    `json:"secondstadium" bson:"secondstadium"`
	// Third_Stadium  string    `json:"thirdstadium" bson:"thirdstadium"`
	// Fourth_Stadium string    `json:"fourthstadium" bson:"fourthstadium"`
	// Sum_Score      string    `json:"sumscore" bson:"sumscore"`
	// No             string    `json:"no" bson:"no"`
	// CreateDate     time.Time `json:"createdate" bson:"createdate"`
	// Lastupdate     time.Time `json:"lastupdate" bson:"lastupdate"`
	// CreateBy       string    `json:"createby" bson:"createby"`
	// LastUpdateBy   string    `json:"lastupdateby" bson:"lastupdateby"`
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	UUID            string             `json:"uuid" bson:"uuid"`
	Field_race_uuid string             `json:"field_race_uuid" bson:"field_race_uuid"`
	Team_uuid       string             `json:"team_uuid" bson:"team_uuid"`
	Score           float64            `json:"score" bson:"score"`
	CreateDate      time.Time          `json:"create_at" bson:"create_at"`
	Lastupdate      time.Time          `json:"update_at" bson:"update_at"`
	CreateBy        string             `json:"create_by" bson:"create_by"`
	UpdateBy        string             `json:"update_by" bson:"update_by"`
}

type MetaDataResponse struct {
	TotalRecords uint `json:"totalRecords" example:"10"`
	Page         uint `json:"page" example:"1"`
	PageSize     uint `json:"pageSize" example:"20"`
}

func (m *MetaDataResponse) Parse(page *int64, pageSize *int64, totalRecords *int64) *MetaDataResponse {
	m.TotalRecords = uint(*totalRecords)

	if page != nil {
		m.Page = uint(*page)
	}
	if pageSize != nil {
		m.PageSize = uint(*pageSize)
	}

	return m
}
