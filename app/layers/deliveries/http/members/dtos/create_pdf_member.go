package dtos

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/errors"
)

type UpdatePDFRequest struct {
	MemberUUID string                `uri:"member_uuid" binding:"required,uuid"`
	Document   *multipart.FileHeader `swaggerignore:"true" form:"certificate"`

	OriginalFileName string `swaggerignore:"true"`
	FileName         string `swaggerignore:"true"`
	FileExtension    string `swaggerignore:"true"`
	FileFullPath     string `swaggerignore:"true"`
	FilePath         string `swaggerignore:"true"`
}

func (req *UpdatePDFRequest) Parse(c *gin.Context) (*UpdatePDFRequest, error) {
	if err := c.ShouldBindUri(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	document, err := c.FormFile("certificate")
	if err != nil {
		return nil, errors.InternalError{Message: err.Error()}
	}
	req.Document = document
	fileExt := filepath.Ext(req.Document.Filename)
	originalFileName := strings.TrimSuffix(filepath.Base(req.Document.Filename), filepath.Ext(req.Document.Filename))
	now := time.Now()
	filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", now.Unix()) + fileExt
	dst := filepath.Join(env.DataPath, "members", "certificate", filename)
	if err := c.SaveUploadedFile(req.Document, dst); err != nil {
		// log.WithError(err).Debugln("UpdateMemberPushDocumentRequest Parse Error")
		return nil, errors.InternalError{Message: err.Error()}
	}
	req.OriginalFileName = originalFileName
	req.FileName = filename
	req.FileExtension = fileExt
	req.FileFullPath = dst
	req.FilePath = filepath.Join("members", "certificate", filename)
	// log.WithField("value", req).Debugln("UpdateMemberPushDocumentRequest Parse")
	return req, nil
}

// type UpdatePDFRespones UpdateMemberPushDocumentRequest

func (req *UpdatePDFRequest) ToEntity() *entities.File {
	return &entities.File{
		OriginalFileName: req.OriginalFileName,
		FileName:         req.FileName,
		FileExtension:    req.FileExtension,
		FileFullPath:     req.FileFullPath,
		FilePath:         req.FilePath,
	}
}
