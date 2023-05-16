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

type AssignmentResponse struct {
	UUID        string        `json:"uuid"`
	TeamUUID    string        `json:"team_uuid"`
	No          int           `json:"no"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Image       *FileResponse `json:"image,omitempty"`
	Document    *FileResponse `json:"document,omitempty"`
	FullScore   float64       `json:"fullScore"`
	IsActive    bool          `json:"isActive"`
	DueDate     time.Time     `json:"dueDate"`
	Year        string        `json:"year"`
	CreatedAt   time.Time     `json:"createdAt"`
	UpdatedAt   time.Time     `json:"updatedAt"`
	CreatedBy   string        `json:"createdBy"`
	UpdatedBy   string        `json:"updatedBy"`
	SendDoc     bool          `json:"senddoc"`
	//OverDue      time.Time     `json:"overdue"`
	DeliveryTime  time.Time                 `json:"delivery_time"`
	Consideration []ConsiderationAssignment `json:"consideration"`
	IsShowMenu    bool                      `json:"isShowMenu"`
	File          []File                    `json:"files" bson:"files"`
	Topic         []string                  `json:"topic"`
}

type File struct {
	// FileUrl    string    `json:"fileurl"  bson:"fileurl"`
	// FileName   string    `json:"filename" bson:"filename"`
	// CreateDate time.Time `json:"createdate" bson:"createdate"`
	// Createby   string    `json:"createby" bson:"createby"`
	Fileurl    string    `json:"fileurl" bson:"fileurl"`
	FileName   string    `json:"filename" bson:"filename"`
	Createby   string    `json:"createby" bson:"createby"`
	CreateDate time.Time `json:"createdate" bson:"createdate"`
}

type TeamAssignmentResponse struct {
	UUID      string  `json:"uuid"`
	TeamUUID  string  `json:"team_uuid"`
	Title     string  `json:"title"`
	FullScore float64 `json:"fullScore"`
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
