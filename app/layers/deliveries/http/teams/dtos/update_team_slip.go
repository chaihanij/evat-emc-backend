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
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/env"
	"gitlab.com/chaihanij/evat/app/errors"
)

type UpdateTeamSlipRequest struct {
	TeamUUID         string                `uri:"team_uuid" binding:"required,uuid"`
	Slip             *multipart.FileHeader `form:"slip"`
	OriginalFileName string                `swaggerignore:"true"`
	FileName         string                `swaggerignore:"true"`
	FileExtension    string                `swaggerignore:"true"`
	FileFullPath     string                `swaggerignore:"true"`
	FilePath         string                `swaggerignore:"true"`
}

func (req *UpdateTeamSlipRequest) Parse(c *gin.Context) (*UpdateTeamSlipRequest, error) {
	if err := c.ShouldBindUri(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	slip, err := c.FormFile("slip")
	if err != nil {
		return nil, errors.InternalError{Message: err.Error()}
	}
	req.Slip = slip
	fileExt := filepath.Ext(slip.Filename)
	originalFileName := strings.TrimSuffix(filepath.Base(slip.Filename), filepath.Ext(slip.Filename))
	filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", time.Now().Unix()) + fileExt
	dir := filepath.Join("teams", "slip", req.TeamUUID)
	os.MkdirAll(filepath.Join(env.DataPath, dir), os.ModePerm)
	dst := filepath.Join(env.DataPath, dir, filename)
	if err := c.SaveUploadedFile(slip, dst); err != nil {
		log.WithError(err).Debugln("UpdateTeamSlipRequest Parse Error")
		return nil, errors.InternalError{Message: err.Error()}
	}
	req.OriginalFileName = originalFileName
	req.FileName = filename
	req.FileExtension = fileExt
	req.FilePath = filepath.Join(dir, filename)
	req.FileFullPath = dst
	log.WithField("value", req).Debugln("UpdateTeamSlipRequest")
	return req, nil
}

func (req *UpdateTeamSlipRequest) ToEntity() (string, *entities.File) {
	return req.TeamUUID,
		&entities.File{
			OriginalFileName: req.OriginalFileName,
			FileName:         req.FileName,
			FileExtension:    req.FileExtension,
			FileFullPath:     req.FileFullPath,
			FilePath:         req.FilePath,
		}
}
