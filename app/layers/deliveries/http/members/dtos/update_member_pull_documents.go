package dtos

import (
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/errors"
)

type UpdateMemberPullDocumentRequest struct {
	UUID         string `uri:"member_uuid" binding:"required,uuid"`
	DocumentUUID string `uri:"document_uuid" binding:"required,uuid"`
}

func (req *UpdateMemberPullDocumentRequest) Parse(c *gin.Context) (*UpdateMemberPullDocumentRequest, error) {
	if err := c.ShouldBindUri(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	return req, nil
}

func (req *UpdateMemberPullDocumentRequest) ToEntity() (string, string) {
	return req.UUID, req.DocumentUUID
}

type UpdateMemberPullDocumentResponseSwagger struct {
	StatusCode    int       `json:"statusCode" example:"1000"`
	StatusMessage string    `json:"statusMessage" example:"Success"`
	Timestamp     time.Time `json:"timestamp" example:"2015-06-30T21:59:59Z"`
}
