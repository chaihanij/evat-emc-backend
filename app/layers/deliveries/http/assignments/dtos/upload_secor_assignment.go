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

type UploadScoreAssingment struct {
	AssignmentUUID   string                `uri:"assignment_uuid"`
	Document         *multipart.FileHeader `swaggerignore:"true" form:"document"`
	OriginalFileName string                `swaggerignore:"true"`
	FileName         string                `swaggerignore:"true"`
	FileExtension    string                `swaggerignore:"true"`
	FileFullPath     string                `swaggerignore:"true"`
	FilePath         string                `swaggerignore:"true"`
	UpdatedBy        string                `swaggerignore:"true"`
}

func (req *UploadScoreAssingment) Parse(c *gin.Context) (*UploadScoreAssingment, error) {
	if err := c.ShouldBindUri(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	document, err := c.FormFile("document")
	if err != nil {
		fmt.Println("err:", err)
		return nil, errors.InternalError{Message: err.Error()}
	}
	// fmt.Println("abc :", document)

	req.Document = document
	fileExt := filepath.Ext(document.Filename)
	originalFileName := strings.TrimSuffix(filepath.Base(document.Filename), filepath.Ext(document.Filename))
	filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", time.Now().Unix()) + fileExt
	dir := filepath.Join("assignment_teams", "documents")
	os.MkdirAll(filepath.Join(env.DataPath, dir), os.ModePerm)
	dst := filepath.Join(env.DataPath, dir, filename)
	if err := c.SaveUploadedFile(document, dst); err != nil {
		log.WithError(err).Debugln("UpdateMemberImageRequest Parse Error")
		return nil, errors.InternalError{Message: err.Error()}
	}
	fmt.Println("fileExt :", filepath.Ext(document.Filename))
	fmt.Println("dtr", dir)
	fmt.Println("dst", dst)
	fmt.Println("document :", document.Size)
	fmt.Println("originalFileName :", originalFileName)

	req.OriginalFileName = originalFileName
	req.FileName = filename
	req.FileExtension = fileExt
	//req.FilePath = filepath.Join(dir, filename)
	//req.FileFullPath = dst
	// log.WithField("value", req).Debugln("SendAssignmentDocumentRequestJSON")

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
