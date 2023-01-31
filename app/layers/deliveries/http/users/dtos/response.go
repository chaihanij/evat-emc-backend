package dtos

import "time"

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

type User struct {
	UID       string     `json:"uid"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	FirstName string     `json:"firstname"`
	LastName  string     `json:"lastname"`
	Address   string     `json:"address"`
	Role      string     `json:"role"`
	Year      string     `json:"year"`
	IsActive  bool       `json:"isActive"`
	LastLogin *time.Time `json:"lastLogin"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	CreatedBy string     `json:"createdBy"`
	UpdatedBy string     `json:"updatedBy"`
}

type UserResponse struct {
	UID           string     `json:"uid"`
	Username      string     `json:"username"`
	Email         string     `json:"email"`
	FirstName     string     `json:"firstname"`
	LastName      string     `json:"lastname"`
	Address       string     `json:"address"`
	Tel           string     `json:"tel"`
	Role          string     `json:"role"`
	Password      string     `json:"password"`
	Year          string     `json:"year"`
	TeamUUID      string     `json:"teamUUID"`
	IsEmailVerify bool       `json:"isEmailVerify"`
	ActivateCode  string     `json:"activateCode"`
	AccessToken   string     `json:"accessToken"`
	IsActive      bool       `json:"isActive"`
	LastLogin     *time.Time `json:"lastLogin"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt"`
	CreatedBy     string     `json:"createdBy"`
	UpdatedBy     string     `json:"updatedBy"`
}
