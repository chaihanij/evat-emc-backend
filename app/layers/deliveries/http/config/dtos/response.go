package dtos

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gitlab.com/chaihanij/evat/app/entities"
	"gitlab.com/chaihanij/evat/app/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Config struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id"`
	RegisterConfig RegisterConfig     `json:"register_config" bson:"register_config"`
}

type RegisterConfig struct {
	StartDate time.Time `json:"start_date" bson:"start_date"`
	EndDate   time.Time `json:"end_date" bson:"end_date"`
}

func (Register *RegisterConfig) Parse(c *gin.Context, data *entities.RegisterConfig) *RegisterConfig {
	copier.Copy(Register, data)
	return Register

}

func (req *Config) Parse(c *gin.Context) (*Config, error) {
	if err := c.ShouldBindJSON(req); err != nil {
		return nil, errors.ParameterError{Message: err.Error()}

	}
	return req, nil
}

type ConfigResponse Config
type RegisterConfigs []RegisterConfig

func (res *ConfigResponse) Parse(c *gin.Context, input *entities.Config) *ConfigResponse {
	var ConFig RegisterConfig

	if val, ok := input.RegisterConfig.(entities.RegisterConfig); ok {

		ConFig.StartDate = val.Start_date
		ConFig.EndDate = val.End_date
		// 	for _, value := range val {
		// 		registerConfig := new(RegisterConfig).Parse(c, &value)

		// 		ConFig = append(ConFig, *registerConfig)

		// 	}
	}

	// ConFig.StartDate = input.RegisterConfig.

	return &ConfigResponse{
		// ID: input.ID,
		RegisterConfig: ConFig,
	}
}

func (req *Config) ToEntity() *entities.Config {
	return &entities.Config{
		RegisterConfig: req.RegisterConfig,
	}
}
