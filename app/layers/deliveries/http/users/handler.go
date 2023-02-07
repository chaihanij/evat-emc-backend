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
		v1.POST("/login", handler.Login)
		v1.POST("/forget-password", handler.ForgetPassword)
		v1.POST("/reset-password-by-otp", handler.ResetPasswordByOTP)
		v1.POST("/superadmin", handler.CreateSuperAdmin)
	}

	v1Auth := ginEngine.Group("v1").Use(authMiddleware.Authentication)
	{
		v1Auth.GET("/users", handler.FindAllUser)
		v1Auth.POST("/users", handler.CreateUser)
		v1Auth.GET("/users/:uid", handler.FineOneUser)
		v1Auth.PUT("/users/:uid", handler.UpdateUser)
		v1Auth.DELETE("/users/:uid", handler.DeleteUser)

		v1Auth.GET("/profiles", handler.FindOneUserProfile)
		v1Auth.PUT("/profiles/:uid", handler.UpdateUserProfile)

	}
}
