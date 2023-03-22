package dtos

import "time"

type ConsiderationResponse struct {
	ID             string           `json:"_id"`
	TotalScore     float64          `json:"total_score"`
	UpdatedAt      time.Time        `json:"update_at"`
	No             int              `json:"no"`
	IndivdualScore []IndivdualScore `json:"indivdual_score"`
}

type IndivdualScore struct {
	Title  string  `json:"title"`
	Score float64 `json:"score"`
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
