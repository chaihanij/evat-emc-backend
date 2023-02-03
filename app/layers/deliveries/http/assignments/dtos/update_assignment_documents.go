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

type UpdateAssignmentDocumentRequest struct {
	AssignmentUUID   string                `uri:"assignment_uuid" binding:"required,uuid"`
	Document         *multipart.FileHeader `swaggerignore:"true" form:"document"`
	OriginalFileName string                `swaggerignore:"true"`
	FileName         string                `swaggerignore:"true"`
	FileExtension    string                `swaggerignore:"true"`
	FileFullPath     string                `swaggerignore:"true"`
	FilePath         string                `swaggerignore:"true"`
}

func (req *UpdateAssignmentDocumentRequest) Parse(c *gin.Context) (*UpdateAssignmentDocumentRequest, error) {
	if err := c.ShouldBindUri(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	document, err := c.FormFile("document")
	if err != nil {
		return nil, errors.InternalError{Message: err.Error()}
	}
	req.Document = document

	fileExt := filepath.Ext(req.Document.Filename)
	originalFileName := strings.TrimSuffix(filepath.Base(req.Document.Filename), filepath.Ext(req.Document.Filename))
	filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", time.Now().Unix()) + fileExt
	dst := filepath.Join(env.DataPath, "assignments", "documents", filename)
	if err := c.SaveUploadedFile(req.Document, dst); err != nil {
		log.WithError(err).Debugln("UpdateAssignmentDocumentRequest Parse Error")
		return nil, errors.InternalError{Message: err.Error()}
	}
	req.OriginalFileName = originalFileName
	req.FileName = filename
	req.FileExtension = fileExt
	req.FilePath = filepath.Join("assignments", "documents", filename)
	req.FileFullPath = dst
	log.WithField("value", req).Debugln("UpdateAssignmentDocumentRequest")
	return req, nil
}

func (req *UpdateAssignmentDocumentRequest) ToEntity() *entities.File {
	log.WithField("value", req).Debugln("UpdateAssignmentDocumentRequest ToEntity")
	return &entities.File{
		OriginalFileName: req.OriginalFileName,
		FileName:         req.FileName,
		FileExtension:    req.FileExtension,
		FileFullPath:     req.FileFullPath,
		FilePath:         req.FilePath,
	}
}

type UpdateAssignmentDocumentResponseSwagger struct {
	StatusCode    int          `json:"statusCode" example:"1000"`
	StatusMessage string       `json:"statusMessage" example:"Success"`
	Timestamp     time.Time    `json:"timestamp" example:"2015-06-30T21:59:59Z"`
	Data          FileResponse `json:"data,omitempty"`
}
