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
	Topic            string `json:"topic"`
}

func (file *FileResponse) Parse(c *gin.Context, data *entities.File) *FileResponse {
	copier.Copy(file, data)
	file.URL = fmt.Sprintf("%s/v1/files/%s", env.BaseUrl, data.UUID)
	return file
}

type FilesResponse []FileResponse

type AssignmentResponse struct {
	AssignmentUUID string         `json:"assignmentUUID"`
	TeamUUID       string         `json:"teamUUID"`
	Description    string         `json:"description"`
	Documents      *FilesResponse `json:"documents"`
	IsConfirmed    bool           `json:"isConfirmed"`
	Score          float64        `json:"score"`
	Document       []Doc          `json:"document"`
}

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
	MemberType   string         `json:"memberType"`
	Documents    *FilesResponse `json:"documents"`
	IsTeamLeader bool           `json:"isTeamLeader"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	CreatedBy    string         `json:"createdBy"`
	UpdatedBy    string         `json:"updatedBy"`
	BirthDay     time.Time      `json:"birth_day"`
	NationalId   string         `json:"national_id"`
	Prefix       string         `json:"prefix"`
	Is_checkin   bool           `json:"is_checkin"`
	Is_national  bool           `json:"is_national"`
	Is_data      bool           `json:"is_data"`
	Is_image     bool           `json:"is_image"`
	Checkin_date time.Time      `json:"checkin_date"`
}

type MembersResponse []MemberResponse

type TeamResponse struct {
	UUID          string           `json:"uuid"`
	Code          string           `json:"code"`
	Name          string           `json:"name"`
	TeamType      string           `json:"teamType"`
	Academy       string           `json:"academy"`
	Major         string           `json:"major"`
	Detail        string           `json:"detail"`
	Year          string           `json:"year"`
	Slip          *FileResponse    `json:"slip,omitempty"`
	IsPaid        bool             `json:"is_paid"`
	PaymentMethod string           `json:"payment_method,omitempty"`
	IsVerify      bool             `json:"is_verify,omitempty"`
	Members       *MembersResponse `json:"members,omitempty"`
	CreatedAt     time.Time        `json:"createdAt"`
	UpdatedAt     time.Time        `json:"updatedAt"`
	CreatedBy     string           `json:"createdBy"`
	UpdatedBy     string           `json:"updatedBy"`
	PaidDateTime  time.Time        `json:"paid_date_time"`
}

type TeamSearchResponse struct {
	UUID     string `json:"uuid"`
	Code     string `json:"code"`
	Name     string `json:"name"`
	TeamType string `json:"team_type"`
	Academy  string `json:"academy"`
	Tel      string `json:"tel"`
	Contact  string `json:"contact"`
}

type MetaDataResponse struct {
	TotalRecords uint `json:"totalRecords" example:"10"`
	Page         uint `json:"page" example:"1"`
	PageSize     uint `json:"pageSize" example:"20"`
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
