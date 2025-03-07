package dtos

import (
	"time"

	"github.com/AlekSi/pointer"
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

type FindOneAssignmentTeamRequestJSON struct {
	TeamUUID       string `json:"-" uri:"team_uuid"`
	AssignmentUUID string `json:"-" uri:"assignment_uuid"`
}

func (req *FindOneAssignmentTeamRequestJSON) Parse(c *gin.Context) (*FindOneAssignmentTeamRequestJSON, error) {
	if err := c.ShouldBindUri(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	return req, nil
}

func (req *FindOneAssignmentTeamRequestJSON) ToEntity() *entities.AssignmentTeamFilter {
	return &entities.AssignmentTeamFilter{
		TeamUUID:       pointer.ToString(req.TeamUUID),
		AssignmentUUID: pointer.ToString(req.AssignmentUUID),
	}
}

type FindOneAssignmentTeamResponseJSON AssignmentResponse

type Doc struct {
	DocumentUUID    string `json:"documentUUID" bson:"documentUUID"`
	AssignmentTopic string `json:"AssignmentTopic" bson:"AssignmentTopic"`
}

func (res *FindOneAssignmentTeamResponseJSON) Parse(c *gin.Context, input *entities.AssignmentTeam) *FindOneAssignmentTeamResponseJSON {
	var documents FilesResponse

	// fmt.Println("Documents :", input.Documents.([]entities.File))

	if val, ok := input.Documents.([]entities.File); ok {
		for i, value := range val {
			document := new(FileResponse).Parse(c, &value)
			document.Topic = input.Document[i].AssignmentTopic
			documents = append(documents, *document)
		}
	}
	var documentsG FilesResponse

	for _, va := range input.Document {
		documentsGs := entities.File{
			UUID:  va.DocumentUUID,
			Topic: va.AssignmentTopic,
		}
		document := new(FileResponse).Parse(c, &documentsGs)
		documentsG = append(documentsG, *document)

	}

	return &FindOneAssignmentTeamResponseJSON{
		TeamUUID:       input.TeamUUID,
		AssignmentUUID: input.AssignmentUUID,
		Description:    input.Description,
		IsConfirmed:    input.IsConfirmed,
		Score:          input.Score,
		Documents:      &documents,
		// Document:       input.Document,
	}
}

type FindOneAssignmentTeamResponseSwagger struct {
	StatusCode    int                               `json:"statusCode" example:"1000"`
	StatusMessage string                            `json:"statusMessage" example:"Success"`
	Timestamp     time.Time                         `json:"timestamp" example:"2015-06-30T21:59:59Z"`
	Data          FindOneAssignmentTeamResponseJSON `json:"data,omitempty"`
}
