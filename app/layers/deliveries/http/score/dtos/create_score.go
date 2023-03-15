package dtos

import (
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
	"gitlab.com/chaihanij/evat/app/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateScoreRequestJSON struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	UUID            string             `json:"uuid" bson:"uuid"`
	Field_race_uuid string             `json:"field_race_uuid" bson:"field_race_uuid"`
	Team_uuid       string             `json:"team_uuid" bson:"team_uuid"`
	Score           float64            `json:"score" bson:"score"`
	CreateDate      time.Time          `json:"create_at" bson:"create_at"`
	Lastupdate      time.Time          `json:"update_at" bson:"update_at"`
	CreateBy        string             `json:"create_by" bson:"create_by"`
	UpdateBy        string             `json:"update_by" bson:"update_by"`
}

type CreateScoreResponseJSON struct {
	// ID             primitive.ObjectID `json:"_id" bson:"_id"`
	UUID            string    `json:"uuid" bson:"uuid"`
	Field_race_uuid string    `json:"field_race_uuid" bson:"field_race_uuid"`
	Team_uuid       string    `json:"team_uuid" bson:"team_uuid"`
	Score           float64   `json:"score" bson:"score"`
	CreateDate      time.Time `json:"create_at" bson:"create_at"`
	Lastupdate      time.Time `json:"update_at" bson:"update_at"`
	CreateBy        string    `json:"create_by" bson:"create_by"`
	UpdateBy        string    `json:"update_by" bson:"update_by"`
}

func (req *CreateScoreRequestJSON) Parse(c *gin.Context) (*CreateScoreRequestJSON, error) {

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
	req.CreateBy = jwtData.UID

	return req, nil

}

func (req *CreateScoreRequestJSON) ToEntity() *entities.Score {

	return &entities.Score{
		UUID: req.UUID,
		Field_race_uuid: req.Field_race_uuid,
		Team_uuid: req.Team_uuid,
		Score: req.Score,
		CreateDate: time.Now(),
		CreateBy:   req.CreateBy,
		Lastupdate: time.Now(),
		Updated_by: req.UpdateBy,
	}

}

func (m *CreateScoreResponseJSON) Parse(input *entities.Score) *CreateScoreResponseJSON {

	return &CreateScoreResponseJSON{

		// ID: string(primitive.NewObjectID()),
		// UID:            input.UUID,
		UUID: input.UUID,
		Field_race_uuid: input.Field_race_uuid,
		Team_uuid: input.Team_uuid,
		Score: input.Score,
		CreateDate: time.Now(),
		CreateBy:   input.CreateBy,
		Lastupdate: time.Now(),
		UpdateBy: input.Updated_by,
	}

}

type CreateScoreResponseSwagger struct {
	StatusCode    int                    `json:"statusCode" example:"1000"`
	StatusMessage string                 `json:"statusMessage" example:"Success"`
	Timestamp     time.Time              `json:"timestamp" example:"2015-06-30T21:59:59Z"`
	Data          CreateScoreRequestJSON `json:"data,omitempty"`
}
