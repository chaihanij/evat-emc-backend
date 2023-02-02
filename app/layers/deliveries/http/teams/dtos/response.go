package dtos

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/entities"
)

type FileResponse struct {
	URL              string `json:"url"`
	UUID             string `json:"uuid"`
	OriginalFileName string `json:"originalFileName"`
	FileName         string `json:"fileName"`
	// FileExtension    string    `json:"fileExtension"`
	// FilePath         string    `json:"filePath"`
	// FileFullPath     string    `json:"fileFullPath"`
	// CreatedAt        time.Time `json:"createdAt"`
	// UpdatedAt        time.Time `json:"updatedAt"`
}

func (file *FileResponse) Parse(c *gin.Context, data *entities.File) *FileResponse {
	copier.Copy(file, data)
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	file.URL = fmt.Sprintf("%s://%s/v1/files/%s", scheme, c.Request.Host, data.UUID)
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
	Image        *FileResponse  `json:"image"`
	Year         string         `json:"year"`
	MemberType   string         `json:"memberType"`
	Documents    *FilesResponse `json:"documents"`
	IsTeamLeader bool           `json:"isTeamLeader"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	CreatedBy    string         `json:"createdBy"`
	UpdatedBy    string         `json:"updatedBy"`
}

type MembersResponse []MemberResponse

type TeamResponse struct {
	UUID      string           `json:"uuid"`
	Code      string           `json:"code"`
	Name      string           `json:"name"`
	TeamType  string           `json:"teamType"`
	Academy   string           `json:"academy"`
	Detail    string           `json:"detail"`
	Year      string           `json:"year"`
	Members   *MembersResponse `json:"members,omitempty"`
	CreatedAt time.Time        `json:"createdAt"`
	UpdatedAt time.Time        `json:"updatedAt"`
	CreatedBy string           `json:"createdBy"`
	UpdatedBy string           `json:"updatedBy"`
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
