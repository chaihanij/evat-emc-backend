package dtos

import (
	"time"

	"github.com/gin-gonic/gin"
	// "github.com/google/martian/log"
	// log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
	"gitlab.com/chaihanij/evat/app/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateFieldRaceTeamRequestJSON struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id"`
	FieldRaceUUID string             `json:"field_race_uuid" bson:"field_race_uuid"`
	TeamUUID      string             `json:"team_uuid" bson:"team_uuid"`
	Score         float64            `score:"score" bson:"score"`
	CreatedAt     time.Time          `hson:"created_at" bson:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at" bson:"updated_at"`
	CreatedBy     string             `json:"created_by" bson:"created_by"`
	UpdatedBy     string             `json:"updated_by" bson:"updated_by"`

}

type CreateFieldRaceTeamResponsJSON struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id"`
	FieldRaceUUID string             `json:"field_race_uuid" bson:"field_race_uuid"`
	TeamUUID      string             `json:"team_uuid" bson:"team_uuid"`
	Score         float64            `score:"score" bson:"score"`
	CreatedAt     time.Time          `hson:"created_at" bson:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at" bson:"updated_at"`
	// CreatedBy     string             `json:"created_by" bson:"created_by"`
	// UpdatedBy     string             `json:"updated_by" bson:"updated_by"`
	// Name          string             `json:"name" bson:"name"`
	// Code          string             `json:"code" bson:"code" `
	// Type          string             `json:"type" bson:"type" `
}

func (req *CreateFieldRaceTeamRequestJSON) Parse(c *gin.Context) (*CreateFieldRaceTeamRequestJSON, error) {
	err := c.ShouldBindJSON(req)
	if err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	err = types.Validate.Struct(req)
	if err != nil {
		if errValidate := types.HandleValidateError(err, req); errValidate != nil {
			return nil, errors.ParameterError{Message: errValidate.Error()}
		}
		return nil, errors.ParameterError{Message: err.Error()}
	}

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

	req.CreatedBy = jwtData.UID

	return req, nil
}

func (req *CreateFieldRaceTeamRequestJSON) ToEntity() *entities.FieldRaceTeam {
	return &entities.FieldRaceTeam{
		FieldRaceUUID: req.FieldRaceUUID,
		TeamUUID:      req.TeamUUID,
		Score:         req.Score,
		CreatedAt:     req.CreatedAt,
		UpdatedAt:     req.UpdatedAt,
		CreatedBy:     req.CreatedBy,
		UpdatedBy:     req.UpdatedBy,
	}
}

func (m *CreateFieldRaceTeamResponsJSON) Parse(input *entities.FieldRaceTeam) *CreateFieldRaceTeamResponsJSON {
	return &CreateFieldRaceTeamResponsJSON{
		ID:            primitive.NewObjectID(),
		FieldRaceUUID: m.FieldRaceUUID,
		TeamUUID:      m.TeamUUID,
		Score:         m.Score,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		// CreatedBy:     m.CreatedBy,
		// UpdatedBy:     m.UpdatedBy,
	}
}
