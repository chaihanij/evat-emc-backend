package dtos

import (
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

type FindOneConsiderationRequestJSON struct {
	UUID string `uri:"consideration_uuid"`
}

func (req *FindOneConsiderationRequestJSON) Parse(c *gin.Context) (*FindOneConsiderationRequestJSON, error) {
	if err := c.ShouldBindUri(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	return req, nil
}

func (req *FindOneConsiderationRequestJSON) ToEntity() *entities.ConsiderationFilter {
	return &entities.ConsiderationFilter{
		TeamUUID: &req.UUID,
	}
}

type FindOneConsiderationResponseJSON ConsiderationResponse

func (m *FindOneConsiderationResponseJSON) Parse(c *gin.Context, input *entities.Consideration) *FindOneConsiderationResponseJSON {
	consideration := &FindOneConsiderationResponseJSON{
		// TeamUUID:    input.TeamUUID,
		// Description: input.Description,
		// CreatedAt:   input.CreatedAt,
		ID:        input.ID,
		Score:     input.Score,
		UpdatedAt: time.Now(),
		// CreatedBy:   input.CreatedBy,
		// UpdatedBy:   input.UpdatedBy,
	}

	return consideration
}
