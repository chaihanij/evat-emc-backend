package dtos

import (
	"time"

	"github.com/AlekSi/pointer"
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

type SendAssignmentTeamPullDocumentRequestJSON struct {
	TeamUUID       string `uri:"team_uuid"`
	AssignmentUUID string `uri:"assignment_uuid"`
	DocumentUUID   string `uri:"document_uuid" `
}

func (req *SendAssignmentTeamPullDocumentRequestJSON) Parse(c *gin.Context) (*SendAssignmentTeamPullDocumentRequestJSON, error) {
	if err := c.ShouldBindUri(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	return req, nil
}

func (req *SendAssignmentTeamPullDocumentRequestJSON) ToEntity() (*entities.AssignmentTeamPartialUpdate, string) {
	return &entities.AssignmentTeamPartialUpdate{
			TeamUUID:       pointer.ToString(req.TeamUUID),
			AssignmentUUID: pointer.ToString(req.AssignmentUUID),
		},
		req.DocumentUUID
}

type SendAssignmentTeamPullDocumentJSONJSwagger struct {
	StatusCode    int       `json:"statusCode" example:"1000"`
	StatusMessage string    `json:"statusMessage" example:"Success"`
	Timestamp     time.Time `json:"timestamp" example:"2015-06-30T21:59:59Z"`
}
