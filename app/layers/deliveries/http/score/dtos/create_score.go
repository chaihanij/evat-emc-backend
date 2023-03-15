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
	ID             primitive.ObjectID `json:"_id" bson:"_id"`
	UID            string             `json:"uid" bson:"uid"`
	NameTeam       string             `json:"nameteam" bson:"nameteam"`
	FirstTeam      string            `json:"firstteam" bson:"firstteam"`
	SecondTeam     string            `json:"secondteam" bson:"secondteam"`
	First_Stadium  string            `json:"firststadium" bson:"firststadium"`
	Second_Stadium string            `json:"secondstadium" bson:"secondstadium"`
	Third_Stadium  string            `json:"thirdstadium" bson:"thirdstadium"`
	Fourth_Stadium string            `json:"fourthstadium" bson:"fourthstadium"`
	Sum_Score      string            `json:"sumscore" bson:"sumscore"`
	No             string            `json:"no" bson:"no"`
	CreateDate     time.Time          `json:"createdate" bson:"createdate"`
	Lastupdate     time.Time          `json:"lastupdate" bson:"lastupdate"`
	CreateBy       string             `json:"createby" bson:"createby"`
	LastUpdateBy   string             `json:"lastupdateby" bson:"lastupdateby"`
}

type CreateScoreResponseJSON struct {
	// ID             primitive.ObjectID `json:"_id" bson:"_id"`
	UID            string    `json:"uid" bson:"uid"`
	NameTeam       string    `json:"nameteam" bson:"nameteam"`
	FirstTeam      string   `json:"firstteam" bson:"firstteam"`
	SecondTeam     string   `json:"secondteam" bson:"secondteam"`
	First_Stadium  string   `json:"firststadium" bson:"firststadium"`
	Second_Stadium string   `json:"secondstadium" bson:"secondstadium"`
	Third_Stadium  string   `json:"thirdstadium" bson:"thirdstadium"`
	Fourth_Stadium string   `json:"fourthstadium" bson:"fourthstadium"`
	Sum_Score      string   `json:"sumscore" bson:"sumscore"`
	No             string   `json:"no" bson:"no"`
	CreateDate     time.Time `json:"createdate" bson:"createdate"`
	Lastupdate     time.Time `json:"lastupdate" bson:"lastupdate"`
	CreateBy       string    `json:"createby" bson:"createby"`
	LastUpdateBy   string    `json:"lastupdateby" bson:"lastupdateby"`
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
		// ID: string(primitive.NewObjectID()),
		// UUID:           req.UID,
		NameTeam:       req.NameTeam,
		FirstTeam:      req.FirstTeam,
		SecondTeam:     req.SecondTeam,
		First_Stadium:  req.First_Stadium,
		Second_Stadium: req.Second_Stadium,
		Third_Stadium:  req.Third_Stadium,
		Fourth_Stadium: req.Fourth_Stadium,
		Sum_Score:      req.Sum_Score,
		No:             req.No,
		CreateDate:     time.Now(),
		CreateBy:       req.CreateBy,
		Lastupdate:     time.Now(),
		LastUpdateBy:   req.LastUpdateBy,
	}

}

func (m *CreateScoreResponseJSON) Parse(input *entities.Score) *CreateScoreResponseJSON {

	return &CreateScoreResponseJSON{

		// ID: string(primitive.NewObjectID()),
		// UID:            input.UUID,
		NameTeam:       input.NameTeam,
		FirstTeam:      input.FirstTeam,
		SecondTeam:     input.SecondTeam,
		First_Stadium:  input.First_Stadium,
		Second_Stadium: input.Second_Stadium,
		Third_Stadium:  input.Third_Stadium,
		Fourth_Stadium: input.Fourth_Stadium,
		Sum_Score:      input.Sum_Score,
		No:             input.No,
		CreateDate:     time.Now(),
		CreateBy:       input.CreateBy,
		Lastupdate:     time.Now(),
		LastUpdateBy:   input.LastUpdateBy,
	}

}

type CreateScoreResponseSwagger struct {
	StatusCode    int                    `json:"statusCode" example:"1000"`
	StatusMessage string                 `json:"statusMessage" example:"Success"`
	Timestamp     time.Time              `json:"timestamp" example:"2015-06-30T21:59:59Z"`
	Data          CreateScoreRequestJSON `json:"data,omitempty"`
}
