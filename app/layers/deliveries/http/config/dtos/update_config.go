package dtos

import (
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

type RequsetConfig struct {
	ConfigUUID    string    `uri:"config_uuid"`
	StartProject  time.Time `json:"start_project"`
	EndProject    time.Time `json:"end_project"`
	StartRegister time.Time `json:"start_register"`
	EndRegister   time.Time `json:"end_register"`
}

func (req *RequsetConfig) Parse(c *gin.Context) (*RequsetConfig, error) {

	if err := c.ShouldBindJSON(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	if err := c.ShouldBindUri(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}
	}
	return req, nil
}

func (req *RequsetConfig) ToEntity() *entities.Config {

	var registerConfig entities.DateRegisterConfig

	registerConfig.Start_date = req.StartRegister
	registerConfig.End_date = req.EndRegister

	var project entities.DateStartProject

	project.Start_date = req.StartProject
	project.End_date = req.EndProject

	config := &entities.Config{
		StartProject:   project,
		RegisterConfig: registerConfig,
		UUID:           req.ConfigUUID,
	}

	return config

}
