package dtos

type ConsiderationResponse struct {
	// ID             string          `json:"_id"`
	Total          float64         `json:"total"`
	Considerations []Consideration `json:"considerations"`
}

// type Considerations struct {
// 	Consideration Consideration `json:"consideration"`
// }

type Consideration struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	TeamName string  `json:"nameteam"`
	Score    float64 `json:"score"`
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
