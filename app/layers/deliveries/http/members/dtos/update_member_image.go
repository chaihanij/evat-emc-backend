package dtos

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/errors"
)

type UpdateMemberImageRequest struct {
	MemberUUID       string                `json:"-" uri:"member_uuid" binding:"required,uuid"`
	Image            *multipart.FileHeader `json:"-" form:"image" binding:"required"`
	OriginalFileName string                `json:"-" swaggerignore:"true"`
	FileName         string                `json:"-" swaggerignore:"true"`
	FileExtension    string                `json:"-" swaggerignore:"true"`
	FilePath         string                `json:"-" swaggerignore:"true"`
}

func (req *UpdateMemberImageRequest) Parse(c *gin.Context) (*UpdateMemberImageRequest, error) {
	if err := c.ShouldBind(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	log.WithField("value", req).Debugln("UpdateMemberImageRequest After ShouldBind Request")

	fileExt := filepath.Ext(req.Image.Filename)
	originalFileName := strings.TrimSuffix(filepath.Base(req.Image.Filename), filepath.Ext(req.Image.Filename))
	now := time.Now()
	filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", now.Unix()) + fileExt
	dst := filepath.Join(env.DataPath, "members", "image", filename)
	if err := c.SaveUploadedFile(req.Image, dst); err != nil {
		log.WithError(err).Debugln("UpdateMemberImageRequest Parse Error")
		return nil, errors.InternalError{Message: err.Error()}
	}
	req.OriginalFileName = originalFileName
	req.FileName = filename
	req.FileExtension = fileExt
	req.FilePath = dst
	log.WithField("value", req).Debugln("UpdateMemberImageRequest Parse")
	return req, nil
}

func (req *UpdateMemberImageRequest) ToEntity() *entities.File {
	return &entities.File{
		OriginalFileName: req.OriginalFileName,
		FileName:         req.FileName,
		FileExtension:    req.FileExtension,
		FilePath:         req.FilePath,
	}
}

type UpdateMemberImageResponseSwagger struct {
	StatusCode    int          `json:"statusCode" example:"1000"`
	StatusMessage string       `json:"statusMessage" example:"Success"`
	Timestamp     time.Time    `json:"timestamp" example:"2015-06-30T21:59:59Z"`
	Data          FileResponse `json:"data,omitempty"`
}
