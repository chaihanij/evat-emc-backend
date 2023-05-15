package dtos

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
)

type FileResponse struct {
	URL              string `json:"url"`
	UUID             string `json:"uuid"`
	OriginalFileName string `json:"originalFileName"`
	FileName         string `json:"fileName"`
}

func (file *FileResponse) Parse(c *gin.Context, data *entities.File) *FileResponse {
	copier.Copy(file, data)
	file.URL = fmt.Sprintf("%s/v1/files/%s", env.BaseUrl, data.UUID)
	return file
}

type FilesResponse []FileResponse

type MemberResponse struct {
	UUID         string         `json:"uuid"`
	FirstName    string         `json:"firstname"`
	LastName     string         `json:"lastname"`
	Address      string         `json:"address"`
	Email        string         `json:"email"`
	Tel          string         `json:"tel"`
	Academy      string         `json:"academy"`
	Major        string         `json:"major"`
	Image        *FileResponse  `json:"image"`
	Year         string         `json:"year"`
	TeamUUID     string         `json:"teamUUID"`
	MemberType   string         `json:"memberType"`
	Documents    *FilesResponse `json:"documents"`
	IsTeamLeader bool           `json:"isTeamLeader"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	CreatedBy    string         `json:"createdBy"`
	UpdatedBy    string         `json:"updatedBy"`
	BirthDay     time.Time      `json:"birth_day" `
	NationalId   string         `json:"national_id" `
	Checkin_date time.Time      `json:"checkin_date"`
	Is_checkin   bool           `json:"is_checkin"`
	Is_data      bool           `json:"is_data"`
	Is_image     bool           `json:"is_image"`
	Is_national  bool           `json:"is_national"`
	Prefix       string         `json:"prefix"`
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
