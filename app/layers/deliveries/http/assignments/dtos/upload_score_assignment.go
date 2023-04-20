package dtos

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gitlab.com/chaihanij/evat/app/constants"
	"gitlab.com/chaihanij/evat/app/entities"

	"gitlab.com/chaihanij/evat/app/errors"
)

type UploadScoreAssingment struct {
	AssignmentUUID          string                    `uri:"assignment_uuid"`
	ConsiderationAssignment []ConsiderationAssignment `json:"consideration" bson:"consideration"`
}

type ConsiderationAssignment struct {
	ID        string  `json:"id" bson:"id"`
	Title     string  `json:"title" bson:"title"`
	Team_type string  `json:"team_type" bson:"team_type"`
	NameTeam  string  `json:"nameteam" bson:"nameteam"`
	Score     float64 `json:"score" bson:"score"`
}

func (req *UploadScoreAssingment) Parse(c *gin.Context) (*UploadScoreAssingment, error) {
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
	// req.UpdatedBy = jwtData.UID
	return req, nil
}

func (req *UploadScoreAssingment) ToEntity() *entities.Assignment {

	var scoreAssignments []entities.ConsiderationAssignment

	for _, value := range req.ConsiderationAssignment {

		scoreAssignment := &entities.ConsiderationAssignment{
			ID:       value.ID,
			NameTeam: value.NameTeam,
			Score:    value.Score,
			Title:    value.Title,
			TeamType: value.Team_type,
		}

		scoreAssignments = append(scoreAssignments, *scoreAssignment)

	}
	return &entities.Assignment{
		UUID:          req.AssignmentUUID,
		Consideration: scoreAssignments,
	}
}

type UploadScoreAssignmentResponseJSON UploadScoreAssingment

func (m *UploadScoreAssignmentResponseJSON) Parse(c *gin.Context, input *entities.Assignment) *UploadScoreAssignmentResponseJSON {

	var scoreAssignments []ConsiderationAssignment

	for _, value := range input.Consideration {

		scoreAssignment := &ConsiderationAssignment{
			ID:       value.ID,
			NameTeam: value.NameTeam,
			Score:    value.Score,
			Title:    value.Title,
		}

		scoreAssignments = append(scoreAssignments, *scoreAssignment)

	}

	assignmentScore := &UploadScoreAssignmentResponseJSON{
		AssignmentUUID:          input.UUID,
		ConsiderationAssignment: scoreAssignments,
	}

	return assignmentScore

}
