package users

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/deliveries/http/middlewares"
	"gitlab.com/chaihanij/evat/app/layers/usecase/users"
)

type Handler struct {
	UsersUseCase users.UseCase
}

func NewEndpointHttpHandler(ginEngine *gin.Engine, authMiddleware middlewares.AuthMiddleware, usersUseCase users.UseCase) {
	handler := &Handler{
		UsersUseCase: usersUseCase,
	}
	v1 := ginEngine.Group("v1")
	{
		v1.POST("/users.login", handler.Login)
		v1.POST("/users.forget-password", handler.ForgetPassword)
		v1.POST("/users.reset-password-by-otp", handler.ResetPasswordByOTP)

	}

	v1Auth := ginEngine.Group("v1").Use(authMiddleware.Authentication)
	{
		v1Auth.GET("/profile.one", handler.FineOneUserProfile)
		v1Auth.POST("/profile.update", handler.UpdateUserProfile)
		// ADMIN
		v1Auth.GET("/users.all", handler.FinAllUser)
		v1Auth.POST("/users.create", handler.CreateUser)
		v1Auth.POST("/users.update", handler.UpdateUser)
		v1Auth.POST("/users.delete/:uid", handler.DeleteUser)
	}
}
