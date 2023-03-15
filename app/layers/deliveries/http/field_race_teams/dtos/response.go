package dtos

import "time"

type FieldRaceTeam struct {
	FieldRaceUUID string    `json:"field_race_uuid"`
	TeamUUID      string    `json:"team_uuid"`
	Description   string    `json:"description"`
	Score         float64   `json:"score"`
	CreatedAt     time.Time `json:"create_at"`
	UpdatedAt     time.Time `json:"update_at"`
	CreatedBy     string    `json:"create_by"`
	UpdatedBy     string    `json:"updat_by"`
}

type FieldRaceTeamResponse struct {
	FieldRaceUUID string    `json:"field_race_uuid"`
	TeamUUID      string    `json:"team_uuid"`
	Description   string    `json:"description"`
	Score         float64   `json:"score"`
	CreatedAt     time.Time `json:"create_at"`
	UpdatedAt     time.Time `json:"update_at"`
	CreatedBy     string    `json:"create_by"`
	UpdatedBy     string    `json:"updat_by"`
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
