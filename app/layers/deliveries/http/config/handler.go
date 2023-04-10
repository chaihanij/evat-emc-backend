package config

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/middlewares"
	"gitlab.com/chaihanij/evat/app/layers/usecase/config"
)

type Handler struct {
	ConfigUseCase config.UseCase
}

func NewEndpointHttpHandler(ginEngine *gin.Engine, authMiddleware middlewares.AuthMiddleware, config config.UseCase) {
	handler := &Handler{
		ConfigUseCase: config,
	}
	v1Auth := ginEngine.Group("v1").Use(authMiddleware.Authentication)
	{
		v1Auth.GET("/config", handler.FindOneConfig)
	}
}
