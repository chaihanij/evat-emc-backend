package files

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/chaihanij/evat/app/layers/usecase/files"
)

type Handler struct {
	FilesUseCase files.UseCase
}

func NewEndpointHttpHandler(ginEngine *gin.Engine, filesUseCase files.UseCase) {
	handler := &Handler{
		FilesUseCase: filesUseCase,
	}
	v1 := ginEngine.Group("v1")
	{
		v1.GET("/files/:file_uuid", handler.FindOneFile)
	}
}
