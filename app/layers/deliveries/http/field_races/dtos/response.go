package dtos

import "time"

type FieldRaces struct {
	UUID        string `json:"uuid"`
	No          int    `json:"no"`
	Title       string `json:"title"`
	Description string `json:"description"`
	// Image       interface{}
	// Document    interface{}
	FullScore float64   `json:"fullscore"`
	IsActive  bool      `json:"isactive"`
	Year      string    `json:"year"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedBy string    `json:"updated_by"`
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

