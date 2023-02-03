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

type UpdateAssignmentImageRequest struct {
	AssignmentUUID   string                `uri:"assignment_uuid" binding:"required,uuid"`
	Image            *multipart.FileHeader `swaggerignore:"true" form:"image"`
	OriginalFileName string                `swaggerignore:"true"`
	FileName         string                `swaggerignore:"true"`
	FileExtension    string                `swaggerignore:"true"`
	FileFullPath     string                `swaggerignore:"true"`
	FilePath         string                `swaggerignore:"true"`
}

func (req *UpdateAssignmentImageRequest) Parse(c *gin.Context) (*UpdateAssignmentImageRequest, error) {
	if err := c.ShouldBindUri(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	image, err := c.FormFile("image")
	if err != nil {
		return nil, errors.InternalError{Message: err.Error()}
	}
	req.Image = image

	fileExt := filepath.Ext(req.Image.Filename)
	originalFileName := strings.TrimSuffix(filepath.Base(req.Image.Filename), filepath.Ext(req.Image.Filename))
	filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", time.Now().Unix()) + fileExt
	dst := filepath.Join(env.DataPath, "assignments", "images", filename)
	if err := c.SaveUploadedFile(req.Image, dst); err != nil {
		log.WithError(err).Debugln("UpdateAssignmentImageRequest Parse Error")
		return nil, errors.InternalError{Message: err.Error()}
	}
	req.OriginalFileName = originalFileName
	req.FileName = filename
	req.FileExtension = fileExt
	req.FilePath = filepath.Join("assignments", "images", filename)
	req.FileFullPath = dst
	log.WithField("value", req).Debugln("UpdateAssignmentImageRequest")
	return req, nil
}

func (req *UpdateAssignmentImageRequest) ToEntity() *entities.File {
	log.WithField("value", req).Debugln("UpdateAssignmentImageRequest ToEntity")
	return &entities.File{
		OriginalFileName: req.OriginalFileName,
		FileName:         req.FileName,
		FileExtension:    req.FileExtension,
		FileFullPath:     req.FileFullPath,
		FilePath:         req.FilePath,
	}
}

type UpdateAssignmentImageResponseSwagger struct {
	StatusCode    int          `json:"statusCode" example:"1000"`
	StatusMessage string       `json:"statusMessage" example:"Success"`
	Timestamp     time.Time    `json:"timestamp" example:"2015-06-30T21:59:59Z"`
	Data          FileResponse `json:"data,omitempty"`
}
