package dtos

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
	"gitlab.com/chaihanij/evat/app/utils"
)

type FindOneFileRequest struct {
	UUID string `uri:"file_uuid"`
}

func (req *FindOneFileRequest) Parse(c *gin.Context) (*FindOneFileRequest, error) {
	if err := c.ShouldBindUri(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	return req, nil
}

func (req *FindOneFileRequest) ToEntity() *entities.FileFilter {
	return &entities.FileFilter{
		UUID: &req.UUID,
	}
}

type FindOneFileResponse struct{}

func (req *FindOneFileResponse) Response(c *gin.Context, file *entities.File) {
	fileBytes, err := ioutil.ReadFile(file.FileFullPath)
	if err != nil {
		utils.JSONErrorResponse(c, err)
		return
	}
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Header().Set("Content-Disposition", `attachment; filename=`+url.QueryEscape(file.FileName))
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Length", fmt.Sprintf("%d", len(fileBytes)))
	c.Writer.Write(fileBytes)
}
