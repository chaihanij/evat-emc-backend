package dtos

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
)

type ConfigReq struct {
	ConfigUUID string `uri:"config_uuid"`

	// RegisterConfig RegisterConfig `json:"register_config" bson:"register_config"`
}

// type RegisterConfig struct {
// 	StartDate time.Time `json:"start_date" bson:"start_date"`
// 	EndDate   time.Time `json:"end_date" bson:"end_date"`
// }

// func (Register *RegisterConfig) Parse(c *gin.Context, data *entities.DateRegisterConfig) *RegisterConfig {
// 	copier.Copy(Register, data)
// 	return Register

// }

func (req *ConfigReq) Parse(c *gin.Context) (*ConfigReq, error) {
	if err := c.ShouldBindUri(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}

	}
	return req, nil
}

func (req *ConfigReq) ToEntity() *entities.Config {

	return &entities.Config{
		UUID: req.ConfigUUID,
	}

}

type ConfigRes struct {
	Uuid           string      `json:"uuid"`
	Registerconfig interface{} `json:"registerconfig"`
	Startproject   interface{} `json:"startproject"`
}

// type Registerconfig struct {
// 	StartDate time.Time `json:"start_date" bson:"start_date"`
// 	EndDate   time.Time `json:"end_date" bson:"end_date"`
// }

// type Startproject struct {
// 	StartDate time.Time `json:"start_date" bson:"start_date"`
// 	EndDate   time.Time `json:"end_date" bson:"end_date"`
// }

func (m *ConfigRes) Parse(c *gin.Context, input *entities.Config) *ConfigRes {

	// var reg interface{} = input.RegisterConfig
	// // fmt.Println("reg :", reg)

	// register, _ := reg.(entities.DateRegisterConfig)
	// fmt.Println("reg :", register)
	// var startproject interface{} = input.StartProject
	// start, _ := startproject.(Startproject)

	// p, ok := i.()

	// var registerconfig Registerconfig
	// var startproject Startproject

	// reg := input.RegisterConfig.(entities.DateRegisterConfig)
	// project := input.RegisterConfig.(entities.DateStartProject)

	// registerconfig.StartDate = reg.Start_date
	// registerconfig.EndDate = reg.End_date

	// startproject.StartDate = project.Start_date
	// startproject.EndDate = project.End_date

	// // startproject := input.StartProject.(entities.Da)

	configRes := &ConfigRes{
		Uuid:           input.UUID,
		Registerconfig: input.RegisterConfig,
		Startproject:   input.StartProject,
	}
	return configRes
}
