package dtos

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

type FindOneTeamRequestJSON struct {
	UUID string `uri:"team_uuid"`
}

func (req *FindOneTeamRequestJSON) Parse(c *gin.Context) (*FindOneTeamRequestJSON, error) {
	if err := c.ShouldBindUri(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	return req, nil
}

func (req *FindOneTeamRequestJSON) ToEntity() *entities.TeamFilter {
	return &entities.TeamFilter{
		UUID: &req.UUID,
	}
}

type FindOneTeamResponseJSON TeamResponse

func (m *FindOneTeamResponseJSON) Parse(input *entities.Team) *FindOneTeamResponseJSON {
	teams := &FindOneTeamResponseJSON{
		UUID:      input.UUID,
		Code:      input.Code,
		Name:      input.Name,
		TeamType:  input.TeamType,
		Academy:   input.Academy,
		Detail:    input.Detail,
		Year:      input.Year,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
		CreatedBy: input.CreatedBy,
		UpdatedBy: input.UpdatedBy,
	}
	if value, ok := input.Members.(entities.Members); ok {
		var members MembersResponse
		for _, m := range value {
			var member MemberResponse
			copier.Copy(member, m)
			if val, ok := m.Image.(entities.File); ok {
				var image FileResponse
				copier.Copy(image, val)
				member.Image = &image
			} else {
				member.Image = nil
			}
			if val, ok := m.Documents.(entities.Files); ok {
				var documents FilesResponse
				for _, value := range val {
					var document FileResponse
					copier.Copy(document, value)
					documents = append(documents, document)
				}
				member.Documents = &documents
			} else {
				member.Documents = &FilesResponse{}
			}
			members = append(members, member)
		}
		teams.Members = &members
	} else {
		teams.Members = &MembersResponse{}
	}
	return teams
}

type FindOneTeamResponseSwagger struct {
	StatusCode    int                     `json:"statusCode" example:"1000"`
	StatusMessage string                  `json:"statusMessage" example:"Success"`
	Timestamp     time.Time               `json:"timestamp" example:"2015-06-30T21:59:59Z"`
	Data          FindOneTeamResponseJSON `json:"data,omitempty"`
}
