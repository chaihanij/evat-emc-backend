package dtos

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/errors"
)

type FileRequestJSON struct {
	AssignmentUUID   string                `uri:"assignment_uuid"`
	Document         *multipart.FileHeader `form:"file"`
	OriginalFileName string                `swaggerignore:"true"`
	FileName         string                `swaggerignore:"true"`
	FileExtension    string                `swaggerignore:"true"`
	FileFullPath     string                `swaggerignore:"true"`
	FilePath         string                `swaggerignore:"true"`
	UpdatedBy        string                `swaggerignore:"true"`
}

func (req *FileRequestJSON) Pares(c *gin.Context) (*FileRequestJSON, error) {

	if err := c.ShouldBindUri(req); err != nil {
		return nil, errors.InternalError{Message: err.Error()}
	}

	filescore, err := c.FormFile("file")
	if err != nil {
		return nil, errors.InternalError{Message: err.Error()}
	}

	req.Document = filescore

	fileExt := filepath.Ext(filescore.Filename)

	originalFileName := strings.TrimSuffix(filepath.Base(filescore.Filename), filepath.Ext(filescore.Filename))
	filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", time.Now().Unix()) + fileExt
	req.FileName = filename
	dir := filepath.Join("assignments", "score", filename)
	os.MkdirAll(filepath.Join(env.DataPath, dir), os.ModePerm)
	dst := filepath.Join(env.DataPath, dir, filename)
	if err := c.SaveUploadedFile(filescore, dst); err != nil {
		log.WithError(err).Debugln("UpdateMemberImageRequest Parse Error")
		return nil, errors.InternalError{Message: err.Error()}
	}

	req.FileExtension = fileExt
	req.FileFullPath = dst
	req.FilePath = filepath.Join(dir, filename)
	jwtRawData, ok := c.Get(constants.JWTDataKey)
	if !ok {
		return nil, errors.InternalError{Message: constants.JWTRestoreFail}
	}

	jwtData, ok := jwtRawData.(entities.JwtData)
	if !ok {
		return nil, errors.InternalError{Message: constants.JWTInvalidStructure}
	}

	if jwtData.UID == "" {
		return nil, errors.ParameterError{Message: constants.UserUIDMissing}
	}
	req.UpdatedBy = jwtData.UID

	return req, nil

}

func (req *FileRequestJSON) ToEntity() *entities.File {
	log.WithField("value", req).Debugln("UpdateAssignmentDocumentRequest ToEntity")
	return &entities.File{
		OriginalFileName: req.OriginalFileName,
		FileName:         req.FileName,
		FileExtension:    req.FileExtension,
		FileFullPath:     req.FileFullPath,
		FilePath:         req.FilePath,
	}
}

// func (req *FileRequestJSON) ToEntity() (*FileRequestJSON, error) {
// 	return &FileRequestJSON{
// 		AssignmentUUID: req.AssignmentUUID,
// 		FileName:       req.FileName,
// 	}, nil
// }
