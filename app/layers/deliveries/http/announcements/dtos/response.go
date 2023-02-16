package dtos

import (
	"time"
)

type AnnouncementResponse struct {
	UUID        string    `json:"uuid"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Year        string    `json:"year"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	CreatedBy   string    `json:"createdBy"`
	UpdatedBy   string    `json:"updatedBy"`
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
