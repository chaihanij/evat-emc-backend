package dtos

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

type UploadScoreFieldRace struct {
	FieldRaceUUID          string                   `uri:"uuid"`
	ConsiderationFieldRace []ConsiderationFieldRace `json:"consideration" bson:"consideration"`
}

type ConsiderationFieldRace struct {
	ID       string  `json:"id" bson:"id"`
	Title    string  `json:"title" bson:"title"`
	NameTeam string  `json:"nameteam" bson:"nameteam"`
	Score    float64 `json:"score" bson:"score"`
}

func (req *UploadScoreFieldRace) Parse(c *gin.Context) (*UploadScoreFieldRace, error) {

	if err := c.ShouldBindUri(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}

	if err := c.ShouldBindJSON(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	log.WithField("value", req).Debugln("SendAssignmentDocumentRequestJSON")

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
	return req, nil

}

func (req *UploadScoreFieldRace) ToEntity() *entities.FieldRace {

	var scoreFieldRaces []entities.ConsiderationFieldRace

	for _, value := range req.ConsiderationFieldRace {

		scoreFieldRace := &entities.ConsiderationFieldRace{
			ID:       value.ID,
			NameTeam: value.NameTeam,
			Score:    value.Score,
			Title:    value.Title,
		}

		scoreFieldRaces = append(scoreFieldRaces, *scoreFieldRace)

	}
	return &entities.FieldRace{
		UUID:          req.FieldRaceUUID,
		Consideration: scoreFieldRaces,
	}
}

type UploadScoreFieldRaceResponseJSON UploadScoreFieldRace

func (m *UploadScoreFieldRaceResponseJSON) Parse(c *gin.Context, input *entities.FieldRace) *UploadScoreFieldRaceResponseJSON {
	var fieldRaces []ConsiderationFieldRace

	for _, value := range input.Consideration {
		scoreFieldRace := &ConsiderationFieldRace{
			ID:       value.ID,
			NameTeam: value.NameTeam,
			Score:    value.Score,
			Title:    value.Title,
		}

		fieldRaces = append(fieldRaces, *scoreFieldRace)
	}

	fieldRaceScore := &UploadScoreFieldRaceResponseJSON{
		FieldRaceUUID:          input.UUID,
		ConsiderationFieldRace: fieldRaces,
	}

	return fieldRaceScore

}
